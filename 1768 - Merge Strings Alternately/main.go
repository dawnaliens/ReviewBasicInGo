package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	n := len(word1)
	m := len(word2)
	// make([ ]Type, length, capacity)
	ans := make([]byte, 0)
	for i := 0; i < n || i < m; i++ {
		if i < n {
			ans = append(ans, word1[i])
		}
		if i < m {
			ans = append(ans, word2[i])
		}
	}
	return string(ans)
}

func main() {
	word1 := "abc"
	word2 := "pqr"
	n := len(word1)
	m := len(word2)
	// make([ ]Type, length, capacity)
	ans := make([]byte, 0)
	for i := 0; i < n || i < m; i++ {
		if i < n {
			ans = append(ans, word1[i])
			fmt.Println("In i < n part, ans =", string(ans))
			fmt.Println()
		}
		if i < m {
			ans = append(ans, word2[i])
			fmt.Println("In i < m part, ans =", string(ans))
			fmt.Println()
		}
	}
	fmt.Println(string(ans))
}
