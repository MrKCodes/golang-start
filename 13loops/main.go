package main

import "fmt"

func main() {
	fmt.Println("Loops in GoLang")

	days := []string{"Sunday", "Tuesday", "Wednessday", "Friday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
}
