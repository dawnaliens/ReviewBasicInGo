package main

import (
	"fmt"
	"strings"
)

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

	another := "LOVE"
	anotherRes := strings.ToLower(another)
	fmt.Println(anotherRes)
}
