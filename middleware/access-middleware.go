package middleware

import (
	"company-config-go/delivery/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AccessMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// middleware
		if c.Request.Header.Get("authorization") == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.DefaultHttpResponse{StatusCode: http.StatusUnauthorized, Message: "unauthorized", Data: map[string]interface{}{"data": nil}})
			return
		} else {
			fmt.Println("ok")
			c.Next()
		}
	}
}
