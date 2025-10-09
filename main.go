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
	Exits       map[string]*Exit
}

type Player struct {
	CurrentRoom *Room
	// Inventory To be added
}
type Exit struct {
	Destination *Room
	Description string
}

func main() {
	beginningRoom := Room{
		Name:        "The Dark Forest",
		Description: "You are in a deep, dark, mysterious forest. Tall, majestuous trees surround you from all sides, except for a tiny path heading north, deeper and deeper into the woods. The air is cool and damp.",
		Exits:       make(map[string]*Exit),
	}
	hut := Room{
		Name:        "The Abandonned Hut",
		Description: "You stumble upon a vast clearing. In it's center, an old, visibly abandonned hut stands. It's thathched roof is smashed on quite a few place.",
		Exits:       make(map[string]*Exit),
	}
	beginningRoom.Exits["north"] = &Exit{
		Destination: &hut,
		Description: "As you walk carefully along the small path, the trees slowly close behind you, hiidng any trace of passage. You walk there a long, long time, so long you even forget the perception of time. Finally, you see light in front of you",
	}
	hut.Exits["south"] = &Exit{
		Destination: &beginningRoom,
		Description: "You head back into the oppressive darkness of the forest",
	}
	player := Player{CurrentRoom: &beginningRoom}
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
		cleanInput := strings.TrimSpace(input)
		fmt.Printf("You entered the command: %s\n", cleanInput)
		fieldsCommand := strings.Fields(cleanInput)
		if len(fieldsCommand) == 0 {
			return
		}
		command := fieldsCommand[0]
		var argument string
		if len(fieldsCommand) > 1 {
			argument = fieldsCommand[1]
		}
		switch command {
		case "help":
			fmt.Println("command list :")
			fmt.Println("help - display this help menu")
			fmt.Println("describe - describe the place you are currently in")
			fmt.Println("go [north/south/east/west] - go in the choosed direction")
		case "quit":
			fmt.Println("Goodbye, traveller...")
			os.Exit(0)
		case "go":
			if exit, ok := player.CurrentRoom.Exits[argument]; ok {
				fmt.Println(exit.Description)
				fmt.Println("////////////")
				player.CurrentRoom = exit.Destination
				fmt.Println(player.CurrentRoom.Description)
			} else {
				fmt.Println("You cannot go this way!")
			}
		case "describe":
			fmt.Println("You are in : " + player.CurrentRoom.Name)
			fmt.Println(player.CurrentRoom.Description)

		default:
			fmt.Println(cleanInput + " : command not found. Type 'help' to get a list of the commands.")
		}

	}

}
