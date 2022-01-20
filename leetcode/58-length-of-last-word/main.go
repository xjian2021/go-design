package main

import "fmt"

func main() {
	//s := "   fly me   to   the moon  "
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	tmp := len(s) - 1
	var lenLast int
	for i := tmp; i >= 0; i-- {
		if s[i] != ' ' {
			lenLast++
		} else {
			if lenLast > 0 {
				return lenLast
			}
			tmp--
		}

	}
	return lenLast
}
