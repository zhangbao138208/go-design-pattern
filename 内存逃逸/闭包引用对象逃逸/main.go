package main

import "fmt"

func Fibonacci() func() int {
	a,b:=0,1
	return func() int {
		a,b = b,b+a
		return a
	}
}
func main()  {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
