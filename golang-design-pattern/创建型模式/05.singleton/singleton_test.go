package singleton

import (
	"fmt"
	"sync"
	"testing"
)

var parallelCount = 100

func TestSingleton(t *testing.T) {
	t1 := GetInstance()
	t2 := GetInstance()
	fmt.Println(t1,t2)
	if t1 != t2 {
		t.Fatal("error")
	}
}

func TestParallelSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(parallelCount)
	instances := make([]*Singleton,0,parallelCount)
	for i := 0; i < parallelCount; i++ {
		go func() {
			defer  wg.Done()
			instances = append(instances,GetInstance())
		}()
	}
	wg.Wait()
	fmt.Println(len(instances))

	for i := 1; i < parallelCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
