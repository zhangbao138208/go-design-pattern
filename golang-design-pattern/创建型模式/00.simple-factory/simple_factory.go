package simplefactory

import "fmt"

type Api interface {
	Say(name string) string
}

type HelloApi struct {
	
}

func (hello HelloApi) Say(name string) string  {
	return fmt.Sprintf("hello %s",name)
}

type HiApi struct {
	
}

func ( HiApi) Say(name string) string {
	return fmt.Sprintf("Hi %s",name)
}

func NewApi(i int) Api {
	if i==1 {
		return &HiApi{}
	}else if i ==2 {
		return &HelloApi{}
	}
	return nil
}