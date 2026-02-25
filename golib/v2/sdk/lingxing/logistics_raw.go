package lingxing

import (
	"context"
	"net/http"
)

func (s *LogisticsService) ChannelListWithRawBody(ctx context.Context, req ChannelListRequest) ([]LogisticsChannel, int, string, error) {
	var items []LogisticsChannel
	env, raw, err := s.c.doSignedJSONWithEnvelopeWithRawBody(ctx, http.MethodPost, "/erp/sc/data/local_inventory/channelList", nil, req, &items)
	if err != nil {
		return nil, 0, raw, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return items, total, raw, nil
}

func (s *LogisticsService) QueryHeadLogisticsProviderListWithRawBody(ctx context.Context, req QueryHeadLogisticsProviderListRequest) (HeadLogisticsProviderListResponse, string, error) {
	var out HeadLogisticsProviderListResponse
	_, raw, err := s.c.doSignedJSONWithEnvelopeWithRawBody(ctx, http.MethodPost, "/basicOpen/logistics/headLogisticsProvider/query/list", nil, req, &out)
	return out, raw, err
}
