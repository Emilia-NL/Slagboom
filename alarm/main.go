package main

import (
	"fmt"
)

var alarmnum int

func readAlarm() {
	fmt.Print("how many times do i have to print alarm? ")
	_, err := fmt.Scanln(&alarmnum)
	if err != nil {
		fmt.Println("error", err)
		return
	}
}

func playAlarm() {

	for i := 0; i < alarmnum; i++ {
		fmt.Printf("alarm %d !\n", i)

	}

}

func main() {

	readAlarm()
	playAlarm()

}
