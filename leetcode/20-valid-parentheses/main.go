package main

import "fmt"

//20. 有效的括号
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
func main() {
	s := "()"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	if len([]rune(s))%2 != 0 {
		return false
	}
	lrPair := map[string]string{"(": ")", "{": "}", "[": "]", ")": "(", "}": "{", "]": "["}
	left := make([]string, 0)
	index := -1
	for _, v := range s {
		if string(v) == "(" || string(v) == "{" || string(v) == "[" {
			left = append(left, string(v))
			index++
		} else {
			if index == -1 || lrPair[string(v)] != left[index] {
				return false
			}
			left = append(left[:index], left[index+1:]...)
			index--
		}
	}
	if len(left) != 0 {
		return false
	}
	return true
}
