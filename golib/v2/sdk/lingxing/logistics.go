package lingxing

import (
	"context"
	"net/http"
)

// LogisticsService 提供物流相关接口（头程物流等）。
type LogisticsService struct {
	c *Client
}

// ChannelList 查询头程物流渠道列表。
//
// API Path: /erp/sc/data/local_inventory/channelList
func (s *LogisticsService) ChannelList(ctx context.Context, req ChannelListRequest) ([]LogisticsChannel, int, error) {
	var items []LogisticsChannel
	env, err := s.c.doSignedJSONWithEnvelope(ctx, http.MethodPost, "/erp/sc/data/local_inventory/channelList", nil, req, &items)
	if err != nil {
		return nil, 0, err
	}
	total, _ := parseIntFromRaw(env.Total)
	return items, total, nil
}

// QueryHeadLogisticsProviderList 查询物流-头程物流商列表。
//
// API Path: /basicOpen/logistics/headLogisticsProvider/query/list
func (s *LogisticsService) QueryHeadLogisticsProviderList(ctx context.Context, req QueryHeadLogisticsProviderListRequest) (HeadLogisticsProviderListResponse, error) {
	var out HeadLogisticsProviderListResponse
	_, err := s.c.doSignedJSONWithEnvelope(ctx, http.MethodPost, "/basicOpen/logistics/headLogisticsProvider/query/list", nil, req, &out)
	return out, err
}
