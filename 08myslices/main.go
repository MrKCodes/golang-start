package main

import "fmt"

func main() {
	fmt.Println("Slices in GoLang")

	// Same as array but without the pre defined length
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("type of fruitlist is %T\n", fruitList)

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

	// Remove a value from a slice based on the index

	var lists = []string{"1", "two", "3", "four", "5"}

	fmt.Println(lists)

	var index int = 2
	// append can be used to remove value as well
	lists = append(lists[:index], lists[index+1:]...)
	fmt.Println(lists)

}
