package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	// days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	// fmt.Println("Loop type 1")
	// for i := 0; i < len(days); i++ {
	// 	day := days[i]
	// 	fmt.Printf("The day is %v, at index of %v\n", day, i)
	// }

	// fmt.Println("Loop type 2")
	// for i := range days {
	// 	fmt.Printf("The day is %v, at index of %v\n", days[i], i)
	// }

	// fmt.Println("Loop type 3")
	// for i, day := range days {
	// 	fmt.Printf("The day is %v, at index of %v\n", day, i)
	// }

	fmt.Println("Loop type 4")
	roughValue := 1

	for roughValue < 10 {
		if roughValue == 5 {
			roughValue++
			continue
		}

		if roughValue == 2 {
			goto lco
		}
		fmt.Println("Value is:", roughValue)
		roughValue++
	}

lco:
	fmt.Println("Jumping at LearnCodeOnline.in")
}
