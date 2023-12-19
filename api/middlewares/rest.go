package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set the necessary CORS headers
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")

		// Handle the OPTIONS method directly in Gin's built-in OPTIONS request handler
		if c.Request.Method == http.MethodOptions {
			c.Writer.WriteHeader(200)
			return
		}

		c.Next()
	}
}
