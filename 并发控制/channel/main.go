package main

import (
	"fmt"
	"sync"
	"time"
)

func Process(ch chan int)  {
	time.Sleep(time.Millisecond*10)
	fmt.Println("测试-----")
	ch<-1
}

func main()  {
	wg := sync.WaitGroup{}

	wg.Add(1)
	chs := make([]chan int,10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Process(chs[i])
	}
	for _, ch := range chs {
		<-ch
	}
	fmt.Println("main exit")
}

