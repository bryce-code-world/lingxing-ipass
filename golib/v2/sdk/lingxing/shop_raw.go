package lingxing

import (
	"context"
	"net/http"
)

func (s *ShopService) AmazonSellerListWithRawBody(ctx context.Context) ([]AmazonSeller, string, error) {
	var out []AmazonSeller
	_, raw, err := s.c.doSignedGETWithEnvelopeWithRawBody(ctx, "/erp/sc/data/seller/lists", nil, &out)
	if err != nil {
		return nil, raw, err
	}
	return out, raw, nil
}

func (s *ShopService) AmazonConceptSellerListWithRawBody(ctx context.Context) (AmazonConceptSellerListResponse, string, error) {
	var list []AmazonConceptSeller
	env, raw, err := s.c.doSignedJSONWithEnvelopeWithRawBody(ctx, http.MethodGet, "/erp/sc/data/seller/conceptLists", nil, nil, &list)
	if err != nil {
		return AmazonConceptSellerListResponse{}, raw, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return AmazonConceptSellerListResponse{
		Total: total,
		List:  list,
	}, raw, nil
}

func (s *ShopService) AmazonSellerBatchRenameWithRawBody(ctx context.Context, req AmazonSellerBatchRenameRequest) (AmazonSellerBatchRenameResponseData, string, error) {
	var out AmazonSellerBatchRenameResponseData
	_, raw, err := s.c.doSignedJSONWithEnvelopeWithRawBody(ctx, http.MethodPost, "/erp/sc/data/seller/batchEditSellerName", nil, req, &out)
	return out, raw, err
}

func (s *ShopService) MultiPlatformStoreListV2WithRawBody(ctx context.Context, req MultiPlatformStoreListV2Request) (MultiPlatformStoreListV2ResponseData, string, error) {
	var out MultiPlatformStoreListV2ResponseData
	_, raw, err := s.c.doSignedJSONWithEnvelopeWithRawBody(ctx, http.MethodPost, "/pb/mp/shop/v2/getSellerList", nil, req, &out)
	return out, raw, err
}
