package prototype

import "testing"

func TestPrototype(t *testing.T) {
	pManager := NewPrototypeManager()

	cat := NewCat("sorry")
	pManager.AddPrototype("cat-sorry",cat)

	dog := &Dog{Age: 10}
	pManager.AddPrototype("dog-10",dog)

	catPrototype := pManager.GetPrototype("cat-sorry")
	cat2 := catPrototype.Clone()
	c2,ok := cat2.(*Cat)
	if !ok {
		t.Error("clone cat fail")
	}
	if c2.Name() != cat.Name() {
		t.Error("clone cat fail name is err")
	}
	t.Logf("cat2 name:%s",c2.Name())
}