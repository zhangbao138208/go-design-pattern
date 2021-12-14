package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)
func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
func TestSetOne(t *testing.T) {
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}
func TestSetTwo(t *testing.T) {
	runtime.GOMAXPROCS(3)
	go a()
	go b()
	time.Sleep(time.Second)
}

