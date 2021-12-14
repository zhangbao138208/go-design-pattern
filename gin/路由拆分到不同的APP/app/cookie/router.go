package cookie

import (
	"gin-demo1/app/middleWare"
	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine)  {
	c := e.Group("/cookies")
	{
		c.GET("/login",loginHandler)
		c.GET("/home",middleWare.AuthMiddleware(),homeHandler)
	}
}
