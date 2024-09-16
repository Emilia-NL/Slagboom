package main

import (
	"fmt"

	"github.com/Emilia-NL/Slagboom/portsniffer/port"
)

func main() {

	fmt.Println("Port Scanner in Go")

	results := port.IniteleScan("localhost")

	fmt.Println(results)
}
