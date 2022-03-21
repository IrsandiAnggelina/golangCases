package main

import "fmt"

//tabahkan nilai yang sama atau lebih besar 7
//[7,8,1,10] = 7+8+10=25

func main() {
	slice := []int{7, 8, 1, 10}

	//tampung nilai yang sama atau lebih besar dari 7
	var temp []int
	for _,value := range(slice) {
		if value >= 7 {
			temp = append(temp, value)
		}
	} //hasil temp = [7,8,10]

	//tambahkan nilai yang difilter
	var sum int
	for _,data := range(temp) {
		sum += data
	}
	fmt.Println(temp)
	fmt.Println(sum)
} //result 25