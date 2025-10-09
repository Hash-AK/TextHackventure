package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type Room struct {
	Name        string
	Description string
}

func main() {
	beginningRoom := Room{
		Name:        "The Dark Forest",
		Description: "You are in a deep, dark, mysterious forest. Tall, majestuous trees surround you from all sides, except for a tiny path heading north, deeper and deeper into the woods. The air is cool and damp.",
	}
	reader := bufio.NewReader(os.Stdin)
	color.Set(color.FgGreen)
	fmt.Print("Welcome to ")
	color.Set(color.FgBlue)
	fmt.Println("TextHackventure!")
	color.Unset()
	fmt.Println("")
	fmt.Println("		////////////////		")
	fmt.Println("")
	fmt.Println("Chapter 1 : " + beginningRoom.Name)
	fmt.Println(beginningRoom.Description)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		command := strings.TrimSpace(input)
		fmt.Printf("You entered the command: %s\n", command)
		switch command {
		case "quit":
			fmt.Println("Goodbye, traveller...")
			os.Exit(0)
		}
	}

}
