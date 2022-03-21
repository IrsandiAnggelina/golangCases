package main

import "fmt"

func main() {
	result := add(2,3)
	fmt.Println(result)
}

func add(num1, num2 int) int {
	return num1 + num2
}