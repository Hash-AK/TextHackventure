package main

import (
	"bufio"
	"fmt"

	"os"

	"github.com/fatih/color"

	"strings"
)

func main() {
	fmt.Print("Enter some text : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	color.Green("You entered, %s", input)
	color.Blue("Is that right?")
	fmt.Print("> ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "yes" {
		color.Green("Good!")
	} else if input == "no" {
		color.Red("Liar!")
	} else {
		color.Yellow("I did not understand")
	}
}
