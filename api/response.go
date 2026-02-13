package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func ok(c *gin.Context, data any) {
	c.JSON(http.StatusOK, resp{Code: 0, Message: "ok", Data: data})
}

func fail(c *gin.Context, httpStatus int, code int, msg string) {
	c.JSON(httpStatus, resp{Code: code, Message: msg})
}
