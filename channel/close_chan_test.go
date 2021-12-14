package channel

import (
	"fmt"
	"testing"
)

func TestCloseChan(t *testing.T) {
	var ch1 chan int = make(chan int)
	var ch2 chan int = make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for true {
			if c,ok := <-ch1;ok {
				ch2 <- c*c
			}else {
				break
			}
			//ch2 <- (<-ch1*10)
		}
		close(ch2)
	}()
	//for c2 := range ch2 {
	//	fmt.Println(c2)
	//}
	for  {
		fmt.Println(<-ch2)
	}
}

func TestNilChan(t *testing.T) {
	var ch chan int
	fmt.Println(ch==nil)
	ch<-1
}
