package lingxing

import (
	"context"
	"time"
)

func (c *Client) ensureAccessToken(ctx context.Context) error {
	if !c.autoToken {
		return nil
	}
	if c.appID == "" {
		return ErrMissingAppID
	}
	if c.appSecret == "" && c.refreshToken == "" && c.accessToken == "" {
		return ErrMissingAppSecret
	}

	c.tokenMu.Lock()
	defer c.tokenMu.Unlock()

	now := c.now()
	if c.accessToken != "" && !c.shouldRefreshLocked(now) {
		return nil
	}

	// 优先用 refresh_token 刷新；失败则降级为重新获取 token。
	if c.refreshToken != "" {
		tok, err := c.Authorization.RefreshToken(ctx, c.refreshToken)
		if err == nil && tok.AccessToken != "" {
			c.setTokenLocked(now, tok)
			return nil
		}
	}

	if c.appSecret == "" {
		return ErrMissingAppSecret
	}
	tok, err := c.Authorization.GetToken(ctx)
	if err != nil {
		return err
	}
	c.setTokenLocked(now, tok)
	return nil
}

func (c *Client) shouldRefreshLocked(now time.Time) bool {
	if c.tokenExpireAt.IsZero() {
		// 未知过期时间：不自动刷新，避免频繁刷新；由调用方按需设置初始 token。
		return false
	}
	return now.Add(c.tokenLeeway).After(c.tokenExpireAt)
}

func (c *Client) setTokenLocked(now time.Time, tok Token) {
	if tok.AccessToken != "" {
		c.accessToken = tok.AccessToken
	}
	if tok.RefreshToken != "" {
		c.refreshToken = tok.RefreshToken
	}
	if tok.ExpiresIn > 0 {
		c.tokenExpireAt = now.Add(time.Duration(tok.ExpiresIn) * time.Second)
	} else {
		c.tokenExpireAt = time.Time{}
	}
}
