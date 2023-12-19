package middlewares

import (
	"net/http"

	"srating/x/rest"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var code int
				if httpError, ok := err.(rest.IHttpError); ok {
					code = httpError.StatusCode()
				} else {
					code = http.StatusInternalServerError
				}
				response := rest.Response{
					Status:  "error",
					Message: err.(error).Error(),
				}
				c.JSON(code, response)

			}
		}()
		c.Next()
	}
}
