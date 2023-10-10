package main

import "fmt"

func main() {
	nums := []int{2, 5, 1, 3, 4, 7}
	//total := []int{}
	x := []int{}
	y := []int{}
	for i := 0; i < 6; i++ {
		if i < 3 {
			x = append(x, nums[i])
		} else {
			y = append(y, nums[i])
		}
	}
	fmt.Println(x)
	fmt.Println(y)
}
