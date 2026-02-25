package lingxing

// Token 表示领星 OAuth token 响应 data 字段。
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}
