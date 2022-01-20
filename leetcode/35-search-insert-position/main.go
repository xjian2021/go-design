package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6, 7, 8, 10, 12}
	target := 4
	fmt.Println(searchInsert(nums, target))
	fmt.Println(searchInsert2(nums, target), 2>>1+3)
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 || target <= nums[0] {
		return 0
	}
	if target == nums[len(nums)-1] {
		return len(nums) - 1
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	var (
		min int
		max = len(nums) - 1
	)
	for {
		median := (min + max) / 2
		fmt.Printf("need:%d nums[min:%d]:%d  nums[median:%d]:%d  nums[max:%d]:%d\n", target, min, nums[min], median, nums[median], max, nums[max])
		switch {
		case min >= median:
			return min + 1
		case max <= median:
			return max + 1
		case nums[median] == target:
			return median
		case nums[median] > target:
			max = median
		case nums[median] < target:
			min = median
		}
	}
}

//searchInsert2 官方解法
func searchInsert2(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		//(right-left)>>1 == (right-left)/2
		mid := (right-left)>>1 + left
		fmt.Printf("need:%d nums[left:%d]:%d  nums[mid:%d]:%d  nums[right:%d]:%d\n", target, left, nums[left], mid, nums[mid], right, nums[right])
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
