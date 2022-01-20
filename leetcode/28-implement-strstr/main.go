package main

import "fmt"

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	if needleLen > haystackLen {
		return -1
	}
	if needle == "" {
		return 0
	}

	for i := 0; i < haystackLen; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
		if haystackLen-i == needleLen {
			return -1
		}
	}
	return -1
}
