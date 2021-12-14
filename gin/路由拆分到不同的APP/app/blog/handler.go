package blog

import (
	"fmt"
	"gin-demo1/app/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {

}

func (Handler)PostHandler(c *gin.Context)  {
	pr := new(model.PostRequest)
	err := c.ShouldBindJSON(pr)
	if err != nil {
		fmt.Print("postHandler ShouldBindJSON err:",err)
		c.JSON(http.StatusBadRequest,gin.H{
			"errors":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
			"status":"ok",
			"data": *pr,
	})
}

func (Handler)CommentHandler(c *gin.Context)  {

	c.JSON(http.StatusOK,nil)
}

func (Handler)XmlHandler(c *gin.Context)  {
	c.XML(http.StatusOK,gin.H{
		"data":"ddd",
		"status":200,
	})
}

func (Handler)YamlHandler(c *gin.Context)  {
	c.YAML(http.StatusOK,gin.H{
		"data":"ddd",
		"status":200,
	})
}

