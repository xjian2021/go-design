package main

import "fmt"

func main() {
	x := 112236
	fmt.Println(mySqrt(x), mySqrt2(x))
}

//mySqrt 暴力法
func mySqrt(x int) int {
	var count int
	defer func() {
		fmt.Println("count:", count)
	}()
	if x == 1 {
		return 1
	}
	var j int
	for i := 1; i < x; i++ {
		count++
		if i*i <= x {
			j = i
		} else {
			return j
		}
	}
	return j
}

//mySqrt2 官方二分查找
func mySqrt2(x int) int {
	var count int
	defer func() {
		fmt.Println("count2:", count)
	}()
	l, r := 0, x
	ans := -1
	for l <= r {
		count++
		mid := (r-l)/2 + l
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
