package factorymethod

type operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type operatorFactory interface {
	Create() operator
}

type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

type PlusOperator struct {
	*OperatorBase
}

func (p *PlusOperator) Result() int {
	return p.a + p.b
}

type MinusOperator struct {
	*OperatorBase
}

func (m *MinusOperator) Result() int {
	 return m.a - m.b
}

type PlusOperatorFactory struct {
}

func (PlusOperatorFactory) Create() operator {
	o := &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
	return o
}

type MinusOperatorFactory struct {
}

func (m MinusOperatorFactory) Create() operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}
