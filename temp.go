package main

import "fmt"

func main() {
	var size int
	fmt.Println("masukkan jumlah array:")
	fmt.Scan(&size)
	num := make([]int, size)
	fmt.Println("masukkan data array: ")
	for i := range num {
		fmt.Scan(&num[i])
	}

	var result int
	for _,value := range num{
		result = result + value
	}

	fmt.Println("Sum array adalah :", result)
	// var size int
	// fmt.Scan(size)
	// result := make([]int,size)
	// fmt.Println(result)
	// for i := range result{
	// 	fmt.Scan(result[i])
	// }

}