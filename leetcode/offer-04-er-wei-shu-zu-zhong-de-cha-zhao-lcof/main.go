package main

import "fmt"

//剑指 Offer 04. 二维数组中的查找
//在一个 n * m 的二维数组中，每一行都按照从左到右 非递减 的顺序排序，每一列都按照从上到下 非递减 的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

func main() {
	m := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(findNumberIn2DArray(m, 5))
}

/*
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]

[
  [1,	4,	7,	11,	15],
  [2,	5,	8,	12,	19],
  [3,	6,	9,	16,	22],
  [10,	13,	14,	17,	24],
  [18,	21,	23,	26,	30]
]

*/
// 暴力法
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, s := range matrix {
		for _, h := range s {
			if h == target {
				return true
			}
			if h > target {
				break
			}
		}
	}
	return false
}

// 二分法
func findNumberIn2DArray1(matrix [][]int, target int) bool {
	return true
}
