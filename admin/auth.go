package admin

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const sessionCookieName = "ipass_admin_session"

func (s *Server) sessionToken() string {
	sum := sha256.Sum256([]byte("ipass:" + s.env.Admin.Password))
	return hex.EncodeToString(sum[:])
}

func (s *Server) requireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(sessionCookieName)
		if err != nil || strings.TrimSpace(cookie) == "" || cookie != s.sessionToken() {
			c.Redirect(http.StatusFound, "/admin/login")
			c.Abort()
			return
		}
		c.Next()
	}
}

func (s *Server) setSession(c *gin.Context) {
	c.SetCookie(sessionCookieName, s.sessionToken(), 86400*7, "/admin", "", false, true)
}

func (s *Server) clearSession(c *gin.Context) {
	c.SetCookie(sessionCookieName, "", -1, "/admin", "", false, true)
}
