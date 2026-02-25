package dsco

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

// OAuth2Service 封装 OAuth2 相关接口。
type OAuth2Service struct {
	c *Client
}

// GetAccessToken 调用 /oauth2/token 获取 access_token。
//
// 注意：该接口不需要 Authorization header，且请求体为 form-urlencoded。
func (s *OAuth2Service) GetAccessToken(ctx context.Context, req OAuth2TokenRequest) (*AccessTokenResponse, error) {
	if strings.TrimSpace(req.ClientID) == "" || strings.TrimSpace(req.ClientSecret) == "" {
		return nil, errors.New("client_id/client_secret 不能为空")
	}
	grantType := strings.TrimSpace(req.GrantType)
	if grantType == "" {
		grantType = "client_credentials"
	}

	values := url.Values{}
	values.Set("client_id", req.ClientID)
	values.Set("client_secret", req.ClientSecret)
	values.Set("grant_type", grantType)
	if strings.TrimSpace(req.Scope) != "" {
		values.Set("scope", req.Scope)
	}

	var out AccessTokenResponse
	if err := s.c.doForm(ctx, http.MethodPost, "/oauth2/token", values, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
