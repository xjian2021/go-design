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
	//m := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	for i := 0; i <= 30; i++ {
		fmt.Println("二分：", i, findNumberIn2DArray2(m, i))
		fmt.Println("暴力：", i, findNumberIn2DArray(m, i))
		//fmt.Println(findNumberIn2DArray1(m, 5))
	}

	//fmt.Println(inArray1(m[0], 1))
}

/*
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
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
//思路：f(n)=f(n/2)+n/2纵向二分   其实横着二分跟纵向二分一样
func findNumberIn2DArray1(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for i, row := range matrix[0] {
		if row > target {
			break
		}
		if row == target {
			return true
		}
		l, r := 0, len(matrix)
		if matrix[r-1][i] < target {
			continue
		}
		for l <= r {
			middle := (l + r) >> 1
			switch {
			case matrix[middle][i] == target:
				return true
			case matrix[middle][i] < target:
				l = middle + 1
			default:
				r = middle - 1
			}
		}
	}
	return false
}

/*
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
*/
// z字抖动
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	portrait, y := len(matrix), len(matrix[0])-1
	var x int
	for x < portrait && y >= 0 {
		switch {
		case matrix[x][y] > target:
			y--
		case matrix[x][y] < target:
			x++
		case matrix[x][y] == target:
			return true
		}
	}
	return false
}
