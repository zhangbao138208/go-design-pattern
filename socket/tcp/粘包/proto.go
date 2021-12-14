package protos

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

func Decode(reader bufio.Reader) (string, error) {
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
