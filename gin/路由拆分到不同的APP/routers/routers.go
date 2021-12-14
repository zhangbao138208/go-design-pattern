package routers

import "github.com/gin-gonic/gin"

type Option func(e *gin.Engine)

var options []Option

func Includes(op ...Option) {
	options = op
}

func Init() *gin.Engine {
	r := gin.Default()
	for _, option := range options {
		option(r)
	}
	return r
}