package adminweb

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const adminCookieName = "ipass_admin_session"

type session struct {
	TS    int64
	Nonce string
	Sig   string
}

func (s *Server) requireAdminForUI() gin.HandlerFunc {
	return func(c *gin.Context) {
		if s.isAuthed(c) {
			c.Next()
			return
		}
		next := c.Request.URL.Path
		if q := strings.TrimSpace(c.Request.URL.RawQuery); q != "" {
			next += "?" + q
		}
		c.Redirect(http.StatusFound, "/admin/ui/login?next="+urlQueryEscape(next))
		c.Abort()
	}
}

func (s *Server) requireAdminForAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		if s.isAuthed(c) {
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
	}
}

func (s *Server) isAuthed(c *gin.Context) bool {
	pass := strings.TrimSpace(s.getAdminPassword())
	if pass == "" {
		// 未配置密码：一期允许直接访问（便于本地），但线上必须配置。
		return true
	}

	// 兼容脚本/CI：允许用请求头携带密码。
	if strings.TrimSpace(c.GetHeader("X-Admin-Password")) == pass {
		return true
	}

	// 浏览器：session cookie。
	raw, err := c.Cookie(adminCookieName)
	if err != nil || strings.TrimSpace(raw) == "" {
		return false
	}
	return s.verifyCookie(raw)
}

func (s *Server) signCookie(ts int64, nonce string) string {
	mac := hmac.New(sha256.New, []byte(s.getAdminPassword()))
	_, _ = mac.Write([]byte(strconv.FormatInt(ts, 10)))
	_, _ = mac.Write([]byte("|"))
	_, _ = mac.Write([]byte(nonce))
	return hex.EncodeToString(mac.Sum(nil))
}

func (s *Server) makeCookieValue(now time.Time) (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	nonce := base64.RawURLEncoding.EncodeToString(b)
	ts := now.UTC().Unix()
	sig := s.signCookie(ts, nonce)
	return base64.RawURLEncoding.EncodeToString([]byte(strconv.FormatInt(ts, 10) + "|" + nonce + "|" + sig)), nil
}

func (s *Server) verifyCookie(raw string) bool {
	decoded, err := base64.RawURLEncoding.DecodeString(strings.TrimSpace(raw))
	if err != nil {
		return false
	}
	parts := strings.Split(string(decoded), "|")
	if len(parts) != 3 {
		return false
	}
	tsStr, nonce, sig := parts[0], parts[1], parts[2]
	if tsStr == "" || nonce == "" || sig == "" {
		return false
	}
	ts, err := strconv.ParseInt(tsStr, 10, 64)
	if err != nil {
		return false
	}

	// 一期不要求次数限制；但仍做一个宽松的时间窗校验，避免永久 cookie。
	// 浏览器 session cookie 理论上不持久化，这里额外兜底 7 天。
	now := s.now().UTC().Unix()
	if ts <= 0 || ts > now+60 {
		return false
	}
	if now-ts > int64((7 * 24 * time.Hour).Seconds()) {
		return false
	}

	want := s.signCookie(ts, nonce)
	return hmac.Equal([]byte(want), []byte(sig))
}

func (s *Server) setSessionCookie(c *gin.Context) error {
	pass := strings.TrimSpace(s.getAdminPassword())
	if pass == "" {
		return errors.New("未配置 IPASS_ADMIN_PASSWORD")
	}
	val, err := s.makeCookieValue(s.now())
	if err != nil {
		return err
	}
	c.SetCookie(adminCookieName, val, 0, "/", "", false, true)
	return nil
}

func (s *Server) clearSessionCookie(c *gin.Context) {
	c.SetCookie(adminCookieName, "", -1, "/", "", false, true)
}

func urlQueryEscape(s string) string { return url.QueryEscape(s) }
