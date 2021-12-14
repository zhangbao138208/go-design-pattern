package adapter

type Target interface {
	Request() string
}

type adaptee interface {
	SpecificRequest() string
}

type adapteeImpl struct {

}

func (adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

func NewAdaptee() adaptee {
	return &adapteeImpl{}
}

type adapter struct {
	adaptee
}

func (a *adapter) Request() string  {
	return a.SpecificRequest()
}

func NewAdapter(a adaptee) Target {
	return &adapter{
		adaptee:a,
	}
}