package middlewares

import (
	"net/http"

	"srating/x/rest"

	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		JwtAuthMiddleware(secret)(c)
		userRole, _ := c.Get("x-user-role")
		if role, ok := userRole.(string); !ok || role != "admin" {
			response := rest.Response{
				Status:  "error",
				Message: "You are not admin",
			}
			c.JSON(http.StatusForbidden, response)
			c.Abort()
			return
		}
		c.Next()
	}
}
