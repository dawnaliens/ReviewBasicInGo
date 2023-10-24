package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

func main() {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	res1 := twoSum(nums1, target1)
	fmt.Println(res1)
}
