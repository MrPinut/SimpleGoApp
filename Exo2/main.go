package main

import (
	"fmt"
)

func main() {

	var nbTx int
	var temps string
	var nbTemps int

	fmt.Print("Enter the number of Tx: ")
	fmt.Scanln(&nbTx)
	fmt.Print("Enter secound, minute, hour : ")
	fmt.Scanln(&temps)
	fmt.Print("Enter number of ", temps, " : ")
	fmt.Scanln(&nbTemps)

	fmt.Print("The volume is :", nbTx/nbTemps, "per ", temps)

}
