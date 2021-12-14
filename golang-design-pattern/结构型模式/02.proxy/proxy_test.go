package proxy

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	var subject Subject
	subject = &Proxy{}
	if subject.Do() != "pre:real:after" {
		t.Fatal()
	}
}

func TestProxy1(t *testing.T) {

	s := &Proxy1{}
	s.R = RealSubject{}
	fmt.Println(s.R.Do())
}
