package main

import "fmt"

func findTheDifference(s string, t string) byte {

	//minLen := len(s)
	//if len(t) < minLen {
	//	minLen = len(t)
	//}
	//
	//for i := 0; i < minLen; i++ {
	//	if s[i] != t[i] {
	//		return s[i]
	//	}
	//}
	//if len(s) > len(t) {
	//	return s[minLen]
	//} else if len(s) < len(t) {
	//	return t[minLen]
	//}
	//return 0
	sum := 0
	for _, v := range s {
		sum -= int(v)
	}
	for _, v := range t {
		sum += int(v)
	}
	return byte(sum)
}

func main() {
	s := "abcd"
	t := "abcde"

	str1 := ""
	str2 := "y"

	res1 := findTheDifference(s, t)
	fmt.Println(res1)

	res2 := findTheDifference(str1, str2)
	fmt.Println(res2)
}
