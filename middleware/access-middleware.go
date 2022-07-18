package middleware

import (
	"company-config-go/delivery/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AccessMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// middleware
		if c.Request.Header.Get("x-endpoint-api-userinfo") == "" && c.Request.Header.Get("x-apigateway-api-userinfo") == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.DefaultHttpResponse{StatusCode: http.StatusUnauthorized, Message: "unauthorized", Data: map[string]interface{}{"data": nil}})
			return
		} else {
			c.Next()
		}
	}
}
