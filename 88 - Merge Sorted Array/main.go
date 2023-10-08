package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	var newArr []int
	for i := 0; i < m; i++ {
		newArr = append(newArr, nums1[i])
	}
	for j := 0; j < n; j++ {
		newArr = append(newArr, nums2[j])
	}
	sort.Ints(newArr)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
