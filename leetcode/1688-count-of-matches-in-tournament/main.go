package main

import "fmt"

//1688. 比赛中的配对次数
//给你一个整数 n ，表示比赛中的队伍数。比赛遵循一种独特的赛制：
//
//如果当前队伍数是 偶数 ，那么每支队伍都会与另一支队伍配对。总共进行 n / 2 场比赛，且产生 n / 2 支队伍进入下一轮。
//如果当前队伍数为 奇数 ，那么将会随机轮空并晋级一支队伍，其余的队伍配对。总共进行 (n - 1) / 2 场比赛，且产生 (n - 1) / 2 + 1 支队伍进入下一轮。
//返回在比赛中进行的配对次数，直到决出获胜队伍为止。
func main() {
	n := 7
	fmt.Println(numberOfMatches(n))
}

func numberOfMatches(n int) int {
	var num int
	for n > 1 {
		num += n / 2
		if n%2 == 1 {
			n++
		}
		n = n / 2
	}
	return num
}

//numberOfMatches1 官方
//数学
//思路与算法
//在每一场比赛中，输的队伍无法晋级，且不会再参加后续的比赛。由于最后只决出一个获胜队伍，因此就有 n−1 个无法晋级的队伍，也就是会有 n−1 场比赛。
func numberOfMatches1(n int) int {
	return n - 1
}
