package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main()  {
   listener,err := net.Listen("tcp","127.0.0.1:8089")
	if err != nil {
		log.Fatal("listen error:",err.Error())
	}
    defer listener.Close()
	for  {
		conn,err := listener.Accept()
		if err != nil {
			fmt.Errorf("listener accept error %s\n",err)
			continue
		}
        go process(conn)
	}
}

func process(conn net.Conn)  {
	defer conn.Close()
	for  {
		reader := bufio.NewReader(conn)
		var bytes [128]byte
		n,err := reader.Read(bytes[:])
		if err != nil {
           fmt.Println("reader read error:",err)
			return
		}
		readStr := string(bytes[:n])
		fmt.Println(time.Now(),"read string:",readStr)
		_,err =conn.Write([]byte(fmt.Sprint("server receiver ",readStr)))
		if err != nil {
			fmt.Println("server write error:",err)
		}
	}
}

