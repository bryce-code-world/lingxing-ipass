package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func bearerAuth(token string) gin.HandlerFunc {
	want := strings.TrimSpace(token)
	return func(c *gin.Context) {
		h := strings.TrimSpace(c.GetHeader("Authorization"))
		if h == "" {
			fail(c, http.StatusUnauthorized, 401, "missing authorization")
			c.Abort()
			return
		}
		parts := strings.SplitN(h, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || strings.TrimSpace(parts[1]) == "" {
			fail(c, http.StatusUnauthorized, 401, "invalid authorization")
			c.Abort()
			return
		}
		got := strings.TrimSpace(parts[1])
		if want == "" || got != want {
			fail(c, http.StatusForbidden, 403, "invalid token")
			c.Abort()
			return
		}
		c.Next()
	}
}
