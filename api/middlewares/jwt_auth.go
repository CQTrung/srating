package middlewares

import (
	"net/http"
	"strings"

	"srating/utils"
	"srating/x/rest"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		t := strings.Split(tokenString, " ")
		if len(t) != 2 {
			errResponse := rest.Response{
				Status:  "error",
				Message: "Missing authorization header",
			}
			c.JSON(http.StatusUnauthorized, errResponse)
			c.Abort()
			return
		}
		tokenString = t[1]
		authorized, err := utils.IsAuthorized(tokenString, secret)
		if err != nil || !authorized {
			errResponse := rest.Response{
				Status:  "error",
				Message: "Invalid or expired token",
			}
			c.JSON(http.StatusUnauthorized, errResponse)
			c.Abort()
			return
		}

		userID, err := utils.ExtractIDFromToken(tokenString, secret)
		if err != nil {
			errResponse := rest.Response{
				Status:  "error",
				Message: "Can't extract user id",
			}
			c.JSON(http.StatusUnauthorized, errResponse)
			c.Abort()
			return
		}

		role, err := utils.ExtractRoleFromToken(tokenString, secret)
		if err != nil {
			errResponse := rest.Response{
				Status:  "error",
				Message: "Can't extract user role",
			}
			c.JSON(http.StatusUnauthorized, errResponse)
			c.Abort()
			return
		}

		location, err := utils.ExtractLocationFromToken(tokenString, secret)
		if err != nil {
			errResponse := rest.Response{
				Status:  "error",
				Message: "Can't extract user location",
			}
			c.JSON(http.StatusUnauthorized, errResponse)
			c.Abort()
			return
		}

		c.Set("x-user-id", userID)
		c.Set("x-user-role", role)
		c.Set("x-user-location", location)

		c.Next()
	}
}
