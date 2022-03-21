//cara 1

package main

import "fmt"

func main() {
	var gedung []int
	gedung = append(gedung, 7)
	gedung = append(gedung, 5)
	gedung = append(gedung, 8)
	gedung = append(gedung, 10)

	for _, data := range(gedung) {
		fmt.Println(data)
	}
}

//cara 2
// package main
// import "fmt"

// func main(){
// 	gedung := []int{7,5,8,10}
// 	for _,data := range(gedung){
// 		fmt.Println(data)
// 	}

// }