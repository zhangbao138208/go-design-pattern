package main
// import this library
import (
	"bufio"
	"fmt"
	"github.com/shiena/ansicolor"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"strings"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
		FullTimestamp: true,

	})
	log.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
	log.SetLevel(log.DebugLevel)
}
func main()  {
	log.Info("\x1b[95m够浪客户端正在启动...")
	client,err := net.Dial("tcp","127.0.0.1:8011")
	if err != nil {
		log.Errorf("够浪客户端启动失败:\x1b[91m%v\n",err)
		return
	}
	defer client.Close()
	log.Info("\x1b[95m够浪客户端启动成功！")
	log.Info("\x1b[95m请输入注册账号:")

	var  registerName string
	fmt.Scan(&registerName)
	registerName = strings.Trim(registerName,"\n")

	// 发送注册名
	_,err = client.Write([]byte(registerName))
	if err != nil {
		log.Errorf("注册失败:\x1b[91m%v\n",err)
		return
	}
	// 读取注册成功信息
	var rBuf [128]byte
	n,err := client.Read(rBuf[:])
	if err != nil {
		log.Errorf("读取注册信息失败:\x1b[91m%v\n",err)
		return
	}
	log.Infof("\x1b%s\n",rBuf[:n])
	log.Info("\x1b[93m⚠温馨提示长时间不在线会强制下线！")
	

	go func() {
		// 处理收到的消息
		for  {
			messBuf  := [1024]byte{}
			n,err := client.Read(messBuf[:])
			if err != nil {
				log.Errorf("处理收到的消息失败:\x1b[91m%v\n",err)
				return
			}
			log.Infof("\x1b[92m%s\n",messBuf[:n])
		}
	}()
	// 发送消息
	for  {
		reader := bufio.NewReader(os.Stdin)
		input,err := reader.ReadString('\n')
		if err != nil {
			log.Errorf("reader ReadString error:\x1b[91m%v\n",err)
			continue
		}
		inputStr := strings.Trim(string(input),"\r\n")
		if inputStr == "Q" {
			return
		}
		_,err = client.Write([]byte(inputStr))
		if err != nil {
			log.Errorf("client Write error:\x1b[91m%v\n",err)
			continue
		}
	}
}