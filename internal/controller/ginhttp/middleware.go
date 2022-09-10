package ginhttp

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
)

func userIdentity(t func(context.Context, string) (id int, role int, err error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authHeader)
		if header == "" {
			errorResponse(c, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			errorResponse(c, http.StatusUnauthorized, "invalid auth header")
			return
		}

		userId, userRole, err := t(c.Request.Context(), headerParts[1])
		if err != nil {
			errorResponse(c, http.StatusUnauthorized, "could not parse token")
			return
		}

		c.Set("userId", userId)
		c.Set("userRole", userRole)
	}
}
