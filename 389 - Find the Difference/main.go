package main

import (
	"fmt"
)

//func findTheDifference(s string, t string) byte {
//
//}

func main() {
	s := "abcd"
	t := "abcde"
	sLen := len(s)
	tLen := len(t)

	// output := make([]byte, 0)
	s1 := []byte(s)
	fmt.Println(s1)

	t1 := []byte(t)
	fmt.Println(t1)

	difference := make([]byte, 0)
	for i := 0; i < sLen || i < tLen; i++ {
		
	}
}
