package main

import "fmt"

func isUgly(n int) bool {
	if n%2 == 0 && n%3 == 0 && n%5 == 0 {
		return true
	} else if n%2 == 0 || n%3 == 0 || n%5 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	res1 := isUgly(6)
	res2 := isUgly(1)
	res3 := isUgly(14)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
}
