package shop

import (
	"gin-demo1/app/middleWare"
	"github.com/gin-gonic/gin"
	"github.com/kubastick/ginception"
)

func Router(e *gin.Engine)  {
	e.LoadHTMLGlob("www/**/*")
	s := e.Group("/shop")
	{
		s.GET("good/:name",goodHandler)
		s.GET("checkout",checkoutHandler)
		s.Use(middleWare.MiddleWare()).GET("index",indexHandler)
		s.Use(middleWare.MiddleWare()).GET("index.html",indexHtmlHandler)
		s.GET("async",longAsyncHandler)
		s.GET("sync",longSyncHandler)
		s.GET("asyncWithoutCopy",longAsyncWithoutCopyHandler)
		s.GET("panic",ginception.Middleware(),PanicHandler)
	}
}
