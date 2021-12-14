package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	listener ,err := net.ListenUDP("udp",&net.UDPAddr{
		 IP: net.IPv4(0,0,0,0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("udp listen err:",err)
		return
	}
	defer listener.Close()
	for  {
		var buf [1024]byte
		n,addr ,err := listener.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("read from udp err:",err)
			continue
		}
		fmt.Printf("udp read 【n=%v,addr=%v】\n",n,addr)
		fmt.Printf("%v,%s\n",time.Now(),buf[:n])
		_,err = listener.WriteToUDP([]byte(fmt.Sprintf("%v----------,%s",time.Now(),buf[:n])),addr)
		if err != nil {
			fmt.Println("udp server write err:\n",err)
		}
	}

}
