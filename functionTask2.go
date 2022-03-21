package main

import (
	"errors"
	"fmt"
)

//result, err := calculate(10,2, "-")
//result, err := calculate(10,2, "*")
//result, err := calculate(10,2, "+")
//result, err := calculate(10,2, "/")
//result, err := calculate(10,2, "=")

func main(){
	result, err := calculate(10,2,"0")
	if err != nil {
		fmt.Println("Unknown Operator")
		fmt.Println(err.Error())
	}
	fmt.Println(result)

}

func calculate(num1 int, num2 int, operation string)(int,error){
	var result int
	var errorResult error

	switch operation {
	case "-" :
		result = num1 - num2
	case "+" :
		result = num1 + num2
	case "*" :
		result = num1 * num2
	case "/" :
		result = num1 / num2
	default :
		errorResult = errors.New("Operasi Gagal")
	}
	return result,errorResult
}