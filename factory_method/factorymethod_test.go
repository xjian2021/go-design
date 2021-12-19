package factory_method

import "testing"

func TestFactory(t *testing.T) {
	rngF := &RngFactory{}
	rng := rngF.Create("大", 3)
	rng.Body()
	rng.Hand()
	rng.Face()

	edgF := &EdgFactory{}
	edg := edgF.Create("小", 1)
	edg.Body()
	edg.Hand()
	edg.Face()
}
