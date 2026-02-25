package dsco

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (s *InventoryService) GetByKeyWithRawBody(ctx context.Context, q InventoryGetQuery) (*ItemInventory, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out ItemInventory
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodGet, "/inventory", q, nil, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *InventoryService) UpsertSingleWithRawBody(ctx context.Context, inv *ItemInventory) (*SuccessFailResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out SuccessFailResponse
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodPost, "/inventory/singleItem", nil, inv, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *InventoryService) UpdateSmallBatchWithRawBody(ctx context.Context, items []ItemInventory, q InventoryUpdateSmallBatchQuery) (*AsyncUpdateResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out AsyncUpdateResponse
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodPost, "/inventory/batch/small", q, items, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *InventoryService) InitLargeBatchWithRawBody(ctx context.Context) (*LargeBatchUpdateResponse, string, error) {
	return s.InitLargeBatchWithQueryWithRawBody(ctx, InventoryUpdateLargeBatchQuery{})
}

func (s *InventoryService) InitLargeBatchWithQueryWithRawBody(ctx context.Context, q InventoryUpdateLargeBatchQuery) (*LargeBatchUpdateResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}

	var out LargeBatchUpdateResponse
	raw, err := s.c.doJSONWithRawBody(ctx, http.MethodPost, "/inventory/batch/large", q, nil, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}

func (s *InventoryService) UploadLargeBatchWithRawBody(ctx context.Context, dataURL string, r io.Reader) (string, error) {
	if s == nil || s.c == nil {
		return "", errors.New("dsco client 不能为 nil")
	}
	dataURL = strings.TrimSpace(dataURL)
	if dataURL == "" {
		return "", errors.New("dataUrl 不能为空")
	}
	u, err := url.Parse(dataURL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return "", errors.New("dataUrl 非法")
	}
	if r == nil {
		return "", errors.New("upload body 不能为空")
	}

	if strings.TrimSpace(s.c.token) == "" {
		return "", ErrMissingToken
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, dataURL, r)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if strings.TrimSpace(s.c.userAgent) != "" {
		req.Header.Set("User-Agent", s.c.userAgent)
	}
	req.Header.Set("Authorization", "bearer "+s.c.token)

	resp, err := s.c.httpCli.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	rawStr := string(raw)

	if resp.StatusCode >= 400 {
		return rawStr, &APIError{
			StatusCode: resp.StatusCode,
			Method:     req.Method,
			URL:        dataURL,
			Body:       raw,
		}
	}
	return rawStr, nil
}
