package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func main() {
	res1 := isPowerOfTwo(1)
	res2 := isPowerOfTwo(16)
	res3 := isPowerOfTwo(3)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
}
