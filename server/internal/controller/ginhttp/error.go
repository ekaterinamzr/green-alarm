package ginhttp

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message"`
}

func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, response{msg})
}
