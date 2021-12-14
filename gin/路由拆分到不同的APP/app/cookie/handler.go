package cookie

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func loginHandler(c *gin.Context)  {
	c.SetCookie("abc","123",30,"/","localhost",false,true)
	c.String(http.StatusOK,"login success")
}

func homeHandler(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{
		"data":"home",
	})
}