package shop

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func goodHandler(c *gin.Context)  {
	goodName := c.Param("name")
	c.String(http.StatusOK,fmt.Sprint("goodName=",goodName))
}

func checkoutHandler(c *gin.Context)  {
	productId := c.DefaultQuery("productId","A01")
	c.String(http.StatusOK,fmt.Sprint("productId=",productId))
}

func indexHandler(c *gin.Context)  {
	v,_:= c.Get("request")
	log.Println("request:",v)
	c.HTML(http.StatusOK,"user/index.html",gin.H{
		"title":"测试",
		"address":"中关村",
	})
}
func indexHtmlHandler(c *gin.Context)  {
	c.Redirect(http.StatusMovedPermanently,"/shop/index")
}

func longAsyncHandler(c *gin.Context)  {
	copyContext := c.Copy()
	go func() {
		time.Sleep(3*time.Second)
		log.Println(copyContext.Request.URL.Path)
	}()
}
func longSyncHandler(c *gin.Context)  {
	time.Sleep(3*time.Second)
	log.Println(c.Request.URL.Path)
}
func longAsyncWithoutCopyHandler(c *gin.Context)  {
	go func() {
		time.Sleep(3*time.Second)
		log.Println(c.Request.URL.Path)
	}()
}

func PanicHandler(c *gin.Context)  {
	panic("test panic")
}