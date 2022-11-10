package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type data struct {
	nbHabitant int
	codePostal int
	surface    int
}

type MapVille struct {
	Ville map[string]data
}

func main() {

	ville := MapVille{Ville: make(map[string]data)}

	getVille(ville)

	printVille(ville)

}

func getVille(mapVille MapVille) {

	words := [3]string{"Number of citizen", "Zip Code", "Space"}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; i++ {

			fmt.Print(words[j])
			reader := bufio.NewReader(os.Stdin)
			// ReadString will block until the delimiter is entered
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("An error occured while reading input. Please try again", err)
				return
			}

			// remove the delimeter from the string
			input = strings.TrimSuffix(input, "\n")
			fmt.Println(input)

			marksStr := input
			marks, err := strconv.Atoi(marksStr)

			if err != nil {
				fmt.Println("Error during conversion")
				return
			}

			fmt.Println(marks)

			mapVille
		}
	}

}

func printVille(mapVille MapVille) {

}
