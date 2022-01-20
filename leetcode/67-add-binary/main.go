package main

import "fmt"

func main() {
	//a, b := "1010", "1011"
	//a, b := "11", "1"
	a, b := "100", "110010"
	fmt.Println(addBinary(a, b), '1', '0')
}

func addBinary(a string, b string) string {
	alen, blen, c, i := len(a)-1, len(b)-1, "", 0
	for alen >= 0 || blen >= 0 || i > 0 {
		addc := "0"
		if alen >= 0 && a[alen] == '1' {
			i++
		}
		if blen >= 0 && b[blen] == '1' {
			i++
		}
		if i == 1 || i == 3 {
			addc = "1"
			if i == 3 {
				i--
			}
		}
		if i > 0 {
			i--
		}
		c = addc + c
		alen--
		blen--
	}
	return c
}

//func addBinary(a string, b string) string {
//	alen, blen, c, i := len(a)-1, len(b)-1, "", 0
//	for alen >= 0 || blen >= 0 || i > 0 {
//		var (
//			addc = "0"
//		)
//		if alen >= 0 && a[alen] == '1' {
//			i++
//		}
//		if blen >= 0 && b[blen] == '1' {
//			i++
//		}
//		switch i {
//		case 1:
//			addc = "1"
//			i--
//		case 2:
//			i--
//		case 3:
//			addc = "1"
//			i -= 2
//		}
//		c = addc + c
//		alen--
//		blen--
//	}
//	return c
//}
