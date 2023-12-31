package main

import "fmt"

func addDigits(num int) int {
	for num >= 10 {
		sum := 0
		for ; num > 0; num /= 10 {
			sum += num % 10
		}
		num = sum
	}
	return num
}

func main() {
	num := 38
	res1 := addDigits(num)
	fmt.Println(res1)
}
