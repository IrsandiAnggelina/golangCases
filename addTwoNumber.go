package main

import "fmt"

// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.

func main() {
	// var l1 []int
	var size int

	fmt.Println("Masukkan jumlah array: ")
	fmt.Scan(&size)

	l1 := make([]int, size)
	l2 := make([]int, size)
	result := make([]int, size)

	fmt.Print("Input array 1:")
	for i:= range l1{
		fmt.Scan(&l1[i])}

	fmt.Print("Input array 2:")
	for j := range l2{
		fmt.Println(&l2[j])
	}

	fmt.Print("Sum: ")
	for x := range result {
		result[x] = l1[x] + l2[x]
		fmt.Print(result[x], " ")
	}
	fmt.Println()
}

