package main

import "fmt"

// 用1或2级一步爬楼梯 求登顶方法数
func main() {
	n := 6
	fmt.Println(climbStairs(n))
}

//climbStairs
//f(x) = f(x-1) + f(x-2) :x级阶梯的最后两步可以分为1级或2级两种方式
//f(0) = 1
//f(1) = 1
//f(2) = f(1) + f(0)
//...
//f(x) = f(x-1) + f(x-2)
func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	s, ss := 1, 1
	for i := 2; i < n; i++ {
		s, ss = ss, s+ss
	}
	return s + ss
}
