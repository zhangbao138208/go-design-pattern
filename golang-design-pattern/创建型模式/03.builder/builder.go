package builder

type Builder interface {
	Part1()
	Part2()
	Part3()
	GetResult() interface{}
}

type Director struct {
	Builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		Builder: builder,
	}
}
func (d *Director) Construct()  {
	d.Builder.Part1()
	d.Builder.Part2()
	d.Builder.Part3()
}

type StringBuilder struct {
	result string
}

func (s *StringBuilder) Part1()  {
	s.result += "1"
}
func (s *StringBuilder) Part2()  {
	s.result += "2"
}
func (s *StringBuilder) Part3()  {
	s.result += "3"
}

func (s *StringBuilder) GetResult() interface{} {
	return s.result
}

type IntBuilder struct {
	result int
}

func (i *IntBuilder) Part1()  {
	i.result += 1
}
func (i *IntBuilder) Part2()  {
	i.result += 2
}
func (i *IntBuilder) Part3()  {
	i.result += 3
}

func (i *IntBuilder) GetResult() interface{} {
	return i.result
}