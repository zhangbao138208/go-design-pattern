package composite

import "fmt"

type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Print(string)
}

func NewComponent(kind int,name string) Component {
	var c Component
	switch kind {
	case LeftNode:
		c = NewLeft()
	case CompositeNode:
		c = NewComposite()
	}

	c.SetName(name)
	return c
}

type component struct {
	parent Component
	name string
}
const (
	LeftNode = iota
	CompositeNode
)
func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(cp Component)  {
	c.parent = cp
}

func (c *component) Name() string  {
	return c.name
}

func (c *component) SetName(n string)  {
	c.name = n
}
func (*component) AddChild( Component)  {

}
func (*component) Print(string)  {

}

type Left struct {
	component
}

func NewLeft() Component {
	return &Left{}
}

func (l *Left) Print(pre string)  {
	fmt.Printf("%s-%s\n",pre,l.name)
}

type Composite struct {
	component
	Children []Component
}

func NewComposite() Component {
	return &Composite{}
}

func (c *Composite) AddChild(com Component)  {
	com.SetParent(c)
	c.Children = append(c.Children,com)
}
func (c *Composite)Print(pre string)  {
	fmt.Printf("%s+%s\n",pre,c.name)
	pre+=" "
	for _, child := range c.Children {
		child.Print(pre)
	}
}