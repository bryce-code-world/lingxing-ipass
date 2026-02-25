package dsco

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// StreamService 封装 Stream 相关接口。
//
// Stream 用于持续拉取与全量同步（sync operation）：
// - inventory：建议用 Stream + sync 进行全量库存拉取/对账，而不是循环调用 GET /inventory（该接口限流且不适合批量）。
type StreamService struct {
	c *Client
}

func escapePath(s string) string {
	return url.PathEscape(strings.TrimSpace(s))
}

func escapeStreamPosition(pos string) string {
	pos = strings.TrimSpace(pos)
	if pos == "" {
		return ""
	}
	// DSCO stream event id/position 在部分场景下会包含已编码的 "%2F"。
	// 为避免出现二次编码（"%252F"）导致 position invalid，这里先做一次解码再重新编码。
	if decoded, err := url.PathUnescape(pos); err == nil {
		pos = decoded
	}
	return url.PathEscape(pos)
}

// CreateStream 创建一个 stream（POST /stream）。
func (s *StreamService) CreateStream(ctx context.Context, stream Stream) (*Stream, error) {
	if s == nil || s.c == nil {
		return nil, errors.New("dsco client 不能为空")
	}
	var out Stream
	if err := s.c.doJSON(ctx, http.MethodPost, "/stream", nil, stream, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ListStreams 列出 streams（GET /stream，可选按 id 过滤）。
func (s *StreamService) ListStreams(ctx context.Context, id string) ([]Stream, error) {
	if s == nil || s.c == nil {
		return nil, errors.New("dsco client 不能为空")
	}
	type q struct {
		ID string `url:"id,omitempty"`
	}
	var out []Stream
	if err := s.c.doJSON(ctx, http.MethodGet, "/stream", q{ID: strings.TrimSpace(id)}, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteStream 删除一个 stream（DELETE /stream/{id}）。
func (s *StreamService) DeleteStream(ctx context.Context, id string) (*Stream, error) {
	if s == nil || s.c == nil {
		return nil, errors.New("dsco client 不能为空")
	}
	id = strings.TrimSpace(id)
	if id == "" {
		return nil, errors.New("stream id 不能为空")
	}
	var out Stream
	if err := s.c.doJSON(ctx, http.MethodDelete, "/stream/"+escapePath(id), nil, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateStreamOperation 创建一个 stream operation（POST /stream/{id}）。
//
// 说明：
// - 对于 inventory 全量拉取，使用 SyncStreamOperation{OperationType:"sync"}。
// - 该操作为异步，服务端会逐步把对象 dump 到 stream。
func (s *StreamService) CreateStreamOperation(ctx context.Context, id string, op any) (*StreamOperationResponse, error) {
	if s == nil || s.c == nil {
		return nil, errors.New("dsco client 不能为空")
	}
	id = strings.TrimSpace(id)
	if id == "" {
		return nil, errors.New("stream id 不能为空")
	}
	var out StreamOperationResponse
	if err := s.c.doJSON(ctx, http.MethodPost, "/stream/"+escapePath(id), nil, op, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetStreamEventsFromPosition 从 position 拉取 events（GET /stream/{id}/{partitionId}/{position}）。
func (s *StreamService) GetStreamEventsFromPosition(ctx context.Context, id string, partitionID int, position string) (*StreamEventWrapper, error) {
	if s == nil || s.c == nil {
		return nil, errors.New("dsco client 不能为空")
	}
	id = strings.TrimSpace(id)
	position = strings.TrimSpace(position)
	if id == "" {
		return nil, errors.New("stream id 不能为空")
	}
	if partitionID < 0 {
		return nil, errors.New("partitionId 不允许为负数")
	}
	if position == "" {
		// 文档允许 position 使用 eventId 或 ISO8601 时间；为空时不明确，直接拒绝。
		return nil, errors.New("position 不能为空")
	}
	var out StreamEventWrapper
	path := fmt.Sprintf("/stream/%s/%d/%s", escapePath(id), partitionID, escapeStreamPosition(position))
	if err := s.c.doJSON(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetStreamEventsInRange 在区间内拉取 events（GET /stream/{id}/{partitionId}/{startPosition}/{endPosition}）。
func (s *StreamService) GetStreamEventsInRange(ctx context.Context, id string, partitionID int, startPosition string, endPosition string) (*StreamEventWrapper, error) {
	if s == nil || s.c == nil {
		return nil, errors.New("dsco client 不能为空")
	}
	id = strings.TrimSpace(id)
	startPosition = strings.TrimSpace(startPosition)
	endPosition = strings.TrimSpace(endPosition)
	if id == "" {
		return nil, errors.New("stream id 不能为空")
	}
	if partitionID < 0 {
		return nil, errors.New("partitionId 不允许为负数")
	}
	if startPosition == "" || endPosition == "" {
		return nil, errors.New("start/end position 不能为空")
	}
	var out StreamEventWrapper
	path := fmt.Sprintf("/stream/%s/%d/%s/%s", escapePath(id), partitionID, escapeStreamPosition(startPosition), escapeStreamPosition(endPosition))
	if err := s.c.doJSON(ctx, http.MethodGet, path, nil, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateStreamPosition 更新 stream partition position（PUT /stream/{id}/{partitionId}/{position}）。
func (s *StreamService) UpdateStreamPosition(ctx context.Context, id string, partitionID int, position string) (*SuccessFailResponse, error) {
	if s == nil || s.c == nil {
		return nil, errors.New("dsco client 不能为空")
	}
	id = strings.TrimSpace(id)
	position = strings.TrimSpace(position)
	if id == "" {
		return nil, errors.New("stream id 不能为空")
	}
	if partitionID < 0 {
		return nil, errors.New("partitionId 不允许为负数")
	}
	if position == "" {
		return nil, errors.New("position 不能为空")
	}
	var out SuccessFailResponse
	path := fmt.Sprintf("/stream/%s/%d/%s", escapePath(id), partitionID, escapeStreamPosition(position))
	if err := s.c.doJSON(ctx, http.MethodPut, path, nil, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// StreamPartitionReader 是一个轻量的“流式拉取器”：按当前 position 拉取 events，并在处理后 ack。
// 调用方负责决定何时停止拉取、以及 ack 的频率。
type StreamPartitionReader struct {
	svc *StreamService

	StreamID    string
	PartitionID int
	Position    string
}

func NewStreamPartitionReader(svc *StreamService, streamID string, partitionID int, position string) *StreamPartitionReader {
	return &StreamPartitionReader{
		svc:         svc,
		StreamID:    strings.TrimSpace(streamID),
		PartitionID: partitionID,
		Position:    strings.TrimSpace(position),
	}
}

// Pull 拉取一批 events（从 reader.Position 之后）。
func (r *StreamPartitionReader) Pull(ctx context.Context) (*StreamEventWrapper, error) {
	if r == nil || r.svc == nil {
		return nil, errors.New("stream reader 未初始化")
	}
	return r.svc.GetStreamEventsFromPosition(ctx, r.StreamID, r.PartitionID, r.Position)
}

// Ack 将 reader.Position 更新为指定 position（一般为最后一个 StreamEvent.id）。
func (r *StreamPartitionReader) Ack(ctx context.Context, position string) (*SuccessFailResponse, error) {
	if r == nil || r.svc == nil {
		return nil, errors.New("stream reader 未初始化")
	}
	position = strings.TrimSpace(position)
	if position == "" {
		return nil, errors.New("position 不能为空")
	}
	resp, err := r.svc.UpdateStreamPosition(ctx, r.StreamID, r.PartitionID, position)
	if err != nil {
		return nil, err
	}
	r.Position = position
	return resp, nil
}

// WaitUntilHasEvents 简单等待：直到 stream 返回至少 1 条 event，或超时/ctx 取消。
// 适用于 sync operation 刚触发时等待首批数据出现。
func (r *StreamPartitionReader) WaitUntilHasEvents(ctx context.Context, pollInterval time.Duration, timeout time.Duration) error {
	if r == nil || r.svc == nil {
		return errors.New("stream reader 未初始化")
	}
	if pollInterval <= 0 {
		pollInterval = 2 * time.Second
	}
	if timeout <= 0 {
		timeout = 2 * time.Minute
	}
	deadline := time.Now().Add(timeout)
	for {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		w, err := r.Pull(ctx)
		if err != nil {
			return err
		}
		if len(w.Events) > 0 {
			return nil
		}
		if time.Now().After(deadline) {
			return errors.New("wait stream events timeout")
		}
		time.Sleep(pollInterval)
	}
}
