package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 4)
	fmt.Println("Result is:", result)

	// proResult := proAdder([]uint8{1, 2, 3, 4, 5})
	proResult, message := proAdder(1, 2, 3, 4, 5)
	fmt.Println("Result is:", proResult, message)
}

// func proAdder(values []uint8) uint8 {
func proAdder(values ...uint8) (uint8, string) {
	sum := uint8(0)
	for _, value := range values {
		sum += value
	}
	return sum, "Salam"
}

func adder(a uint8, b uint8) uint8 {
	return a + b
}

func greeter() {
	fmt.Println("Salam from golang")
}
