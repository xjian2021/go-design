package main

import (
	"fmt"
	"math"
)

//剑指 Offer 30. 包含min函数的栈
//定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。
/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-1)
	fmt.Println(obj.Min()) // -2
	fmt.Println(obj.Top()) // -1
	obj.Pop()
	fmt.Println(obj.Min()) // -2
}

type MinStack struct {
	valArr []int
	minArr []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		valArr: []int{},
		minArr: []int{math.MaxInt},
	}
}

func (this *MinStack) Push(x int) {
	min := x
	this.valArr = append(this.valArr, x)
	if x > this.minArr[len(this.minArr)-1] {
		min = this.minArr[len(this.minArr)-1]
	}
	this.minArr = append(this.minArr, min)
}

func (this *MinStack) Pop() {
	this.valArr = this.valArr[:len(this.valArr)-1]
	this.minArr = this.minArr[:len(this.minArr)-1]
}

//Top 栈顶，而非最大值
func (this *MinStack) Top() int {
	return this.valArr[len(this.valArr)-1]
}

func (this *MinStack) Min() int {
	return this.minArr[len(this.minArr)-1]
}
