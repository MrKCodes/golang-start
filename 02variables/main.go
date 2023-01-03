package main

import "fmt"

const LoginToken string = "aX12TyuOoo#997fd" // Public ( beacuse the const name start with capital letter)

func main() {
	var username string = "kartikeya"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 254
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloat float64 = 2255.234235345
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	// default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	var anotherStrVariable string
	fmt.Println(anotherStrVariable)
	fmt.Printf("variable is of type: %T \n", anotherStrVariable)

	//  implicit type

	var website = "linkedin.com"
	fmt.Println(website)

	// no var style ( without using var )
	numberOfUser := 1500
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)

}
