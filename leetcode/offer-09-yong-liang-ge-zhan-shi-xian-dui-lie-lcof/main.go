package main

import "fmt"

//剑指 Offer 09. 用两个栈实现队列
//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
// 如何用两个先入后出的栈实现一个先入先出的队列？
func main() {
	obj := Constructor()
	obj.AppendTail(3)
	param_2 := obj.DeleteHead()
	fmt.Println(param_2, obj.DeleteHead(), obj.DeleteHead())
}

type CQueue struct {
	inputStack  []int
	outputStack []int
}

func Constructor() CQueue {
	return CQueue{
		inputStack:  []int{},
		outputStack: []int{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.inputStack = append(this.inputStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.outputStack) == 0 {
		for i := len(this.inputStack) - 1; i >= 0; i-- {
			this.outputStack = append(this.outputStack, this.inputStack[i])
			this.inputStack = this.inputStack[:i]
		}
	}
	if len(this.outputStack) > 0 {
		out := this.outputStack[len(this.outputStack)-1]
		this.outputStack = this.outputStack[:len(this.outputStack)-1]
		return out
	}
	return -1
}
