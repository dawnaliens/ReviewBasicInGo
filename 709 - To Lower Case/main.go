package main

import "fmt"

func main() {
	s := "Hello"
	res := []byte(s)
	fmt.Println(res)
	for i := 0; i < len(res); i++ {
		if 'A' <= res[i] && res[i] <= 'Z' {
			res[i] |= 32
		}
	}
	fmt.Println(string(res))
}
