package middleware

import (
	"github.com/gin-gonic/gin"
)

func AccessMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// // middleware
		// if c.Request.Header.Get("authorization") == "" {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, dto.DefaultHttpResponse{StatusCode: http.StatusUnauthorized, Message: "unauthorized", Data: map[string]interface{}{"data": nil}})
		// 	return
		// } else {
		// 	fmt.Println("ok")
		c.Next()
		// }
	}
}
