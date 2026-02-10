package main

import "fmt"

const LoginToken string = "bhajbfhjabhfb"  // Public

func main() {
	var username string = "Mayank Mathur"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
	
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.5687326547843678563
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherStringVariable string
	fmt.Println(anotherStringVariable)
	fmt.Printf("Variable is of type: %T \n", anotherStringVariable)


	//implicit way of declaring a variable

	var website = "google.com"
	fmt.Println(website)

	// no var style

	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
