package main

import (
	"fmt"

	"github.com/fatih/color"
)

type Room struct {
	Name        string
	Description string
}

func main() {
	beginningRoom := Room{
		Name:        "Dark Forest",
		Description: "You in a deep, dark, mysterious forest.",
	}
	color.Set(color.FgGreen)
	fmt.Print("Welcome to ")
	color.Set(color.FgBlue)
	fmt.Println("TextHackventure!")
	color.Unset()
	fmt.Println("")
	fmt.Println("		////////////////		")
	fmt.Println("")
	fmt.Println("Chapter 1 : The " + beginningRoom.Name)
	fmt.Println(beginningRoom.Description)

}
