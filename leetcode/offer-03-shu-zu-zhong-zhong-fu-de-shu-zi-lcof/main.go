package main

import "fmt"

//剑指 Offer 03. 数组中重复的数字
//找出数组中重复的数字。
//
//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

func main() {
	fmt.Println(findRepeatNumber1([]int{2, 3, 1, 0, 2, 5, 3}))
}

//findRepeatNumber 哈希表，省时间
func findRepeatNumber(nums []int) int {
	tmp := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := tmp[num]; ok {
			return num
		}
		tmp[num] = struct{}{}
	}
	return -1
}

/*
v: 2, 3, 1, 0, 2, 5, 3
i: 0, 1, 2, 3, 4, 5, 6

v: 1, 3, 2, 0, 2, 5, 3
i: 0, 1, 2, 3, 4, 5, 6

v: 3, 1, 2, 0, 2, 5, 3
i: 0, 1, 2, 3, 4, 5, 6

v: 0, 1, 2, 3, 2, 5, 3
i: 0, 1, 2, 3, 4, 5, 6
               ↓
v: 0, 1, 2, 3, 2, 5, 3
i: 0, 1, 2, 3, 4, 5, 6

*/

//findRepeatNumber1 原地交换法，省空间
func findRepeatNumber1(nums []int) int {
	var i int
	for i < len(nums) {
		if nums[i] == i {
			i++
			continue
		}
		// 如果遇到值与索引不对应，但以值为索引的值却相等，则说明找到相同值
		if nums[i] == nums[nums[i]] {
			return nums[i]
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
	return -1
}
