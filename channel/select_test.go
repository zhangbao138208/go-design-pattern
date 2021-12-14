package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestFullChan(t *testing.T) {
	ch := make(chan string,10)

	go func() {
		for s := range ch {
			fmt.Println("res:", s)
			time.Sleep(time.Second)
		}
	}()
	for {
		select {
		// 写数据
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
