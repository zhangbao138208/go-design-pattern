package middleWare

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		context.Set("request","中间件")
		context.Next()
		status := context.Writer.Status()
		fmt.Println("中间件执行完毕",status)
		fmt.Println("time:",time.Since(t))
	}
}

func AuthMiddleware() gin.HandlerFunc  {
	return func(c *gin.Context) {
		if cookie,err := c.Cookie("abc");err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized,gin.H{"error":"StatusUnauthorized"})
		c.Abort()
	}
}
