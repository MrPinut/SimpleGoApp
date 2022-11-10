package main

import (
	"fmt"
	"sort"
)

func main() {

	var words []string
	var inputs string
	var bool bool
	bool = true

	fmt.Print("Enter a word: ")
	fmt.Scanln(&inputs)

	for bool {

		if inputs != "Stop" {
			words = append(words, inputs)
			fmt.Print("Enter a word or Stop if you want to stop: ")
			fmt.Scanln(&inputs)
		} else {
			bool = false
		}
	}
	fmt.Print(words)

	sort.Strings(words)
	fmt.Print(words)
}
