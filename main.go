package main

import (
	"bufio"
	"fmt"

	"os"

	"github.com/fatih/color"
)

func main() {
	fmt.Print("Enter some text : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	color.Green("You entered, %s", input)
	color.Blue("Is that right?")
	fmt.Print("> ")
	input, _ = reader.ReadString('\n')
	if input == "yes\n" {
		color.Green("Good!")
	} else if input == "no\n" {
		color.Red("Liar!")
	} else {
		color.Yellow("I did not understand")
	}
}
