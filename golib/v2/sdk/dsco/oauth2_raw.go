package dsco

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

func (s *OAuth2Service) GetAccessTokenWithRawBody(ctx context.Context, req OAuth2TokenRequest) (*AccessTokenResponse, string, error) {
	if s == nil || s.c == nil {
		return nil, "", errors.New("dsco client 不能为 nil")
	}
	if strings.TrimSpace(req.ClientID) == "" || strings.TrimSpace(req.ClientSecret) == "" {
		return nil, "", errors.New("client_id/client_secret 不能为空")
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
	raw, err := s.c.doFormWithRawBody(ctx, http.MethodPost, "/oauth2/token", values, &out)
	if err != nil {
		return nil, raw, err
	}
	return &out, raw, nil
}
