package main

import "fmt"

type message struct {
	id      int
	name    string
	address string
	phone   int
}
type Option func(m *message)

func WithId(id int) Option {
	return func(m *message) {
		m.id = id
	}
}
func WithName(name string) Option {
	return func(m *message) {
		m.name = name
	}
}

func WithAddress(address string) Option {
	return func(m *message) {
		m.address = address
	}
}

func WithPhone(phone int) Option {
	return func(m *message) {
		m.phone = phone
	}
}

func NewByOption(options ...Option) message {
	m := message{}
	for _, option := range options {
		option(&m)
	}
	return m
}

func NewByOptionWithoutId(Id int, options ...Option) message {
	m := message{id: Id}
	for _, option := range options {
		option(&m)
	}
	return m
}

func main() {
	m1 := message{}
	fmt.Println("m1=", m1)
	m2 := NewByOption(WithId(11),
		WithAddress("addr"),
		WithPhone(155555),
		WithName("timothy"))
	fmt.Println("m2=", m2)
	m3 := NewByOptionWithoutId(111,
		WithAddress("addr111"),
		WithPhone(15555511),
		WithName("timothy1"))
	fmt.Println("m3=", m3)
}
