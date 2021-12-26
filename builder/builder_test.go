package builder

import "testing"

func TestBuilder(t *testing.T) {
	xJain := &XJianEat{}
	lunch1 := NewLunch(xJain)
	lunch1.DoEat()
	xJain.GetResult()

	xNiu := &XNiuEat{}
	lunch2 := NewLunch(xNiu)
	lunch2.DoEat()
	xNiu.GetResult()
}
