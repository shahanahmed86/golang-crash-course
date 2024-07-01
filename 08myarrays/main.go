package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruits = [4]string{}
	fruits[0] = "Apple"
	fruits[1] = "Mango"
	fruits[3] = "Banana"

	fmt.Println("Fruits list is: ", fruits)
	fmt.Println("The length of the fruits list is: ", len(fruits))

	var vegetables = [3]string{"Potato", "Bean", "Mushroom"}
	fmt.Println("Vegetables list is: ", vegetables)
	fmt.Println("The length of the vegetables list is: ", len(vegetables))
}
