package main

import "fmt"

//14. 最长公共前缀
//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
func main() {
	strs := []string{"ab", "a"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}
	for i := 0; len(strs[0]) > i; i++ {
		for _, s := range strs {
			if len(s) == i || s[i] != strs[0][i] {
				return s[:i]
			}
		}
	}
	return ""
}
