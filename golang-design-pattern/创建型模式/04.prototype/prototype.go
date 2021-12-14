package prototype

type Cloneable interface {
	Clone() Cloneable
}

type PrototypeManager struct {
	prototypes map[string] Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *PrototypeManager) Set(k string,v Cloneable)  {
	p.prototypes[k] = v
}

func (p *PrototypeManager) Get(k string) Cloneable {
	return p.prototypes[k]
}