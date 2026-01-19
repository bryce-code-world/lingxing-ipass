package adminweb

import (
	"io"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *Server) handleStatic(c *gin.Context) {
	name := strings.TrimSpace(c.Param("name"))
	if name == "" {
		c.Status(http.StatusNotFound)
		return
	}
	name = path.Base(name)
	switch name {
	case "admin.css", "admin.js":
	default:
		c.Status(http.StatusNotFound)
		return
	}
	b, err := assetsFS.ReadFile("static/" + name)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	if strings.HasSuffix(name, ".css") {
		c.Header("Content-Type", "text/css; charset=utf-8")
	} else if strings.HasSuffix(name, ".js") {
		c.Header("Content-Type", "application/javascript; charset=utf-8")
	}
	_, _ = io.WriteString(c.Writer, string(b))
}
