package main

import (
	"gin-demo1/app/blog"
	"gin-demo1/app/cookie"
	"gin-demo1/app/shop"
	"gin-demo1/routers"
	"log"
)

func main()  {
	routers.Includes(blog.Router,shop.Router,cookie.Router)
	r := routers.Init()
	if err := r.Run(":8082"); err != nil {
		log.Fatal("start error:",err)
	}
}
