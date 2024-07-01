package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers in golang")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("actual pointer is", *ptr)

	*ptr *= 2
	fmt.Println("New value is", myNumber)
}
