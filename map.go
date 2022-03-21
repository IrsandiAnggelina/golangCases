// package main

// import "fmt"

// func main() {
// 	var myMap map[string]int
// 	myMap = map[string]int{}

// 	myMap["ujang"] = 5
// 	myMap["go"] = 6
// 	myMap["python"]= 2

// 	for _,value := range(myMap){
// 	fmt.Println(value)}

// 	}

//cara 2
package main

import "fmt"

func main() {
	myMap := map[string]string{"go": "len","javacsript":"length","java": "panjang"}

	for index,value := range(myMap) {
		fmt.Println("Key:", index, "value:", value)
	}
}