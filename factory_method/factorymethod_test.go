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

	rngF2 := &RngFactory{}
	rng2 := CreateFactory(rngF2, "中", 2)
	rng2.Body()
	rng2.Hand()
	rng2.Face()

}
