package easy_factory

import "testing"

func TestNewSay(t *testing.T) {
	sayHello := NewSay(1)
	sayHello.Say("man")

	sayBye := NewSay(2)
	sayBye.Say("man")
}
