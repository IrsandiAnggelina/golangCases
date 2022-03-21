package main

import "fmt"

//parameter struct
type User struct {
	ID int
	FirstName string
	LastName string
	Email string
}

//parameter Embedded Struct
type Group struct {
	Name string
	Admin User
	Users []User
}

func main() {
	//input data struct sesuai parameter
	user1 := User{1, "Nur", "Aini", "aini@gmail.com"}
	user2 := User{2, "Nur", "Hasyifa", "syifa@gmail.com"}

	//input data Embedded Struct
	users := []User{user1, user2}
	group1 := Group{"Divisi IT Group", user1, users}

	// userFunc(user1)
	groupFunc(group1)
}

//fungsi struct
func userFunc(user User)string{
	return fmt.Sprintf("Nama: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
}

//fungsi Embedded Struct
//Nama Group, Jumlah User, User Names
func groupFunc(group Group){
	fmt.Printf("Nama Group: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Jumlah User: %d", len(group.Users))
	fmt.Println("")
	fmt.Printf("User Names: ")
	for _,user := range group.Users{
		fmt.Println(user.FirstName, user.LastName)
}
}