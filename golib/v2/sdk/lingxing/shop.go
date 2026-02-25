package lingxing

import (
	"context"
	"net/http"
)

// ShopService 提供店铺（店铺列表、店铺重命名等）相关接口。
type ShopService struct {
	c *Client
}

// AmazonSellerList 查询亚马逊店铺列表。
//
// API Path: /erp/sc/data/seller/lists
func (s *ShopService) AmazonSellerList(ctx context.Context) ([]AmazonSeller, error) {
	var out []AmazonSeller
	if err := s.c.doSignedGET(ctx, "/erp/sc/data/seller/lists", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// AmazonConceptSellerList 查询亚马逊概念店铺列表。
//
// 注意：该接口的 total 位于响应顶层（不在 data 内部）。
//
// API Path: /erp/sc/data/seller/conceptLists
func (s *ShopService) AmazonConceptSellerList(ctx context.Context) (AmazonConceptSellerListResponse, error) {
	var list []AmazonConceptSeller
	env, err := s.c.doSignedJSONWithEnvelope(ctx, http.MethodGet, "/erp/sc/data/seller/conceptLists", nil, nil, &list)
	if err != nil {
		return AmazonConceptSellerListResponse{}, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return AmazonConceptSellerListResponse{
		Total: total,
		List:  list,
	}, nil
}

// AmazonSellerBatchRename 批量修改店铺名称（最多 10 个）。
//
// API Path: /erp/sc/data/seller/batchEditSellerName
func (s *ShopService) AmazonSellerBatchRename(ctx context.Context, req AmazonSellerBatchRenameRequest) (AmazonSellerBatchRenameResponseData, error) {
	var out AmazonSellerBatchRenameResponseData
	_, err := s.c.doSignedJSONWithEnvelope(ctx, http.MethodPost, "/erp/sc/data/seller/batchEditSellerName", nil, req, &out)
	return out, err
}

// MultiPlatformStoreListV2 查询多平台店铺信息。
//
// API Path: /pb/mp/shop/v2/getSellerList
func (s *ShopService) MultiPlatformStoreListV2(ctx context.Context, req MultiPlatformStoreListV2Request) (MultiPlatformStoreListV2ResponseData, error) {
	var out MultiPlatformStoreListV2ResponseData
	_, err := s.c.doSignedJSONWithEnvelope(ctx, http.MethodPost, "/pb/mp/shop/v2/getSellerList", nil, req, &out)
	return out, err
}
