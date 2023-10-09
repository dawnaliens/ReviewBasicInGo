package main

import (
	"fmt"
	"math"
)

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for i := 0; i < n; i++ {
		if math.Pow(float64(3), float64(i)) == float64(n) {
			return true
		}
	}
	return false
}

func main() {
	num1 := 27
	res := math.Pow(float64(3), float64(4))
	fmt.Println(res)

	for i := 0; i < num1; i++ {
		if math.Pow(float64(3), float64(i)) == float64(num1) {
			fmt.Println(float64(i))
			fmt.Println(true)
		}
	}
}
