package proxy

type Subject interface {
	Do() string
}

type RealSubject struct {

}

func (RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	Real RealSubject
}

func (p Proxy) Do() string {
	var res = ""
	res += "pre:"
	res += p.Real.Do()
	res += ":after"
	return res
}

type Proxy1 struct {
	R RealSubject
}

func (p Proxy1) Do() string {
	return "1111"
}