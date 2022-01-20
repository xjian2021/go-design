package main

import "fmt"

func main() {
	a := []int{2, 3, 1, 9, 9}
	fmt.Println(plusOne(a))
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}
