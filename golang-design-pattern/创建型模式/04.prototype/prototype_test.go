package prototype

import (
	"fmt"
	"testing"
)

var manager *PrototypeManager

type Type1 struct {
  name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}

type Type2 struct {

}

func (t *Type2) Clone() Cloneable  {
	tc := *t
	return &tc
}

func init()  {
	manager = NewPrototypeManager()
	var t1 Cloneable
	t1 = &Type1{
		name: "Type1",
	}
	manager.Set("t1",t1)
}

func TestClone(t *testing.T) {
	t1 := manager.Get("t1")
	tc := t1.Clone()
	fmt.Println(t1,tc)
	fmt.Println(t1==t1,t1==tc,tc==tc)
	if t1 == tc {
		t.Fatal("error! get clone not working")
	}
}

func TestCloneFromManager(t *testing.T) {
	c := manager.Get("t1")
	t1 := c.(*Type1)
	if t1.name != "Type1" {
		t.Fatal("error")
	}
}