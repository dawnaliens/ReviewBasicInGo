package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	res := make([]bool, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				res = append(res, nums[i] == nums[j])
			}
		}
	}
	return len(res)
}

func main() {
	nums1 := []int{1, 2, 3, 1, 1, 3}
	nums2 := []int{1, 1, 1, 1}
	res1 := numIdenticalPairs(nums1)
	res2 := numIdenticalPairs(nums2)
	fmt.Println(res1)
	fmt.Println(res2)

}
