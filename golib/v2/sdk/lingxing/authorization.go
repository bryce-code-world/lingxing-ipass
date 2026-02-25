package lingxing

import (
	"context"
	"net/url"
)

// AuthorizationService 提供 access_token 获取与刷新接口。
type AuthorizationService struct {
	c *Client
}

// GetToken 获取 access_token / refresh_token。
//
// API Path: /api/auth-server/oauth/access-token
func (s *AuthorizationService) GetToken(ctx context.Context) (Token, error) {
	if s.c == nil {
		return Token{}, ErrMissingAppID
	}
	if s.c.appID == "" {
		return Token{}, ErrMissingAppID
	}
	if s.c.appSecret == "" {
		return Token{}, ErrMissingAppSecret
	}

	var out Token
	if err := s.c.doMultipartForm(ctx, "/api/auth-server/oauth/access-token", url.Values{
		"appId":     []string{s.c.appID},
		"appSecret": []string{s.c.appSecret},
	}, &out); err != nil {
		return Token{}, err
	}
	return out, nil
}

// RefreshToken 使用 refresh_token 刷新 access_token。
//
// 注意：文档说明每个 refresh_token 只能使用一次，因此刷新成功后需要保存新的 refresh_token。
//
// API Path: /api/auth-server/oauth/refresh
func (s *AuthorizationService) RefreshToken(ctx context.Context, refreshToken string) (Token, error) {
	if s.c == nil {
		return Token{}, ErrMissingAppID
	}
	if s.c.appID == "" {
		return Token{}, ErrMissingAppID
	}
	if refreshToken == "" {
		// refresh_token 不存在时，调用方应走 GetToken。
		return Token{}, ErrMissingAccessToken
	}

	var out Token
	if err := s.c.doMultipartForm(ctx, "/api/auth-server/oauth/refresh", url.Values{
		"appId":        []string{s.c.appID},
		"refreshToken": []string{refreshToken},
	}, &out); err != nil {
		return Token{}, err
	}
	return out, nil
}
