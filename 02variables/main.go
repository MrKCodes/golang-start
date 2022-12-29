package main

import "fmt"

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
}
