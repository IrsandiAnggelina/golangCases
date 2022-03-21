package main

import "fmt"

func main() {
	scores := []int{100, 80, 75, 92, 70, 93, 88, 67}

	//memasukkan nilai yang >=90 kedalam goodScores
	var goodScores []int
	for _, data := range scores {
		if data >= 90 {
			goodScores = append(goodScores, data)
		}
	} //result = [100,92,93]
	
	//tambahkan semua nilai filter
	var totalFilter int
	for _, value := range goodScores {
		totalFilter = totalFilter + value
	} //result = 285

	fmt.Println(goodScores)
	fmt.Println(totalFilter)
}