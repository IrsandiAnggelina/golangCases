package main

import "fmt"

//nama saya adalah: anggel
func main() {
	result := stringNama("Anggel")
	fmt.Println(result)

}

func stringNama(nama string)string{
	newNama := "Nama Saya :" + nama
	return newNama

	//isinya : input/parameter , proses, output(return)
}
