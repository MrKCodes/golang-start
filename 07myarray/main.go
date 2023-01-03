package main

import "fmt"

func main() {
	fmt.Println("Arrays in GoLang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Strawberry"

	fmt.Println("Fruit list is: ", fruitList)

	var vegList = [3]string{"potato", "beans", "peas"}
	fmt.Println("Vegy list is: ", vegList)
}
