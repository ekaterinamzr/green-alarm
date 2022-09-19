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

type middleware struct {
	parseToken func(context.Context, string) (id int, role int, err error)
}

func newMiddleware(parseToken func(context.Context, string) (id int, role int, err error)) *middleware{
	return &middleware{parseToken: parseToken}
}

func (m *middleware) checkRole(requiredRole int) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetInt("userRole") > requiredRole {
			errorResponse(c, http.StatusUnauthorized, "permission denied")
			return
		}
		c.Next()
	}
}

func (m *middleware) userIdentity() gin.HandlerFunc {
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

		userId, userRole, err := m.parseToken(c.Request.Context(), headerParts[1])
		if err != nil {
			errorResponse(c, http.StatusUnauthorized, "could not parse token")
			return
		}

		c.Set("userId", userId)
		c.Set("userRole", userRole)
	}
}
