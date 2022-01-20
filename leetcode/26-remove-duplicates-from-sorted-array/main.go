package main

import "fmt"

//26. 删除有序数组中的重复项
//给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
//
//不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n := removeDuplicates(nums)
	fmt.Println(n, nums[:n])
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var n int
	temp := nums[0] - 1
	for _, num := range nums {
		if num != temp {
			nums[n] = num
			n++
			temp = num
		}
	}
	nums = nums[:n]
	return n
}
