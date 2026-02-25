package lingxing

import (
	"context"
	"net/url"
)

func (s *AuthorizationService) GetTokenWithRawBody(ctx context.Context) (Token, string, error) {
	if s.c == nil {
		return Token{}, "", ErrMissingAppID
	}
	if s.c.appID == "" {
		return Token{}, "", ErrMissingAppID
	}
	if s.c.appSecret == "" {
		return Token{}, "", ErrMissingAppSecret
	}

	var out Token
	_, raw, err := s.c.doMultipartFormWithEnvelopeWithRawBody(ctx, "/api/auth-server/oauth/access-token", url.Values{
		"appId":     []string{s.c.appID},
		"appSecret": []string{s.c.appSecret},
	}, &out)
	if err != nil {
		return Token{}, raw, err
	}
	return out, raw, nil
}

func (s *AuthorizationService) RefreshTokenWithRawBody(ctx context.Context, refreshToken string) (Token, string, error) {
	if s.c == nil {
		return Token{}, "", ErrMissingAppID
	}
	if s.c.appID == "" {
		return Token{}, "", ErrMissingAppID
	}
	if refreshToken == "" {
		return Token{}, "", ErrMissingAccessToken
	}

	var out Token
	_, raw, err := s.c.doMultipartFormWithEnvelopeWithRawBody(ctx, "/api/auth-server/oauth/refresh", url.Values{
		"appId":        []string{s.c.appID},
		"refreshToken": []string{refreshToken},
	}, &out)
	if err != nil {
		return Token{}, raw, err
	}
	return out, raw, nil
}
