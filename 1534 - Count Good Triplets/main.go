package main

import (
	"fmt"
	"math"
)

func countGoodTriplets(arr []int, a int, b int, c int) int {
	res := make([]bool, 0)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if math.Abs(float64(arr[i]-arr[j])) <= float64(a) &&
					math.Abs(float64(arr[j]-arr[k])) <= float64(b) &&
					math.Abs(float64(arr[i]-arr[k])) <= float64(c) {
					res = append(res, math.Abs(float64(arr[i]-arr[j])) <= float64(a) &&
						math.Abs(float64(arr[j]-arr[k])) <= float64(b) &&
						math.Abs(float64(arr[i]-arr[k])) <= float64(c))
				}
			}
		}
	}
	return len(res)
}

func main() {
	array1 := []int{3, 0, 1, 1, 9, 7}
	a := 7
	b := 2
	c := 3
	res1 := countGoodTriplets(array1, a, b, c)
	fmt.Println(res1)

	array2 := []int{1, 1, 2, 2, 3}
	d := 0
	e := 0
	f := 1
	res2 := countGoodTriplets(array2, d, e, f)
	fmt.Println(res2)
}
