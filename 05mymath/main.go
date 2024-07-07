package main

import (
	// "math/rand"
	// "time"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

func main() {
	// fmt.Println("Welcome to math in golang")

	// var numOne int = 2
	// var numTwo float64 = 4

	// fmt.Println("The sum is", numOne+numTwo)

	// random number - using "math/rand"
	// rand.New(rand.NewSource(time.Now().UnixNano()))
	// fmt.Println(rand.Intn(6) + 1)

	// random number - using "crypto/rand"
	myRand, err := rand.Int(rand.Reader, big.NewInt(6))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The random number is", myRand)
}
