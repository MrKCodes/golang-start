package main

import "fmt"

func main() {
	fmt.Println("Slices in GoLang")

	// Same as array but without the pre defined length
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("tyope of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "mango", "banana")

	fmt.Println("The fruitlist is: ", fruitList)

	fruitList = append(fruitList[1:3])

	fmt.Println("FruitList ", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 220
	highScores[1] = 229
	highScores[2] = 240
	highScores[3] = 245

	fmt.Println(highScores)
}
