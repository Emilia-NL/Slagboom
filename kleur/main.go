package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var color string
	var text string

	fmt.Println("Welke kleur?")
	fmt.Scanln(&color)

	switch color {
	case "rood":
		text = "Rood met passie."

	case "blauw":
		text = "Blauw zoals de lucht."

	case "geel":
		text = "Geel als de stralen van de zon."

	case "groen":
		text = "Groen van de natuur."

	default:
		fmt.Println("Die ken ik niet, weg ermee!")

	}

	file, err := os.Create("mooie kleur.txt")
	if err != nil {

		fmt.Println("er is een fout gemaakt tijdens creatie")
		return

	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(text)
	if err != nil {
		fmt.Println("Er is een fout opgetreden bij het schrijven naar het bestand:", err)
		return

	}
	writer.Flush()
	fmt.Println("De tekst is succesvol opgeslagen in 'lievelingskleur.txt'.")

}
