package 子测试

import "testing"

func sub1(t *testing.T)  {
	var a,b = 1,2
	var expected =3
	actual := add(a,b)
	if expected != actual {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
	}
}
func sub2(t *testing.T)  {
	var a,b = 1,2
	var expected =3
	actual := add(a,b)
	if expected != actual {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
	}
}
func sub3(t *testing.T)  {
	var a,b = 1,2
	var expected =3
	actual := add(a,b)
	if expected != actual {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
	}
}

func TestSub(t *testing.T) {
	t.Run("A=1",sub1)
	t.Run("A=2",sub2)
	t.Run("B=1",sub3)
}
