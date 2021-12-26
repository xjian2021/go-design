package builder

import "fmt"

//Eat 把一顿饭分解为几个动作
type Eat interface {
	DrinkSoup() // 喝汤
	EatChickenLeg() // 吃鸡腿
	EatRice() // 吃米饭
}

type Lunch struct {
	eat Eat
}

func NewLunch(e Eat) *Lunch {
	return &Lunch{eat: e}
}

func (l *Lunch) DoEat() {
	l.eat.DrinkSoup()
	l.eat.EatChickenLeg()
	l.eat.EatRice()
}

type XJianEat struct {
	result string
}

func (x *XJianEat) DrinkSoup()  {
	x.result += "喝了汤 "
}

func (x *XJianEat) EatChickenLeg() {
	x.result += "吃了鸡腿 "
}

func (x *XJianEat) EatRice() {
	x.result += "吃了米饭 "
}

func (x *XJianEat) GetResult() {
	fmt.Printf("xJian %s\n",x.result)
}

type XNiuEat struct {
	result string
}

func (x *XNiuEat) DrinkSoup()  {
	x.result += "喝了汤 "
}

func (x *XNiuEat) EatChickenLeg() {
	x.result += "吃了鸡腿 "
}

func (x *XNiuEat) EatRice() {
	x.result += "不吃米饭 "
}

func (x *XNiuEat) GetResult() {
	fmt.Printf("XNiu %s\n",x.result)
}