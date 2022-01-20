package main

import "fmt"

//1. 两数之和
//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
//你可以按任意顺序返回答案。
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	temp := make(map[int]int)
	for i, v := range nums {
		if val, ok := temp[target-v]; !ok {
			temp[v] = i
		} else {
			return []int{val, i}
		}
	}
	return nil
}
