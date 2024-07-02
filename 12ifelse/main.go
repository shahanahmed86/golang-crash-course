package main

import "fmt"

func main() {
	fmt.Println("Welcome to if else in golang")

	loginCount := 10

	var message string

	if loginCount < 10 {
		message = "Regular user"
	} else if loginCount > 10 {
		message = "Watch out"
	} else {
		message = "Exactly 10 login count"
	}

	fmt.Println(message)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}

}
