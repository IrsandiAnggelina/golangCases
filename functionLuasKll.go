package main

import "fmt"

//mempunyai dua nilai kembalian return, yaitu: luas dan kll

func main(){
	hasilLuas, hasilKeliling := calculate(2,3)
	fmt.Println(hasilLuas)
	fmt.Println(hasilKeliling)
}

func calculate(panjang int, lebar int) (int, int){
	luas := panjang * lebar
	keliling := 2 *(panjang + lebar)
	return luas,keliling
}

//model2

// func calculate(panjang int, lebar int) (luas int, keliling int){
// 	luas = panjang * lebar
// 	keliling = 2 * (panjang+lebar)
// 	return
// }