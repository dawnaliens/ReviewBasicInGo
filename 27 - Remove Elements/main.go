package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if val != nums[i] {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	return j
}

func main() {
	nums1 := []int{3, 2, 2, 3}
	val1 := 3
	res1 := removeElement(nums1, val1)
	fmt.Println(res1)
}
