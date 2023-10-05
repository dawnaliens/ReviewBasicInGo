package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	num1Int, _ := strconv.Atoi(num1)
	num2Int, _ := strconv.Atoi(num2)
	sum := num1Int + num2Int
	sumString := strconv.Itoa(sum)
	return sumString
}

func main() {
	num1 := "11"
	num2 := "123"
	res1 := addStrings(num1, num2)
	fmt.Println(res1)
	fmt.Println(reflect.TypeOf(res1))

	num3 := "456"
	num4 := "77"
	res2 := addStrings(num3, num4)
	fmt.Println(res2)
	fmt.Println(reflect.TypeOf(res2))

	num5 := "0"
	num6 := "0"
	res3 := addStrings(num5, num6)
	fmt.Println(res3)
	fmt.Println(reflect.TypeOf(res3))

	num7 := "3876620623801494171"
	num8 := "6529364523802684779"
	res4 := addStrings(num7, num8)
	fmt.Println(res4)

	fmt.Println(3876620623801494171 + 6529364523802684779)
}
