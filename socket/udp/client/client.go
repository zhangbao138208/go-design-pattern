package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	socket,err := net.Dial("udp","127.0.0.1:30000")
	if err != nil {
		fmt.Println("upd client dail err:",err)
		return
	}
	defer socket.Close()
	for true {
		reader := bufio.NewReader(os.Stdin)
		var buf [128]byte
		input,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("client stdin read err:",err)
			continue
		}
        inputInfo := strings.Trim(string(input),"\r\n")
		if inputInfo == "Q" {
			return
		}
		_,err = socket.Write([]byte(inputInfo))
		if err != nil {
			fmt.Println("client socket write err:",err)
			continue
		}
		n,err  :=socket.Read(buf[:])
		if err != nil {
			fmt.Println("socket read err:",err)
			continue
		}
		fmt.Printf("%s\n",buf[:n])
	}
}
