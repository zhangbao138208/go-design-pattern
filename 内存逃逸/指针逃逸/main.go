package main

type Student struct {
	Name string
	Age  int
}

func StudentRegister(name string, age int) *Student {
	s := new(Student) //局部变量s逃逸到堆

	s.Name = name
	s.Age = age

	return s
}
//通过编译参数-gcflag=-m可以查看编译过程中的逃逸分析：
func main() {
	StudentRegister("Jim", 18)
}