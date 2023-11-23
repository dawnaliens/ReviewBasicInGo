package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3}
	arr1[len(arr1)-1] = arr1[len(arr1)-1] + 1
	fmt.Println(arr1)
}
