package main

import "fmt"

//剑指 Offer 53 - I. 在排序数组中查找数字 I
//统计一个数字在排序数组中出现的次数。

func main() {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
}

func search(nums []int, target int) int {
	if len(nums) == 0 || nums[0] > target || nums[len(nums)-1] < target {
		return 0
	}
	var i int
	for _, num := range nums {
		if num < target {
			continue
		}
		if num > target {
			break
		}
		i++
	}
	return i
}

//TODO 二分法
func search1(nums []int, target int) int {
	return 0
}
