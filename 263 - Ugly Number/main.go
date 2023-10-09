package main

import "fmt"

var factors = []int{2, 3, 5}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for _, f := range factors {
		for n%f == 0 {
			n /= f
		}
	}
	return n == 1
}

func main() {
	res1 := isUgly(6)
	res2 := isUgly(1)
	res3 := isUgly(14)
	res4 := isUgly(8)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)

}
