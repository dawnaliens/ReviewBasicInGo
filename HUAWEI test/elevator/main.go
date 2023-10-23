package main

import (
	"fmt"
	"sort"
)

//func reverseSlice(slice []int) {
//	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
//		slice[i], slice[j] = slice[j], slice[i]
//	}
//}

func UniqueSequence(sequence []int) []int {
	var UniSequence []int
	check := make(map[int]bool)
	for i := 0; i < len(sequence); i++ {
		if !check[i] {
			UniSequence = append(UniSequence, i)
			check[i] = true
		}
	}
	sort.Slice(UniSequence, func(i, j int) bool {
		return UniSequence[i] > UniSequence[j]
	})
	return UniSequence
}

// Count the sequence the stop at the floor
func getStops(targetFloor int, sequence []int) []int {
	uniSequence := UniqueSequence(sequence)
	first := 0
	way := 1
	var res []int
	for i := 0; i < len(uniSequence); i++ {
		if first+i <= targetFloor {
			first = i + 1
		} else {
			first = i - 1
		}
		res = append(res, i)
		way = -1 * way
	}
	return res
}

func main() {
	var targetFloor int
	var sequenceLen int
	var queue []int
	fmt.Scan(&targetFloor, &sequenceLen)
	for i := 0; i < sequenceLen; i++ {
		fmt.Scan(&queue)
	}

	res := getStops(targetFloor, queue)
	for i := 0; i < len(res); i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println()

}
