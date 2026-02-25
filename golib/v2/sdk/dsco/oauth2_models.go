package dsco

// OAuth2TokenRequest 表示 OAuth2 client_credentials 请求体（form urlencoded）。
type OAuth2TokenRequest struct {
	// ClientID 对应接口参数 client_id。
	ClientID string
	// ClientSecret 对应接口参数 client_secret。
	ClientSecret string
	// GrantType 对应接口参数 grant_type（当前仅支持 client_credentials）。
	GrantType string
	// Scope 对应接口参数 scope（可选）。
	Scope string
}

// AccessTokenResponse 表示 POST /oauth2/token 的响应体。
type AccessTokenResponse struct {
	// AccessToken 访问令牌（bearer token）。
	AccessToken string `json:"access_token"`
	// TokenType 令牌类型（通常为 Bearer）。
	TokenType string `json:"token_type"`
	// ExpiresIn 过期秒数。
	ExpiresIn int `json:"expires_in"`
	// Scope 返回的 scope（可选）。
	Scope *string `json:"scope,omitempty"`
}
