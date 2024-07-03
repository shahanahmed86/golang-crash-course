package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch case in golang")

	rand.Seed(time.Now().UnixNano())

	diceNum := rand.Intn(6) + 1

	switch diceNum {
	case 1:
		fmt.Println("Dice value is 1, and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
		fallthrough
	case 3:
		fmt.Println("You can move 3 spot")
		fallthrough
	case 4:
		fmt.Println("You can move 4 spot")
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot, and roll the dice again")
	default:
		fmt.Println("What was that?")
	}
}
