package main

import "fmt"

func main() {
	nums1 := []int{4, 5, 6, 0, 0, 0}
	nums2 := []int{1, 2, 3}
	m, n := 3, 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

//merge 逆向双指针
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	if n == 0 {
		return
	}
	lenn1 := len(nums1) - 1
	m--
	n--
	for i := lenn1; i >= 0; i-- {
		if n < 0 {
			return
		}
		if m < 0 {
			copy(nums1[:n+1], nums2[:n+1])
			return
		}
		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
	}
}
