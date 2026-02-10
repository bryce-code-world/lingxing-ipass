package dsco_lingxing

import (
	"context"
	"time"

	"gitee.com/lsy007/golibv2/v2/sdk/dsco"
	"gitee.com/lsy007/golibv2/v2/sdk/lingxing"
)

func (d *Domain) dscoClient() (*dsco.Client, error) {
	return dsco.New(dsco.Config{
		BaseURL:    d.env.Integration.DSCO.BaseURL,
		Token:      d.env.Auth.DSCO.Token,
		HTTPClient: d.httpClient,
		UserAgent:  "lingxing-ipass",
	})
}

func (d *Domain) lingxingClient(ctx context.Context) (*lingxing.Client, error) {
	return lingxing.New(lingxing.Config{
		BaseURL:     d.env.Integration.LingXing.BaseURL,
		AppID:       d.env.Auth.LingXing.AppID,
		AppSecret:   d.env.Auth.LingXing.AppSecret,
		AccessToken: "",
		AutoToken:   true,
		HTTPClient:  d.httpClient,
		UserAgent:   "lingxing-ipass",
		Now:         func() time.Time { return time.Now().UTC() },
	})
}
