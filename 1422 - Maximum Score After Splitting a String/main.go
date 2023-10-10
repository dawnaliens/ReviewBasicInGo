package main

import (
	"fmt"
	"strings"
)

var vowels = "aeiouAEIOU"

func vowelStrings(words []string, left int, right int) int {
	var trueRes []bool
	for i := left; i <= right; i++ {
		firstChar := strings.ToLower(words[i][:1])
		lastChar := strings.ToLower(words[i][len(words[i])-1:])
		if strings.Contains(vowels, firstChar) && strings.Contains(vowels, lastChar) {
			trueRes = append(trueRes, true)
		}
	}
	return len(trueRes)
}

func main() {
	words := []string{"hey", "aeo", "mu", "ooo", "artro"}
	res1 := vowelStrings(words, 1, 4)
	fmt.Println(res1)

	second_words := []string{"are", "amy", "u"}
	res2 := vowelStrings(second_words, 0, 2)
	fmt.Println(res2)
}
