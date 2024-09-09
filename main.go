package main

//imports the function FMT

import (
	"fmt"
	"time"
)

//Prints text to the console.

func main() {

	tijdophalen := time.Now().Hour()

	fmt.Print("Welcome to Fonteyn, Het beste vakantiepark volgens de consumentenbond\n")

	switch {

	case tijdophalen > 7 && tijdophalen < 12:
		fmt.Print("Goedemorgen")

	case tijdophalen >= 12 && tijdophalen <= 18:
		fmt.Print("Goedemiddag")

	case tijdophalen >= 18 && tijdophalen <= 23:
		fmt.Print("goedeavond")

	default:
		fmt.Print("We zijn gesloten")

	}
}
