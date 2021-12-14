package adapter

import "testing"

var expect = "adaptee method"

func TestNewAdapter(t *testing.T) {
    adaptee := NewAdaptee()
    adapter := NewAdapter(adaptee)
    rs := adapter.Request()
	if rs != expect {
		t.Fatalf("expect: %s,actual :%s\n",expect,rs)
	}
}
