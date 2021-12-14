package concurrency

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)
var lock1 sync.Mutex
type ServiceData struct {
	ch chan string
	data []string
}

func (s *ServiceData)Schedule()  {
	for ch := range s.ch {
		s.data = append(s.data,ch)
	}
}

func (s *ServiceData)Close()  {
	close(s.ch)
}
func (s *ServiceData)AddData(d string)  {
	s.ch <- d
}
func NewScheduleJob(size int,done func()) *ServiceData {
	s := &ServiceData{
		ch:make(chan string,size),
		data: make([]string,0),
	}
	go func() {
		s.Schedule()
		done()
	}()
	return s
}
var s11 = make([]string,0)
func TestChanSlice(t *testing.T) {
	var size = 10000
	wg := sync.WaitGroup{}
	var start,end time.Time
	start = time.Now()
	wg.Add(size)
	var ch1 chan struct{} =make(chan struct{})
	s := NewScheduleJob(size, func() {
		ch1<- struct{}{}
	})
	for i := 0; i < size; i++ {
		go func(n int) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				s.AddData(strconv.Itoa(n))
			}
		}(i)
	}
	wg.Wait()
	s.Close()
	<-ch1
	end = time.Now()
	fmt.Printf("s.data length =%v,%v\n", len(s.data),end.Sub(start))


	start = time.Now()
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func(n int) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				lock1.Lock()
				s11 = append(s11,strconv.Itoa(n))
				lock1.Unlock()
			}
		}(i)
	}
	wg.Wait()
	end = time.Now()
	fmt.Printf("s11 length =%v,%v\n", len(s11),end.Sub(start))
}
