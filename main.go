package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Hello World1")
	color.Red("test")
	fmt.Println("test2")
	string1 := "test3" + " " + "lol" + "1 + 1"
	color.Green(string1)
}
