package facade

import "fmt"

func NewAPI() API {
	return &apiImpl{
		a:&AModuleImpl{},
		b:&BModuleImpl{},
	}
}

type API interface {
	Test() string
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a apiImpl) Test() string  {
	as := a.a.TestA()
	bs := a.b.TestB()
	return fmt.Sprintf("%s\n%s",as,bs)
}

type AModuleAPI interface {
	TestA() string
}

type AModuleImpl struct {

}

func (a AModuleImpl) TestA() string {
	return "A module running"
}

type BModuleAPI interface {
	TestB() string
}

type BModuleImpl struct {

}

func (b BModuleImpl) TestB() string {
	return "B module running"
}