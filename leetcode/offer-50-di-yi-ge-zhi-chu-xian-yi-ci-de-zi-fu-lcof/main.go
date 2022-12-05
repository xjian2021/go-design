package main

import "fmt"

//剑指 Offer 50. 第一个只出现一次的字符
//在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

func main() {
	s := "abaccdeff"
	fmt.Println(string(firstUniqChar1(s)))
}

func firstUniqChar(s string) byte {
	if s == "" {
		return ' '
	}
	chs := [26]int{}
	for _, ch := range s {
		chs[ch-'a']++
	}
	for _, ch := range s {
		if chs[ch-'a'] == 1 {
			return byte(ch)
		}
	}
	return ' '
}

// 队列解法，延迟删除重复元素
func firstUniqChar1(s string) byte {
	n := len(s)
	chs := [26]int{}
	for i := range chs {
		chs[i] = n
	}
	ps := []byte{}
	for i, ch := range []byte(s) {
		if chs[ch-'a'] == n { // 第一次出现
			chs[ch-'a'] = i
			ps = append(ps, ch)
		} else { // 第二次出现
			chs[ch-'a'] = n + 1
			// 这里延迟删除重复元素，只要第一个元素还没出现重复的元素，就不需要往后删除。
			// 一旦出现重复的，就把前面所有的重复的元素都清理了，剩下最新不重复元素。
			for len(ps) > 0 && chs[ps[0]-'a'] == n+1 {
				ps = ps[1:]
			}
		}
	}
	if len(ps) > 0 {
		return ps[0]
	}
	return ' '
}
