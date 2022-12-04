package main

import (
	"fmt"
)

//剑指 Offer 11. 旋转数组的最小数字
//把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
//
//给你一个可能存在 重复 元素值的数组 numbers ，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。请返回旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一次旋转，该数组的最小值为 1。
//
//注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
//
//示例 1：
//
//输入：numbers = [3,4,5,1,2]
//输出：1

//示例 2：
//
//输入：numbers = [2,2,2,0,1]
//输出：0

func main() {
	n := []int{2, 2, 2, 0, 1}
	fmt.Println(minArray1(n))
}

func minArray(numbers []int) int {
	numbers = sortArray(numbers)
	return numbers[0]
}

//快速排序，分治法
//关键在于合并数组
func sortArray(nums []int) []int {
	l, r := 0, len(nums)
	if r <= 1 || nums[0] < nums[r-1] {
		return nums
	}
	mid := l + (r-l)>>1
	la := sortArray(nums[:mid])
	ra := sortArray(nums[mid:])
	return mergeArray(la, ra)
}

func mergeArray(a, b []int) []int {
	al, bl := len(a), len(b)
	if al == 0 {
		return b
	}
	if bl == 0 {
		return a
	}
	var i, j int
	newArrry := make([]int, 0, al+bl)
	for {
		if a[i] <= b[j] {
			newArrry = append(newArrry, a[i])
			i++
		} else {
			newArrry = append(newArrry, b[j])
			j++
		}
		if i == al {
			newArrry = append(newArrry, b[j:]...)
			break
		}
		if j == bl {
			newArrry = append(newArrry, a[i:]...)
			break
		}
	}
	return newArrry
}

//二分法
func minArray1(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	l, r := 0, len(numbers)-1
	for l < r {
		mid := l + (r-l)>>1
		// 如果中位值比右端值小，说明中位与右端是关系是顺序关系，可以忽略这段区间的数。
		// 但不能忽略中位值(不能+1)，因为中位值有可能是最小值，如果忽略了中位值，将找不到最小值
		if numbers[mid] < numbers[r] {
			r = mid
			// 如果中位值比右端值大，说明中位与右端是关系并非顺序关系，那么左端到中位这段区间的数肯定是顺序的，可忽略。
		} else if numbers[mid] > numbers[r] {
			l = mid + 1
			// 如果中位值与右端值相等，并不能说明什么，因为中位的左右两边都可能会存在最小值，但可以忽略右端值，缩小区间继续找。
		} else {
			r--
		}
	}
	return numbers[l]
}
