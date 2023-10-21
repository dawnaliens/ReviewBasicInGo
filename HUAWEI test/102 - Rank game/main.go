package main

import "fmt"

func main() {
	array := []int{23, 45, 100, 200, 99, -1, -2}
	for k, v := range array {
		fmt.Printf("The %v has value %v\n", k, v)
	}
}
