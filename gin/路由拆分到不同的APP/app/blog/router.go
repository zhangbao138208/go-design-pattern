package blog

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

func Router(e *gin.Engine)  {
	b := e.Group("blog")
	{
		//b.POST("post",Handler{}.postHandler)
		//b.POST("comment",Handler{}.commentHandler)
		//b.GET("xml",Handler{}.xmlHandler)
		//b.GET("yaml",Handler{}.yamlHandler)
		ht := reflect.TypeOf(Handler{})
		hv := reflect.ValueOf(Handler{})
		for i := 0; i < ht.NumMethod(); i++ {
			//fmt.Println("method name :",ht.Method(i).Name,strings.TrimRight(ht.Method(i).Name,"Handler"))
			b.GET(strings.ToLower(ht.Method(i).Name[:len(ht.Method(i).Name)-7]),hv.Method(i).Interface().(func(*gin.Context)))
		}
	}
}
