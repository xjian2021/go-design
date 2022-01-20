package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray2(nums))
}

// 贪心算法
// 主要思路是：确保计算最大值的过程中是正增长的
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max, res, p := nums[0], 0, 0
	for p < len(nums) {
		//当加上新元素后比之前的最大值要大 则更新最大值
		res += nums[p]
		if res > max {
			max = res
		}
		//当加上新元素后却为负数 说明新元素是个绝对值比前面子数组元素相加还大的数
		//跳过这个元素 开始新的子数组累加 寻找总和比目前为止最大值还大的子数组
		if res < 0 {
			res = 0
		}
		p++
	}
	return max
}

//maxSubArray2 动态规划DP
func maxSubArray2(nums []int) int {
	MaxSum := nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		//取 当前元素和前面连续子数组总和(简称之前累加数)相加的数(简称当前累加数) 和 当前元素 的最大值为累加数
		//当前元素是正数时 肯定累加的数最大
		//当前元素是负数时
		//	如果之前累加数为负数 则当前元素最大
		//	如果之前累加数为正数 则当前累加数最大
		res = max(res+nums[i], nums[i])
		//取 累加数 和 目前最大值 的最大值为目前最大值
		MaxSum = max(res, MaxSum)
	}
	return MaxSum
}

// 对于一个区间 [l,r][l,r]，我们可以维护四个量：
// lSum 表示 [l,r][l,r] 内以 ll 为左端点的最大子段和
// rSum 表示 [l,r][l,r] 内以 rr 为右端点的最大子段和
// mSum 表示 [l,r][l,r] 内的最大子段和
// iSum 表示 [l,r][l,r] 的区间和
type Status struct {
	lSum int
	rSum int
	mSum int
	iSum int
}

//maxSubArray3 官方解法 分治法
func maxSubArray3(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func get(nums []int, l int, r int) Status {
	if l >= r {
		return Status{
			lSum: nums[l],
			rSum: nums[l],
			mSum: nums[l],
			iSum: nums[l],
		}
	}
	middle := (r-l)>>1 + l
	ls := get(nums, l, middle)
	rs := get(nums, middle+1, r)
	return pushUp(ls, rs)
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSum := max(l.mSum, max(l.rSum+r.lSum, r.rSum))
	return Status{
		lSum: lSum,
		rSum: rSum,
		mSum: mSum,
		iSum: iSum,
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
