package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer in golang")
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")

	myDefer()
}

func myDefer() {
	for i := uint8(0); i < 5; i++ {
		defer fmt.Print(i)
	}
}

// input: world one two 43210
// output: hello 43210 two one world
