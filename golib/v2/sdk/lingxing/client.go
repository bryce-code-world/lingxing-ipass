package lingxing

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	// BaseURLProd 生产环境基础地址。
	BaseURLProd = "https://openapi.lingxing.com"
)

// Config 表示领星 SDK 客户端配置。
type Config struct {
	// BaseURL API 基础地址，默认 BaseURLProd。
	BaseURL string

	// HTTPClient 自定义 http 客户端（可用于设置超时、代理等），默认 10 秒超时。
	HTTPClient *http.Client

	// AppID 即 app_key（文档中的 AppID）。
	AppID string

	// AppSecret 获取/续约 token 时使用（multipart/form-data）。
	AppSecret string

	// AccessToken 业务接口调用使用（Query Param：access_token）。
	AccessToken string

	// AutoToken 表示是否在业务请求前自动获取/刷新 access_token。
	// 开启后，业务请求无需显式传入 AccessToken（但必须配置 AppID + AppSecret）。
	AutoToken bool

	// TokenLeeway 表示提前刷新 token 的时间窗口（例如 30s 表示将在过期前 30 秒刷新）。
	// 仅在 AutoToken=true 时生效。默认 30 秒。
	TokenLeeway time.Duration

	// UserAgent 可选 User-Agent。
	UserAgent string

	// Now 用于生成 timestamp（Unix 秒）。不传则使用 time.Now。
	Now func() time.Time
}

// Client 是领星 API 客户端。
type Client struct {
	baseURL   *url.URL
	httpCli   *http.Client
	userAgent string

	appID       string
	appSecret   string
	accessToken string

	now func() time.Time

	autoToken   bool
	tokenLeeway time.Duration

	tokenMu       sync.Mutex
	refreshToken  string
	tokenExpireAt time.Time

	// 服务分组（按业务域）。
	Order         *OrderService
	Logistics     *LogisticsService
	Inventory     *InventoryService
	Warehouse     *WarehouseService
	Shop          *ShopService
	Config        *ConfigService
	Authorization *AuthorizationService
}

// New 创建一个新的领星客户端。
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

	nowFn := cfg.Now
	if nowFn == nil {
		nowFn = time.Now
	}

	c := &Client{
		baseURL:     u,
		httpCli:     httpCli,
		userAgent:   strings.TrimSpace(cfg.UserAgent),
		appID:       strings.TrimSpace(cfg.AppID),
		appSecret:   strings.TrimSpace(cfg.AppSecret),
		accessToken: strings.TrimSpace(cfg.AccessToken),
		now:         nowFn,
		autoToken:   cfg.AutoToken,
	}

	leeway := cfg.TokenLeeway
	if leeway <= 0 {
		leeway = 30 * time.Second
	}
	c.tokenLeeway = leeway

	c.Order = &OrderService{c: c}
	c.Logistics = &LogisticsService{c: c}
	c.Inventory = &InventoryService{c: c}
	c.Warehouse = &WarehouseService{c: c}
	c.Shop = &ShopService{c: c}
	c.Config = &ConfigService{c: c}
	c.Authorization = &AuthorizationService{c: c}

	return c, nil
}
