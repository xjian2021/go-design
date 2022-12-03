package main

import "fmt"

//剑指 Offer 53 - I. 在排序数组中查找数字 I
//统计一个数字在排序数组中出现的次数。

func main() {
	nums := []int{5, 7, 7, 8, 9, 10}
	//fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
	//fmt.Println(missingNumber([]int{0}))
	fmt.Println(search1(nums, 7))
}

// 暴力法
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

// 二分法
func searchBoundary(nums []int, l, r, target int) int {
	for l <= r {
		middle := l + (r-l)>>1
		if nums[middle] < target {
			l = middle + 1
		} else {
			r = middle - 1
		}
	}
	return l
}

func search1(nums []int, target int) int {
	li := searchBoundary(nums, 0, len(nums)-1, target)
	ri := searchBoundary(nums, li, len(nums)-1, target+1)
	return ri - li
}

//剑指 Offer 53 - II. 0～n-1中缺失的数字
//一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
func missingNumber(nums []int) int {
	for i, num := range nums {
		if i != num {
			return i
		}
	}
	return len(nums)
}
