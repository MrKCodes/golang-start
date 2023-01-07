package main

import "fmt"

func main() {
	fmt.Println("Loops in GoLang")

	days := []string{"Sunday", "Tuesday", "Wednessday", "Friday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	rValue := 1

	for rValue < 10 {

		if rValue == 2 {
			goto plchldr
		}

		if rValue == 5 {
			rValue++
			continue
		}
		fmt.Println("Value is: ", rValue)
		rValue++
	}

plchldr:
	fmt.Println("Jumping to placehodler")

}
