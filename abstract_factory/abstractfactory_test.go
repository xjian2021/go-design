package abstract_factory

import "testing"

func getParentsAndSon(factory HomeFactory) {
	factory.CreateParents().GoToWork()
	factory.CreateSon().Study()
}

func TestHomeFactory(t *testing.T) {
	j := &JJHome{}
	getParentsAndSon(j)

	x := &XjianHome{}
	getParentsAndSon(x)
}
