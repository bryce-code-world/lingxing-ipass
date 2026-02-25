package dsco

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// InventoryService 封装 Inventory 相关接口。
type InventoryService struct {
	c *Client
}

// InventoryUpdateSmallBatchQuery 表示 Update Inventory Small Batch 的查询参数。
type InventoryUpdateSmallBatchQuery struct {
	// SkipItemsThatDontExist 为 true 时，会在处理前过滤掉“不存在的 item”。
	SkipItemsThatDontExist *bool `url:"skipItemsThatDontExist,omitempty"`
}

// InventoryUpdateLargeBatchQuery 表示 Update Inventory Large Batch 的查询参数。
type InventoryUpdateLargeBatchQuery struct {
	// SkipItemsThatDontExist 为 true 时，会在处理前过滤掉“不存在的 item”。
	SkipItemsThatDontExist *bool `url:"skipItemsThatDontExist,omitempty"`
}

// LargeBatchUpdateResponse 表示 Large Batch API 的响应体（返回上传用 dataUrl）。
type LargeBatchUpdateResponse struct {
	Status    string               `json:"status"`
	RequestID string               `json:"requestId"`
	EventDate string               `json:"eventDate"`
	DataURL   string               `json:"dataUrl"`
	Messages  []APIResponseMessage `json:"messages,omitempty"`
}

// GetByKey 获取单个库存对象（GET /inventory）。
func (s *InventoryService) GetByKey(ctx context.Context, q InventoryGetQuery) (*ItemInventory, error) {
	var out ItemInventory
	if err := s.c.doJSON(ctx, http.MethodGet, "/inventory", q, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpsertSingle 创建/更新单个库存对象（POST /inventory/singleItem）。
func (s *InventoryService) UpsertSingle(ctx context.Context, inv *ItemInventory) (*SuccessFailResponse, error) {
	var out SuccessFailResponse
	if err := s.c.doJSON(ctx, http.MethodPost, "/inventory/singleItem", nil, inv, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateSmallBatch 异步批量更新库存（POST /inventory/batch/small）。
//
// 返回 AsyncUpdateResponse，其中 requestId 可用于后续通过 stream/change log 追踪处理结果。
func (s *InventoryService) UpdateSmallBatch(ctx context.Context, items []ItemInventory, q InventoryUpdateSmallBatchQuery) (*AsyncUpdateResponse, error) {
	if s == nil || s.c == nil {
		return nil, errors.New("dsco client 不能为空")
	}
	var out AsyncUpdateResponse
	if err := s.c.doJSON(ctx, http.MethodPost, "/inventory/batch/small", q, items, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// InitLargeBatch 初始化大批量库存更新（POST /inventory/batch/large）。
//
// 返回 dataUrl：调用方需要再对该 URL 发起 PUT（JSON Lines）上传数据。
func (s *InventoryService) InitLargeBatch(ctx context.Context) (*LargeBatchUpdateResponse, error) {
	return s.InitLargeBatchWithQuery(ctx, InventoryUpdateLargeBatchQuery{})
}

// InitLargeBatchWithQuery 初始化大批量库存更新（POST /inventory/batch/large），支持 skipItemsThatDontExist 查询参数。
//
// 返回 dataUrl：调用方需要再对该 URL 发起 PUT（JSON Lines）上传数据。
func (s *InventoryService) InitLargeBatchWithQuery(ctx context.Context, q InventoryUpdateLargeBatchQuery) (*LargeBatchUpdateResponse, error) {
	if s == nil || s.c == nil {
		return nil, errors.New("dsco client 不能为空")
	}
	var out LargeBatchUpdateResponse
	if err := s.c.doJSON(ctx, http.MethodPost, "/inventory/batch/large", q, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UploadLargeBatch 将 JSON Lines 数据上传到 InitLargeBatch 返回的 dataUrl（PUT dataUrl）。
//
// DSCO 文档要求：PUT 时使用与原 API 调用相同的 HTTP headers（至少 Authorization/Accept/Content-Type）。
func (s *InventoryService) UploadLargeBatch(ctx context.Context, dataURL string, r io.Reader) error {
	if s == nil || s.c == nil {
		return errors.New("dsco client 不能为空")
	}
	dataURL = strings.TrimSpace(dataURL)
	if dataURL == "" {
		return errors.New("dataUrl 不能为空")
	}
	u, err := url.Parse(dataURL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return errors.New("dataUrl 非法")
	}
	if r == nil {
		return errors.New("upload body 不能为空")
	}

	if strings.TrimSpace(s.c.token) == "" {
		return ErrMissingToken
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, dataURL, r)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if strings.TrimSpace(s.c.userAgent) != "" {
		req.Header.Set("User-Agent", s.c.userAgent)
	}
	req.Header.Set("Authorization", "bearer "+s.c.token)

	resp, err := s.c.httpCli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return &APIError{
			StatusCode: resp.StatusCode,
			Method:     req.Method,
			URL:        dataURL,
			Body:       raw,
		}
	}
	return nil
}
