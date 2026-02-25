package dsco

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// BaseURLProd 生产环境基础地址。
	BaseURLProd = "https://api.dsco.io/api/v3"
	// BaseURLStaging 沙箱环境基础地址。
	BaseURLStaging = "https://staging-api.dsco.io/api/v3"
)

// Config 表示 Dsco SDK 客户端配置。
//
// 默认行为：使用静态 bearer token（即每次请求都带 Authorization: bearer <Token>）。
// 同时也支持 OAuth2：通过 c.OAuth2.GetAccessToken 手动获取 token（后续可扩展自动刷新）。
type Config struct {
	// BaseURL API 基础地址，默认为 BaseURLProd。
	BaseURL string

	// HTTPClient 自定义 http 客户端（可用于设置超时、代理等），默认 10 秒超时。
	HTTPClient *http.Client

	// Token 静态 bearer token（推荐默认用法）。
	Token string

	// UserAgent 可选 User-Agent。
	UserAgent string
}

// Client 是 Dsco API 客户端。
type Client struct {
	baseURL   *url.URL
	httpCli   *http.Client
	userAgent string

	token string

	// 服务分组（按 OpenAPI tag/业务域）。
	Order     *OrderService
	Inventory *InventoryService
	Invoice   *InvoiceService
	Return    *ReturnService
	Cancel    *CancelService
	Shipment  *ShipmentService
	Stream    *StreamService
	OAuth2    *OAuth2Service
}

// New 创建一个新的 Dsco 客户端。
func New(cfg Config) (*Client, error) {
	base := strings.TrimSpace(cfg.BaseURL)
	if base == "" {
		base = BaseURLProd
	}
	u, err := url.Parse(base)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "" || u.Host == "" {
		return nil, errors.New("BaseURL 非法：缺少 scheme/host")
	}
	if u.Path == "" {
		u.Path = "/"
	}

	httpCli := cfg.HTTPClient
	if httpCli == nil {
		httpCli = &http.Client{Timeout: 10 * time.Second}
	}

	c := &Client{
		baseURL:   u,
		httpCli:   httpCli,
		userAgent: strings.TrimSpace(cfg.UserAgent),
		token:     strings.TrimSpace(cfg.Token),
	}

	c.Order = &OrderService{c: c}
	c.Inventory = &InventoryService{c: c}
	c.Invoice = &InvoiceService{c: c}
	c.Return = &ReturnService{c: c}
	c.Cancel = &CancelService{c: c}
	c.Shipment = &ShipmentService{c: c}
	c.Stream = &StreamService{c: c}
	c.OAuth2 = &OAuth2Service{c: c}

	return c, nil
}
