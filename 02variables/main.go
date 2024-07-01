package main

import "fmt"

const LoginToken string = "gibisflkamsn;io.lasjdfkljaklsjdf" // Public

func main() {
	var username string = "shahan"
	fmt.Println(username)
	fmt.Printf("Variable  is of type %T: \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T: \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type %T: \n", smallVal)

	var smallFloat float64 = 255.43434534534534534
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type %T: \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var defaultStr string
	fmt.Println(defaultStr)
	fmt.Printf("Variable is of type: %T \n", defaultStr)

	// implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// no var style
	numberOfUser := 30_000.00
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}