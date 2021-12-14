package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	var x float64 = 3.4
	t1 := reflect.TypeOf(x)
	fmt.Println("type:",t1)

	v := reflect.ValueOf(&x)
	fmt.Println("value:",v,v.Type())

	v.Elem().SetFloat(666)
	y1 := v.Interface().(*float64)
	y := v.Elem().Interface().(float64)

	fmt.Println("value:",y,*y1)
}

type MyInt int32
func Test2(t *testing.T) {
	var m MyInt = 33
	v := reflect.ValueOf(m)
	fmt.Println("kind =",v.Kind(),"type = ",v.Type())
}
