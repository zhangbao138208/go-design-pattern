package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main()  {
	client,err := net.Dial("tcp","127.0.0.1:4577")
	if err != nil {
		log.Fatal("client dial error:",err)
	}
	log.Println("client start ....")
	for true {
		var input string
		//var inputBuf = [256]byte{}

		readerIn := bufio.NewReader(os.Stdin)
		input,_=readerIn.ReadString('\n')

		if err != nil {
			log.Println("scan error:",err)
			continue
		}
		inputStr := strings.Trim(input,"\r\n")
		pkg ,_:= Encode(inputStr)
		for i := 0; i < 30; i++ {
			_,err = client.Write(pkg)
			if err != nil {
				log.Println("client write error:",err)
				continue
			}
		}
		go func() {
			for true {
				var buf [1024]byte
				n,err :=client.Read(buf[:])
				if err != nil {
					fmt.Println("client read error:",err)
					continue
				}
				log.Printf("client read %s\n",buf[:n])
			}
		}()
	}
}
func Encode(message string) ([]byte, error) {
	length := int32(len(message))
	var pkg = new(bytes.Buffer)
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		fmt.Println("binary Write error:", err)
		return nil, err
	}
	// 写入实体消息
	err = binary.Write(pkg,binary.LittleEndian,[]byte(message))
	if err != nil {
		fmt.Println("实体消息 binary Write error:", err)
		return nil, err
	}
	return pkg.Bytes(),nil
}

