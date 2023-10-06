package main

import (
	"fmt"
	"math/big"
)

func binaryToDecimal(binaryStr string) (*big.Int, error) {
	decimal := new(big.Int)
	decimal, success := decimal.SetString(binaryStr, 2)
	if !success {
		return nil, fmt.Errorf("Invalid binary string: %s", binaryStr)
	}
	return decimal, nil
}

func main() {
	binaryStr := "101010" // 二进制字符串
	decimal, err := binaryToDecimal(binaryStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Binary: %s\nDecimal: %s\n", binaryStr, decimal.String())
}
