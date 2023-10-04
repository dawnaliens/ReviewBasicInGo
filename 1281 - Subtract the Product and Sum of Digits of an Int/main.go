package main

import "fmt"

func subtractProductAndSum(n int) int {
	m := 1
	s := 0
	for n > 0 {
		x := n % 10
		fmt.Println("x =", x)
		n /= 10
		fmt.Println("n =", n)
		m = m * x
		fmt.Println("m =", m)
		s += x
		fmt.Println("s =", s)
	}
	return m - s
}

func main() {
	res1 := subtractProductAndSum(234)
	fmt.Println(res1)
	fmt.Println(234 % 10)
}
