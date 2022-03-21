package main

import "fmt"

//struct : untuk membuat struktur data yang diinginkan/oop

type User struct {
	ID int
	FirstName string
	LastName string
	Email string
	IsActive bool
}

func main(){
	//nilai struct
	user1 := User{1, "Nur", "Aisyah", "aisyah@gmail.com", true}
	user2 := User{2, "Nur", "Hasyifa", "Syifa@gmail.com", true}

	//masukkan fungsi displayUser
	hasil1 := displayUser(user1)
	hasil2 := displayUser(user2)

	fmt.Println(hasil1)
	fmt.Println(hasil2)

}

func displayUser(user User)string{
	result := fmt.Sprintf("Nama: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
	return result
}