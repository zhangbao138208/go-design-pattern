package main

import (
	"fmt"
	"runtime"
)

//出CPU时间片，重新等待安排任务(大概意思就是本来计划的好好的周末出去烧烤，
//但是你妈让你去相亲,两种情况第一就是你相亲速度非常快，见面就黄不耽误你继续烧烤，
//第二种情况就是你相亲速度特别慢，见面就是你侬我侬的，耽误了烧烤，但是还馋就是耽误了烧烤你还得去烧烤)
func main() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}

//func main() {
//	f,err := os.Create("./trace.out")
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//
//	err = trace.Start(f)
//	if err != nil {
//		panic(err)
//	}
//	defer trace.Stop()
//	fmt.Println("Hello World!")
//}
