package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:4577")
	if err != nil {
		log.Fatal("net listen error:",err)
	}
	defer listen.Close()
	log.Println("server start success! on port:4577")
	for  {
		conn,err := listen.Accept()
		if err != nil {
			log.Println("listen accept error :",err)
			continue
		}
		go process(conn)

	}
}

func process(conn net.Conn)  {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for true {
		msg,err := Decode(reader)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Printf("Decode error:%v\n",err)
			return
		}
		fmt.Println(time.Now(),msg)
		conn.Write([]byte(msg))
	}
}

func Decode(reader *bufio.Reader) (string, error) {
	lengthByte, err := reader.Peek(4)
	if err != nil {
		fmt.Println("peek error:", err)
		return "", err
	}
	lengthBuf := bytes.NewBuffer(lengthByte)
	var length int32
	err = binary.Read(lengthBuf, binary.LittleEndian, &length)
	if err != nil {
		fmt.Println("binary Read error:", err)
		return "", err
	}
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	var readBuf = make([]byte, length+4)
	_, err = reader.Read(readBuf)
	if err != nil {
		fmt.Println("reader read error:", err)
		return "", err
	}
	return string(readBuf[4:]), nil
}