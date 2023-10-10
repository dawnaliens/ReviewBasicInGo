package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}

func main() {
	nums1 := []int{1, 1, 2}
	res1 := removeDuplicates(nums1)
	fmt.Println(res1)

	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	res2 := removeDuplicates(nums2)
	fmt.Println(res2)
}
