package concurrency

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)
var x int64
var wg sync.WaitGroup
var lock sync.Mutex
var rwLock sync.RWMutex
func TestLock(t *testing.T) {
	var start,end time.Time
	start = time.Now()
	wg.Add(100)
	for i := 0; i < 100; i++ {
		//go add()
		go atomicAdd()
	}
	wg.Wait()
	end = time.Now()
	fmt.Println(x,end.Sub(start))
	x = 0
	start = time.Now()
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go add()
		//go atomicAdd()
	}
	wg.Wait()
	end = time.Now()
	fmt.Println(x,end.Sub(start))


}
func add() {
	for i := 0; i < 5000; i++ {
		//lock.Lock()
		rwLock.Lock()
		x = x + 1
		//lock.Unlock()
		rwLock.Unlock()
	}
	wg.Done()
}

func atomicAdd()  {
	for i := 0; i < 5000; i++ {
		atomic.AddInt64(&x,int64(1))
	}
	wg.Done()
}

var m = make(map[string]int)

func set(k string,v int)  {
	rwLock.Lock()
	m[k] = v
	rwLock.Unlock()
}
func get(k string) int {
	defer rwLock.RUnlock()
	rwLock.RLock()
	return m[k]

}

var s = make([]string,0)

func setSlice(v string)  {
	s = append(s,v)
}
func getSlice(n int)string  {
	return s[n]
}


func TestConcurrentMap(t *testing.T) {
	var num = 2
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(n int) {
			defer wg.Done()
			set(strconv.Itoa(n),n)
			fmt.Printf("map[%v]=%v\n",n,get(strconv.Itoa(n)))
		}(i)
	}
	wg.Wait()
}
func TestConcurrentSlice(t *testing.T) {
	var num = 1000
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(n int) {

			defer wg.Done()
			setSlice(strconv.Itoa(n))
			fmt.Printf("s[%v]=%v\n",n,getSlice(0))
		}(i)
	}
	wg.Wait()
	fmt.Println("s len:", len(s))
}
func TestSyncMap(t *testing.T) {
	var num = 100
	wg := sync.WaitGroup{}
	wg.Add(num)
	sm := sync.Map{}
	for i := 0; i < num; i++ {
		go func(n int) {
			defer wg.Done()
			sm.Store(n,n*n)
			v,_:=sm.Load(n)
			fmt.Printf("map[%v]=%v\n",n,v)
		}(i)
	}
	wg.Wait()
}

type priority int

const (
	Log_emergey priority = iota
	log_fatla
	log_err =iota
	log_info
)

func TestIota(t *testing.T) {
	fmt.Println(log_info,log_err)
	fmt.Println(mutexLocked,mutexWoken,mutexStarving,mutexWaiterShift,starvationThresholdNs)
}
const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
	starvationThresholdNs = 1e6
)
func AddElement(slice []int, e int) []int {
	return append(slice, e)
}

func TestSlice(t *testing.T) {

	var array [10]int

	var slice1 = array[5:6]

	fmt.Println("lenth of slice: ", len(slice1))
	fmt.Println("capacity of slice: ", cap(slice1))
	fmt.Println(&slice1[0] == &array[5])
	var slice []int = make([]int,0,4)
	slice = append(slice, 1, 2, 3)
	fmt.Println(cap(slice))
	newSlice := AddElement(slice, 4)
	fmt.Println(&slice[0] == &newSlice[0])
}
func TestSlice2(t *testing.T) {
	var arr = [10]int{1,2,3,4,5,6,7,8,9,10}
	var s1 = arr[5:7]
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s1[0],s1[1])
	s1 = append(s1,11)
	s1 = append(s1,12,13,14,15)
	fmt.Println(arr,&arr[5]==&s1[0])
}

func TestDefer(t *testing.T)  {
	fmt.Println(deferFuncReturn())
	fmt.Println(deferFuncReturn2())
}
func deferFuncReturn() (result int) {
	i := 1

	defer func() {
		result++
	}()

	return i
}
func deferFuncReturn2()  int {
	i := 1

	defer func() {
		i++
	}()

	return i
}
func TestDeferPanic(t *testing.T) {
	//Dived(0)
	defer IsPanic()
	DivedWithIf(0)
}
func NoPanic() {
	if err := recover(); err != nil {
		fmt.Println("Recover success...:ERR",err)
	}
}
func IsPanic() bool {
	if err := recover(); err != nil {
		fmt.Println("Recover success...:ERR",err)
		return true
	}
	return false
}

func DivedWithIf(n int) {
	//defer func() {
	//     defer IsPanic()
	//}()

	fmt.Println(1/n)
}
func Dived(n int) {
	defer NoPanic()

	fmt.Println(1/n)
}

func TestChannel(t *testing.T) {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go func() {
		close(ch1)
	}()
	go func() {
		//close(ch2)
	}()

	select {
	case c1,ok :=<-ch1:
		fmt.Println("c1=",c1,ok)
	case c2,ok :=<-ch2:
		fmt.Println("c2=",c2,ok)
	}
	fmt.Println("main exit")
}

func TestRange(t *testing.T) {
	v := []int{1, 2, 3}
	for i:= range v {
		v = append(v, i)
	}
	fmt.Println(v)
}