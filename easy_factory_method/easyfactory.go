package easy_factory_method

import "fmt"

type SayFactory interface {
	Say(string)
}

func NewSay(t int) SayFactory {
	switch t {
	case 1:
		return &Hello{}
	case 2:
		return &Bye{}
	}
	return nil
}

type Hello struct{}

func (h *Hello) Say(s string) {
	fmt.Printf("say: Hello %s\n", s)
}

type Bye struct{}

func (b *Bye) Say(s string) {
	fmt.Printf("say: Bye %s\n", s)
}
