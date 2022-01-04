package adapter

import "testing"

func TestNewRoundHole(t *testing.T) {
	hole := NewRoundHole(5) //实例化一个半径为5的圆孔
	rPeg := NewRoundPeg(5) //实例化一个半径为5的圆钉
	if !hole.Fits(rPeg) {
		t.Errorf("圆钉半径:%f 圆孔半径:%f 钉不进去",rPeg.GetRadius(),hole.GetRadius())
	}
	t.Logf("圆钉半径:%f 圆孔半径:%f 可以使用",rPeg.GetRadius(),hole.GetRadius())
}

func TestSquarePegAdapter(t *testing.T) {
	hole := NewRoundHole(5) //实例化一个半径为5的圆孔
	sPeg := NewSquarePeg(7) //实例化一个宽为7的方钉
	//hole.Fits(sPeg) //方钉给不了圆孔用
	sAdapter := NewSquarePegAdapter(sPeg) //使用方钉适配器
	if !hole.Fits(sAdapter) {
		t.Errorf("方钉宽:%f 对角的一半:%f 圆孔半径:%f 钉不进去",sPeg.GetWidth(),sAdapter.GetRadius(),hole.GetRadius())
	}
	t.Logf("方钉宽:%f 对角的一半:%f 圆孔半径:%f 可以使用",sPeg.GetWidth(),sAdapter.GetRadius(),hole.GetRadius())
}