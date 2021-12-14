package 性能测试

func MakeSliceWithoutAlloc() []int {
	var s []int
	for i := 0; i < 100000; i++ {
		s = append(s,i)
	}
	return s
}

func MakeSliceWithPreAlloc() []int  {
	var s = make([]int,0,100000)
	for i := 0; i < 100000; i++ {
		s = append(s,i)
	}
	return s
}
func MakeSliceWithPreAlloc2() []int  {
	var s = make([]int,100000)
	for i := 0; i < 100000; i++ {
		s[i] = i
	}
	return s
}
