package main

import "fmt"

func main() {
	n := 5
	start := 0
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = start + 2*i
	}
	fmt.Println(nums)
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	fmt.Println(res)
}
