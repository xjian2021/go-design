package main

import (
	"fmt"
	"strings"
)

//剑指 Offer 05. 替换空格
//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
}

func replaceSpace(s string) string {
	b := make([]rune, 0, len(s)*3)
	for _, i := range []rune(s) {
		if i == ' ' {
			b = append(b, '%', '2', '0')
		} else {
			b = append(b, i)
		}
	}
	return string(b)
}

func replaceSpace1(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}
