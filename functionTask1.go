//sum function
//scores := []int{10,5,8,9,7}
//total := sum(scores)
//hasil = 39

package main

import "fmt"

func main(){
	scores := []int{10,5,8,9,7}
	sum := sum(scores)
	fmt.Println(sum)
}

func sum(scores []int)int{
	var total int
	for _,value := range(scores){
		total = total + value
	}
	return total
}

