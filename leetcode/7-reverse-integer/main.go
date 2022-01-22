package main

import "fmt"

//7. 整数反转
//给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//
//如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
//
//假设环境不允许存储 64 位整数（有符号或无符号）。
func main() {
	x := 123
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	var rev int = 0
	for x != 0 {
		rev = rev*10 + (x % 10)
		if rev > 1<<31-1 || rev < -1<<31 {
			return 0
		}
		x = x / 10
	}
	return rev
}
