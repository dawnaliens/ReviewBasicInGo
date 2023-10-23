package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)

	var res int
	if input%3 == 0 {
		res = input / 3
		// 11 % 3 = 2
		// 3 + 3 + 3 + 2 (4 steps)
	} else if input%3 == 2 {
		res = input/3 + 1
	} else if input == 1 {
		res = 2
	} else {
		res = input/3 + 1
	}
	fmt.Println(res)
}
