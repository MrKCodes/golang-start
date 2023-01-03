package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang")
	// no inheritence; No super or parent

	kartik := User{"Kartikeya", "kartikeya@test.com", true, 27}
	fmt.Println(kartik)
	fmt.Printf("Kartik details are : %+v\n", kartik)
	fmt.Printf("Name is %v and email is %v.", kartik.Name, kartik.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
