package main

//imports the function FMT

import (
	"fmt"
	"time"
)

//Making a hardcoded string of license plates for testing purposes

var kentekens = []string{

	"ken-tek-en", "tes-ten", "344-321",
}

// functie om te controleren of kenteken correct is
func scanteken(kents string) bool {

	for _, k := range kentekens {

		if k == kents {

			return true
		}
	}
	return false
}

func main() {

	fmt.Print("Welcome to Fonteyn, Het beste vakantiepark volgens de consumentenbond\n")

	var kenteken string
	fmt.Print("Voer uw kenteken in: ")
	fmt.Scanln(&kenteken)

	if scanteken(kenteken) {

		tijdophalen := time.Now().Hour()

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
	} else {
		fmt.Print("U heeft helaas geen toegang tot het parkeerterrein")
	}
}
