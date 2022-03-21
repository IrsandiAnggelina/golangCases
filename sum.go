package main

import "fmt"

func main() {
	//hitung rata-rata
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var total int

	for _, value := range scores {
		total = total + value
	}
	
	length := len(scores)
	avarage := float64(total)/float64(length)

	fmt.Println(avarage)
}