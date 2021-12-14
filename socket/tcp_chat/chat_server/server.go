package main

// import this library
import (
	"fmt"
	"github.com/shiena/ansicolor"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net"
	"os"
	"time"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	log.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
	log.SetLevel(log.DebugLevel)
}

var message chan []byte

var chats chan ChatMessage = make(chan ChatMessage,10)

type UserInfo struct {
	UserName string
	Chats    chan []byte // 用户广播用户消息
	NewUsers chan []byte // 用于广播用户上线下线
}

var onlineUsers = map[string]UserInfo{}

func main() {
	log.Info("\x1b[92m够浪服务器正在启动...")
	listener, err := net.Listen("tcp", "0.0.0.0:8011")
	if err != nil {
		log.Errorf("够浪服务器启动失败:\x1b[91m%v\n", err)
		return
	}
	defer listener.Close()
	log.Info("\x1b[92m够浪服务器启动成功！")
	//无缓冲channel
	message = make(chan []byte)
	go handleMessage()
	go InsertDB()

	for {
		accept, err := listener.Accept()
		if err != nil {
			log.Errorf("listener accept err:\x1b[91m%v\n", err)
			continue
		}
		go HandleConnection(accept)
	}

}

func HandleConnection(accept net.Conn) {
	defer accept.Close()

	log.Infof("\x1b[96m%s已连接到服务器.....\n", accept.RemoteAddr().String())

	var nameBuf [128]byte
	n, err := accept.Read(nameBuf[:])
	if err != nil {
		log.Errorf("accept Read err:\u001B[91m%v\n", err)
		return
	}
	thisUser := nameBuf[:n]
	// 注册在线
	chanChats := make(chan []byte)
	chanNewUsers := make(chan []byte)
	onlineUsers[accept.RemoteAddr().String()] = UserInfo{
		UserName: string(thisUser),
		Chats:    chanChats,
		NewUsers: chanNewUsers,
	}
	log.Infof("用户[%s]注册成功\n", thisUser)
	go handleSelfAcceptMessage(accept)
	// 通知用户上线
	go func() {
		for _, info := range onlineUsers {
			info.NewUsers <- []byte(fmt.Sprintf("????\u001B[96m用户[%s"+"]已加入当前聊天室\n", thisUser))
		}
	}()

	// 处理自身的消息

	//超时chan
	//overtimeChan := make(chan bool)

	// 超时timer
	overTime := time.Duration(time.Second*60)
	overTimer := time.NewTicker(overTime)

	// 处理读到的消息类型
	go func() {
		for {
			var messageBuf [1024]byte
			n, err := accept.Read(messageBuf[:])
			if err != nil {
				delete(onlineUsers, accept.RemoteAddr().String())
				log.Errorf("accept accept read err:\u001B[91m%v\n", err)
				return
			}
			mess := string(messageBuf[:n])
			switch {
			case n == 0:
				// 断开连接
				for _, info := range onlineUsers {
					info.NewUsers <- []byte(fmt.Sprintf("????用户[\x1b[96m%s]已退出当前聊天室\n", thisUser))
				}
				delete(onlineUsers, accept.RemoteAddr().String())
				return
			case mess == "who":
				// 查询在线用户
				m := ""
				for _, info := range onlineUsers {
					m += fmt.Sprintf("\t[\x1b[96m%s]\n", info.UserName)
				}
				onlineUsers[accept.RemoteAddr().String()].Chats <- []byte(fmt.Sprintf("\n%s", m))
				//overtimeChan <- true
				overTimer.Reset(overTime)
				continue

			case len(mess) > 7 && mess[:7] == "rename|":
				onlineUsers[accept.RemoteAddr().String()] = UserInfo{
					UserName: mess[7:],
					Chats:    chanChats,
					NewUsers: chanNewUsers,
				}
				onlineUsers[accept.RemoteAddr().String()].Chats <- []byte("您已成功修改用户!")
				//accept.Write([]byte(fmt.Sprintf("\x1b[97m您已成功修改用户!\n")))
				//overtimeChan <- true
				overTimer.Reset(overTime)
				continue
			}
			// 通知消息
			//for addr, info := range onlineUsers {
			//	if addr != accept.RemoteAddr().String() {
			//		info.Chats <- []byte(message)
			//	}
			//}
			message <- append([]byte("????["+onlineUsers[accept.RemoteAddr().String()].UserName+"]对大家说:"), mess...)
			//overtimeChan <- true
			overTimer.Reset(overTime)
		}
	}()

	// 超时处理
	for {
		select {
		//case <-overtimeChan:
		//case <-time.Tick(time.Second * 60 * 3):
		case <-overTimer.C:

			// 超时自动下线
			for _, info := range onlineUsers {
				info.NewUsers <- []byte(fmt.Sprintf("????用户[\x1b[96m%s]由于长时间未发送消息已被踢出当前聊天室\n", thisUser))
			}
			// 删除在线用户
			delete(onlineUsers, accept.RemoteAddr().String())
			return
		}
	}
}

func handleSelfAcceptMessage(accept net.Conn) {
	for true {
		select {
		case c := <-onlineUsers[accept.RemoteAddr().String()].Chats:
			_, err := accept.Write(c)
			if err != nil {
				log.Errorf("handleSelfAcceptMessage Chats write err:\u001B[91m%v\n", err)
			}
		case n := <-onlineUsers[accept.RemoteAddr().String()].NewUsers:
			_, err := accept.Write(n)
			if err != nil {
				log.Errorf("handleSelfAcceptMessage write NewUsers err:\u001B[91m%v\n", err)
			}
		}

	}
}

type ChatMessage struct {
	gorm.Model
	Address string
	Name    string
	Message string
}

func handleMessage() {
	for true {
		select {
		case msg := <-message:
			for addr, userInfo := range onlineUsers {
				userInfo.Chats <- msg
				// 数据库处理channel
				chats <- ChatMessage{
					Message: string(msg),
					Address: addr,
					Name:    userInfo.UserName,
				}
			}
		}
	}
}
// 入库
func InsertDB()  {
	db,err := gorm.Open(sqlite.Open("chat.db"),&gorm.Config{})
	if err != nil {
		log.Fatal("sqlite.Open error:",err)
	}
	db.AutoMigrate(&ChatMessage{})

	var chatMessages = make([]ChatMessage,0)
	for true {
		select {
		case chat:=<-chats:
			chatMessages = append(chatMessages,chat)
			if len(chatMessages)>=10 {
				go func(cs []ChatMessage) {
					db.CreateInBatches(cs,10)

				}(chatMessages)
				chatMessages = chatMessages[:0]
			}
		}
	}
}
