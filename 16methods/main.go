package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods in golang")
	// no inheritance in golang, no super, no parent

	shahan := User{"shahan", "shahan.khaan@gmail.com", true, 38}
	fmt.Println(shahan)

	fmt.Printf("shahan details are: %+v\n", shahan)
	fmt.Printf("Name is: %v and email is: %v\n", shahan.Name, shahan.Email)

	shahan.GetStatus()
	shahan.NewEmail()
	fmt.Printf("Name is: %v and email is: %v\n", shahan.Name, shahan.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    uint8
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
