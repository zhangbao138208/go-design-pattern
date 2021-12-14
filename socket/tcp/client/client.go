package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main()  {
	dail,err := net.Dial("tcp","127.0.0.1:8089")
	if err != nil {
		log.Fatal("dail error:",err)
	}

	for  {
		reader := bufio.NewReader(os.Stdin)
		inputInfo,err := reader.ReadString('\n')
		if err != nil {
			fmt.Errorf("reader readString error:%s\n",err.Error())
			continue
		}
		inputInfoStr := strings.Trim(inputInfo,"\r\n")
		if inputInfoStr == "Q" {
			return
		}
		for i := 0; i < 20; i++ {
			_,err = dail.Write([]byte(inputInfoStr))
			if err != nil {
				fmt.Println("client write error:",err)
				continue
			}
		}
		var bytes [128]byte
		n,err := dail.Read(bytes[:])
		if err != nil {
			fmt.Println("client read error :",err)
			continue
		}
		fmt.Println(time.Now(),"client readStr【",string(bytes[:n]),"】")
	}
}
