package lingxing

import (
	"context"
	"errors"
	"net/http"
)

// ConfigService 提供业务配置相关接口。
type ConfigService struct {
	c *Client
}

// GetPairListV2 查询多平台 SKU 配对列表。
//
// API Path: /pb/mp/listing/v2/getPairList
func (s *ConfigService) GetPairListV2(ctx context.Context, req PairListV2Request) (PairListV2ResponseData, error) {
	if req.Offset < 0 {
		return PairListV2ResponseData{}, errors.New("offset 不能小于 0")
	}

	// 文档中 length 为可选，但服务端通常会对 length 有最小值/非 0 限制；这里给一个安全默认值。
	if req.Length <= 0 {
		req.Length = 20
	}

	if req.UseCursor != nil && *req.UseCursor && req.CursorID == nil {
		return PairListV2ResponseData{}, errors.New("use_cursor=true 时 cursor_id 必填")
	}

	var out PairListV2ResponseData
	_, err := s.c.doSignedJSONWithEnvelope(ctx, http.MethodPost, "/pb/mp/listing/v2/getPairList", nil, req, &out)
	return out, err
}
