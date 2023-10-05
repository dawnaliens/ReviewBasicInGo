package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	num1 := "11"
	num2 := "123"
	another1, _ := strconv.Atoi(num1)
	another2, _ := strconv.Atoi(num2)
	fmt.Println(another1)
	fmt.Println(another2)
	fmt.Println("The type of another1 is", reflect.TypeOf(another1))
	fmt.Println("The type of another2 is", reflect.TypeOf(another2))
	num1String := strconv.Itoa(another1)
	num2String := strconv.Itoa(another2)
	fmt.Println(num1String)
	fmt.Println(num2String)
	fmt.Println("The type of num1String is", reflect.TypeOf(num1String))
	fmt.Println("The type of num2String is", reflect.TypeOf(num2String))
}
