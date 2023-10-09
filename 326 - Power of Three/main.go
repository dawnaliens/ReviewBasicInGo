package main

import (
	"fmt"
	"math"
)

//func isPowerOfThree(n int) bool {
//	if n <= 0 {
//		return false
//	}
//	for i := 0; i < n; i++ {
//		if math.Pow(float64(3), float64(i)) == float64(n) {
//			return true
//		}
//	}
//	return false
//}

// Official solution
func isPowerOfThree(n int) bool {
	for n > 0 && n%3 == 0 {
		n = n / 3
	}
	return n == 1
}

func main() {
	num1 := 14348906
	res := math.Pow(float64(3), float64(4))
	fmt.Println(res)
	res1 := isPowerOfThree(num1)
	fmt.Println(res1)
	//for i := 0; i < num1; i++ {
	//	if math.Pow(float64(3), float64(i)) == float64(num1) {
	//		fmt.Println(float64(i))
	//		fmt.Println(true)
	//	}
	//}
}
