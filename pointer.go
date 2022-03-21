package main

import "fmt"

func main() {
	number := 5
	fmt.Println("Alamat memory:", &number)
	fmt.Println("Nilai awal:", number)

	pointer(&number, 100)

	fmt.Println("Nilai Akhir:",number)
	fmt.Println("Alamat memory:", &number )
}

func pointer(old *int, new int) {
	*old = new
	fmt.Println("alamat memory function:", old)
	fmt.Println("Nilai Akhir:", *old)
}

// Alamat memory: 0xc0000aa058
// Nilai awal: 5

// alamat memory function: 0xc0000aa058
// Nilai Akhir: 100

// Nilai Akhir: 100
// Alamat memory: 0xc0000aa058

