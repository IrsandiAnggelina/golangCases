// OUTPUT
// masukkan jumlah array:
// 4
// masukkan data array:
// 1 1 1 1
// Sum array adalah : 4

package main

import "fmt"

func main(){
	var ukuran int
	var data []int
	var result int

	fmt.Println("Masukkan Jumlah Angka:")
	fmt.Scan(&ukuran)

	data = make([]int, ukuran)
	fmt.Println("Masukkan Data:")
	for i := range data {
		fmt.Scan(&data[i])
	}

	for _,value := range data{
		result = result + value
	}

	fmt.Println("Hasil Penjumlahan:", result)
}