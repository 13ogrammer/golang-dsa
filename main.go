package main

import (
	"dsa/algos"
	"dsa/ds"
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

	// isPalindrome test
	input := "A man, a plan, a canal: Panama"
	fmt.Printf("Is palindrome?: %s", input)
	fmt.Println()
	fmt.Println(algos.IsPalindrome(input))

	// middleOfLinkedList test
	listData := [7]int{2, 3, 5, 6, 7, 11, 15}
	list := &ds.LinkedList{}
	for _, v := range listData {
		list.Append(v)
	}
	fmt.Printf("Middle of LinkedList %s", "2,3,5,6,7,11,15")
	fmt.Println()
	fmt.Println(algos.MiddleOfLinkedList(list.Head))

	// maxSumSubArray
	maxSumInput := []int{1, 2, 3, 7, 4, 1}
	fmt.Printf("Max sum input: %s", "1,2,3,7,4,1")
	fmt.Println()
	fmt.Println(algos.MaxSumSubArray(maxSumInput, 3))

}
