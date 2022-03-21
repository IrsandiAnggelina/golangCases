package main

import "fmt"

func main() {
	students := []map[string]string{
		{"name": "Agung", "score": "A"},
		{"name": "Link", "score": "B"},
		{"name": "Mario", "score": "C"},
	}

	for _, value := range (students) {
		fmt.Println(value)
	}
}