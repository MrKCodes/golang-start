package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in GoLang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["PS"] = "PowerShell"

	fmt.Println("All languages: ", languages)
	fmt.Println("PS short for: ", languages["PS"])

	// deleting the value

	delete(languages, "RB")
	fmt.Println("All languages: ", languages)

	// Iterate through the maps
	// interesting loops

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("Value is %v\n", value)
	}
}
