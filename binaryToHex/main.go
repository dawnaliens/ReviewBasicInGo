package main

import (
	"fmt"
	"strconv"
)

func binaryToHex(binaryStr string) (string, error) {
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		return "", err
	}
	hexStr := fmt.Sprintf("%X", decimal)
	return hexStr, nil
}

func main() {
	binaryStr := "101010" // 二进制字符串
	hexStr, err := binaryToHex(binaryStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Binary: %s\nHexadecimal: %s\n", binaryStr, hexStr)
}
