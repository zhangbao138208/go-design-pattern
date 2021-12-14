package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main()  {
	resp,err := http.Get("http://127.0.0.1:8089")
	if err != nil {
		log.Fatal("client get error :",err)
	}
	defer resp.Body.Close()
	//
	fmt.Println(resp.Header)
	fmt.Println(resp.Status)
    buf := make([]byte,1024)
	for true {
		n,err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}  else {
			fmt.Println("读取完毕")
			res := string(buf[:n])
			fmt.Println(res)
		}
		break
	}
}