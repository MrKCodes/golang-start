package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "This is for user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter anything random:")

	// comma ok syntax || err ok syntax

	// input, err := reader.ReadString('\n')
	input, _ := reader.ReadString('\n')
	fmt.Println("Your random value you entered: ", input)
	fmt.Printf("The type of your input is %T ", input)
}
