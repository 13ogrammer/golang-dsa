package main

import (
	"dsa/algos"
	"fmt"
)

func main() {
	println("DSA Practice in Go")

	// toSum test
	arr := []int{2, 3, 5, 6, 7, 11, 15}
	target := 8
	fmt.Printf("Two sums: [%s] with target %i", "2,3,5,6,7,11,15", 8)
	fmt.Println()
	fmt.Println(algos.TwoSum(arr, target))
}
