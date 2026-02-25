package dsco

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"strings"
	"time"
)

// InventoryFullPullOptions 表示“全量库存拉取”的选项（基于 Stream + sync operation）。
//
// 设计取舍：
// - 本方法默认每次创建新 stream，拉取完成后删除（不保留 position 状态）。
// - stream sync 为异步操作，因此这里通过“轮询 + 空结果超时”判断完成；调用方可按数据量调大超时。
type InventoryFullPullOptions struct {
	// Description 可选，便于在 DSCO 后台排查。
	Description string

	// SKUs 可选：只同步这些 SKU（供应商侧 SKU）。
	SKUs []string

	// MaxEventsObjectCount 可选：单次拉取 events 的最大数量（DSCO 默认 1000）。
	MaxEventsObjectCount int

	// PollInterval 轮询间隔；默认 2 秒。
	PollInterval time.Duration

	// IdleTimeout 当持续无 events 时，达到该时长则认为已完成；默认 30 秒。
	IdleTimeout time.Duration

	// MaxDuration 最大执行时长（包含等待与拉取）；默认 30 分钟。
	MaxDuration time.Duration
}

func (o *InventoryFullPullOptions) normalize() {
	if o.PollInterval <= 0 {
		o.PollInterval = 2 * time.Second
	}
	if o.IdleTimeout <= 0 {
		o.IdleTimeout = 30 * time.Second
	}
	if o.MaxDuration <= 0 {
		o.MaxDuration = 30 * time.Minute
	}
}

// StartFullSyncStream 创建 inventory stream 并触发 sync operation，返回一个 reader 与 delete 函数。
func (s *InventoryService) StartFullSyncStream(ctx context.Context, opt InventoryFullPullOptions) (*StreamPartitionReader, func(context.Context) error, error) {
	if s == nil || s.c == nil {
		return nil, nil, errors.New("dsco client 不能为空")
	}
	if s.c.Stream == nil {
		return nil, nil, errors.New("dsco stream service 未初始化")
	}

	opt.normalize()

	var maxEvents *StreamMaxEvents
	if opt.MaxEventsObjectCount > 0 {
		n := opt.MaxEventsObjectCount
		maxEvents = &StreamMaxEvents{ObjectCount: &n}
	}

	stream, err := s.c.Stream.CreateStream(ctx, Stream{
		ObjectType:    StreamObjectTypeInventory,
		Description:   strings.TrimSpace(opt.Description),
		NumPartitions: 1,
		MaxEvents:     maxEvents,
		Query: InventoryQuery{
			QueryType: "inventory",
		},
	})
	if err != nil {
		return nil, nil, err
	}
	if stream == nil || strings.TrimSpace(stream.ID) == "" {
		return nil, nil, errors.New("create stream failed: missing stream id")
	}
	if len(stream.Partitions) == 0 {
		return nil, nil, errors.New("create stream failed: missing partitions")
	}

	partitionID := stream.Partitions[0].PartitionID
	if partitionID < 0 {
		return nil, nil, errors.New("create stream failed: invalid partitionId")
	}
	// 对于“全量拉取”，我们不使用服务端返回的当前 position：
	// - sync operation dump 出来的 events，其 event.id 可能早于“当前 position”（例如按对象 lastUpdate 生成 id），导致一直拉不到数据。
	// - 新建 stream 时使用 epoch 作为起点更稳妥。
	position := "1970-01-01T00:00:00Z"

	_, err = s.c.Stream.CreateStreamOperation(ctx, stream.ID, SyncStreamOperation{
		OperationType: "sync",
		SKU:           opt.SKUs,
	})
	if err != nil {
		// 尝试清理 stream
		_, _ = s.c.Stream.DeleteStream(context.Background(), stream.ID)
		return nil, nil, err
	}

	reader := NewStreamPartitionReader(s.c.Stream, stream.ID, partitionID, position)
	deleteFn := func(delCtx context.Context) error {
		_, err := s.c.Stream.DeleteStream(delCtx, stream.ID)
		return err
	}
	return reader, deleteFn, nil
}

// PullAll 通过 Stream + sync operation 拉取“全量库存”。
//
// handler 约束：
// - 必须是幂等或可容忍重复（stream 拉取在网络抖动/超时下可能出现重复拉取）。
func (s *InventoryService) PullAll(ctx context.Context, opt InventoryFullPullOptions, handler func(inv *ItemInventory) error) error {
	if handler == nil {
		return errors.New("handler 不能为空")
	}
	opt.normalize()

	runCtx, cancel := context.WithTimeout(ctx, opt.MaxDuration)
	defer cancel()

	reader, deleteFn, err := s.StartFullSyncStream(runCtx, opt)
	if err != nil {
		return err
	}
	defer func() { _ = deleteFn(context.Background()) }()

	// 等待首批 events 出现（避免刚触发 sync 时误判“已完成”）。
	// 注意：sync operation 是异步的，大数据量下可能需要更久才会产出首批 events。
	firstEventTimeout := opt.IdleTimeout
	if firstEventTimeout < 2*time.Minute {
		firstEventTimeout = 2 * time.Minute
	}
	if err := reader.WaitUntilHasEvents(runCtx, opt.PollInterval, firstEventTimeout); err != nil {
		return err
	}

	lastNonEmptyAt := time.Now()
	lastMaxPos := ""
	stableMaxPosSince := time.Time{}

	normalizePosForCompare := func(pos string) string {
		pos = strings.TrimSpace(pos)
		if pos == "" {
			return ""
		}
		if decoded, err := url.PathUnescape(pos); err == nil {
			pos = decoded
		}
		return pos
	}

	for {
		if runCtx.Err() != nil {
			return runCtx.Err()
		}

		w, err := reader.Pull(runCtx)
		if err != nil {
			return err
		}

		if len(w.Events) == 0 {
			// 尝试更快结束：
			// - DSCO Stream 的 maxPosition 表示当前“最新事件位置”（近似值，可能滞后约 1 分钟）
			// - 当 position == maxPosition 且连续多次拉不到新 events 时，可认为全量已拉完。
			// - 若 ListStreams 失败（网络/权限/服务端异常），则忽略并回退到 IdleTimeout。
			if s.c != nil && s.c.Stream != nil && strings.TrimSpace(reader.StreamID) != "" {
				if streams, err := s.c.Stream.ListStreams(runCtx, reader.StreamID); err == nil && len(streams) > 0 {
					for _, p := range streams[0].Partitions {
						if p.PartitionID != reader.PartitionID {
							continue
						}
						maxPos := strings.TrimSpace(p.MaxPosition)
						if maxPos == "" {
							break
						}
						cur := normalizePosForCompare(reader.Position)
						max := normalizePosForCompare(maxPos)
						if cur != "" && cur == max {
							// maxPosition 是“近似值”，且可能滞后，短时间内相等并不代表真正结束。
							// 这里要求 position==maxPosition 持续稳定至少 1 分钟，才认为已完成，
							// 否则可能过早退出导致只拉到部分全量数据。
							if maxPos == lastMaxPos {
								if stableMaxPosSince.IsZero() {
									stableMaxPosSince = time.Now()
								}
							} else {
								lastMaxPos = maxPos
								stableMaxPosSince = time.Now()
							}
						} else {
							lastMaxPos = maxPos
							stableMaxPosSince = time.Time{}
						}
						break
					}
				}
			}
			// 仅当“相等且稳定”持续一段时间，才允许提前结束。
			if !stableMaxPosSince.IsZero() && time.Since(stableMaxPosSince) >= time.Minute {
				return nil
			}
			if time.Since(lastNonEmptyAt) >= opt.IdleTimeout {
				return nil
			}
			time.Sleep(opt.PollInterval)
			continue
		}

		lastNonEmptyAt = time.Now()
		stableMaxPosSince = time.Time{}

		lastID := ""
		for _, ev := range w.Events {
			lastID = strings.TrimSpace(ev.ID)
			if lastID == "" {
				continue
			}
			var inv ItemInventory
			if err := json.Unmarshal(ev.Payload, &inv); err != nil {
				return err
			}
			if err := handler(&inv); err != nil {
				return err
			}
		}
		if lastID != "" {
			if _, err := reader.Ack(runCtx, lastID); err != nil {
				return err
			}
		}
	}
}
