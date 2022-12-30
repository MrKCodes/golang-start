package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please rate this product")
	fmt.Println("Please rate this between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	// numRating, err := strconv.ParseFloat(input, 64) // this one will throw the error as it has space and \n
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
		// panic(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

}
