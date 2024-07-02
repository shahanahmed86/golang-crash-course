package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")

	var fruits = []string{"Apple", "Mango", "Peach"}
	fmt.Printf("Type of fruits is %T \n", fruits)

	fruits = append(fruits, "Banana", "Grapes", "Orange")
	fmt.Println(fruits)

	fruits = append(fruits[1:3], "Apple")
	fmt.Println(fruits)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted((highScores)))
	sort.Ints(highScores)
	fmt.Println(highScores)

	// how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var idx int = 2
	courses = append(courses[:idx], courses[idx+1:]...)
	fmt.Println(courses)
}
