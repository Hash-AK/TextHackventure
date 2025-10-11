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
	Features    map[string]string
	Items       map[string]*Item
	NPCs        map[string]*NPC
}

type Player struct {
	CurrentRoom *Room
	Inventory   map[string]*Item
}
type Exit struct {
	Destination *Room
	Description string
}

type Item struct {
	Name        string
	Description string
}
type NPC struct {
	Name        string
	Description string
	Dialogue    string
	TalkedTo    bool
}

func main() {
	hermit := &NPC{
		Name:        "hermit",
		Description: "A very old man with a long, white beard sits by the empty fireplace, rocking slowly. He seems lost in toughts",
		Dialogue:    "He looks up at you with cloudy eyes. 'The wyrm of the gate is a creature of shadow and flame...' He paused, looking lost. Then with hysteria he added 'They told me the trasure was worth it! The Alchemists Guild, the geologists... They all lied! It's all a trap! The gass... the flame... the GREEN, THE GREEN LIGHT!'. He then fell back on his seat, mumbling understandable words.",
	}
	fmt.Println(hermit.Name)
	beginningRoom := Room{
		Name:        "The Dark Forest",
		Description: "You are in a deep, dark, mysterious forest. Tall, majestuous trees surround you from all sides, except for a tiny path heading north, deeper and deeper into the woods. The air is cool and damp.",
		Exits:       make(map[string]*Exit),
	}
	hut := Room{
		Name:        "The Abandonned Hut",
		Description: "You stumble upon a vast clearing. A path continue toward the mountain on the east, and a small stream goes toward the west. In the center of the clearing, an old, visibly abandonned hut stands. It's thathched roof is smashed on quite a few place. As you enter it, you see and old, dusty table. A small note is placed on it. In the corner, an old hermit sits rocking in a chair, starring blankly at the empty fireplace.Leaning against the wall, you see a sturdy-looking pickaxe",
		Exits:       make(map[string]*Exit),
		Features:    make(map[string]string),
		Items:       make(map[string]*Item),
		NPCs:        make(map[string]*NPC),
	}
	caveEntrance := Room{
		Name:        "The Cave Entrance",
		Description: "After entering the cave entrance, you feel the heavy weight of tons and tons of rock surroudning you. Suddenly, you reach the end of the tunnel. Everything beyond is obstructed by big boulders. To continue, you will need to clean up the path...",
		Exits:       make(map[string]*Exit),
	}

	crystalCave := Room{
		Name:        "The Crystal Cave",
		Description: "As you look around, you realize that weird little crystals emerge of the rock. They cover the walls and ceiling of the cave, makink long, brilliant veins. A long-forgotten sign lies on one wall. You realize that the little lamps you saw earlier aren't usual lamps : they glow from a weird, magical green flame...",
		Exits:       make(map[string]*Exit),
		Items:       make(map[string]*Item),
		Features:    make(map[string]string),
	}
	riverbank := Room{
		Name:        "The Riverbank",
		Description: "You reach the bank of the river. Here, it become pretty wide. In the clear, cold water, you see lots of rocks. A path with a sign beside it goes toward the north, near what look like old ruins. Turning back your attention to the river and it's mesmerizing sound, you realize that the stones on the bank look really much like a flint...",
		Exits:       make(map[string]*Exit),
		Items:       make(map[string]*Item),
		Features:    make(map[string]string),
	}
	guardianGate := Room{
		Name:        "The Guardian Gate",
		Description: "You stand on a wide, rocky plateau. In front of you stand a gigantic gate, made of an unknown, dark stone. A fearsome dragon, it's scales the color of obsidian, lies coiled before it. In the center of the plateau is a small, unlit fire pit...",
		Exits:       make(map[string]*Exit),
		Features:    make(map[string]string),
	}
	treasureRoom := Room{
		Name:        "The Treasure Room",
		Description: "You stuble inside a vast, stone-paved room. In it's center lies an enormous chest, debording with gold, jewels and other treasures. You made it! You found the legendary treasure of the Guardian Gate!",
	}
	beginningRoom.Exits["north"] = &Exit{
		Destination: &hut,
		Description: "As you walk carefully along the small path, the trees slowly close behind you, hiidng any trace of passage. You walk there a long, long time, so long you even forget the perception of time. Finally, you see light in front of you",
	}
	hut.Exits["south"] = &Exit{
		Destination: &beginningRoom,
		Description: "You head back into the oppressive darkness of the forest",
	}
	hut.Exits["east"] = &Exit{
		Destination: &caveEntrance,
		Description: "You head on the small path going east. The forest become lighter, and you see big boulders on the side of the path. Soon enough you reach the foot of the mountain. A mysterious cave entrance is there, just in front of you.",
	}
	hut.Exits["west"] = &Exit{
		Destination: &riverbank,
		Description: "You follow the gargling stream. It's sound is like an invitation to continue, mesmerizing and beautiful. You walk beside it for a good while, before reaching a place where the stream becomes a river.",
	}
	hut.Features["note"] = "The note reads: \nThe beast of fire, \nGuardian of The Gate \nFears only the green sparks \nOf the earthsoul crystals..."
	hut.Items["pickaxe"] = &Item{
		Name:        "pickaxe",
		Description: "A sturdy-looking, iron-headed pickaxe. A standart tool for geology and mining, Perfect to break big rocks.",
	}
	caveEntrance.Exits["west"] = &Exit{
		Destination: &hut,
		Description: "You head back to the west, toward the clearing and the abandonned hut you saw earlier.",
	}
	crystalCave.Exits["west"] = &Exit{
		Destination: &caveEntrance,
		Description: "You head back toward the cave entrance, in the tight tunnel. You're temporarly blinded due to the change of the ligh, but then you can continue forward. Soon enough you see the light of the day",
	}
	crystalCave.Features["sign"] = "You have some trouble reading the sign, but here's what you where able to decipher : \nProperty of the Alchemists' Guild. \nHigh Concentration of Malachite\nCAUTION : Unstable. "
	player := Player{
		CurrentRoom: &beginningRoom,
		Inventory:   make(map[string]*Item),
	}
	crystalCave.Items["crystals"] = &Item{
		Name:        "crystals",
		Description: "Shards of a vibrant green Malachite, a copper carbonate hydroxide mineral. When heated, the copper ions should produce a distinct green flame.",
	}
	riverbank.Exits["east"] = &Exit{
		Destination: &hut,
		Description: "You walk back on your step, going east toward the old hut. You watch the big river becoming a small stream made of melt mountain snow.",
	}
	riverbank.Exits["north"] = &Exit{
		Destination: &guardianGate,
		Description: "Using all your courage, you follow up the path trought the old, abandonned ruins, toward the north. You are now sure of where this path led : the mystical Guardian Gate. As you look around in these long-forgotten rocks, you keep seeing sign to go back : carving in the rotten doors, painting on the walls... Soon enough, you reach a bare rocky plateau.",
	}
	riverbank.Features["sign"] = "As you approach, you feel the fear in you : the sign show's a skull, with the following text writen next to it :\nWARNING - Geological Survey Team 7: High levels of methane and sulfur gas reported beyond this point. Open flames are strictly prohibited. EXTREME DANGER.\n You also see a smaller note carved by hand under it. It reads :\nBEWARE THE GUARDIAN...\nTurn back immediatly...\nLife is worth way more than an hypotethic treasure..."
	riverbank.Items["flint"] = &Item{
		Name:        "flint",
		Description: "A piece of flint, a form of microcrystalline quartz. It has a conchoidal fracture, leaving a very sharp edge. Striking it against steel or pyrite can create high-temperature sparks. Perhaps to light up a fire?",
	}
	guardianGate.Features["fire pit"] = "A circle of blackened stones, clearly used to make a fire in an ancient time. It is currently unlit."
	guardianGate.Exits["south"] = &Exit{
		Destination: &riverbank,
		Description: "You flee back south, the dragon's glare burning at your back. You can hear back the calming noise of the river's flow.",
	}
	hut.NPCs["hermit"] = hermit
	reader := bufio.NewReader(os.Stdin)
	color.Set(color.FgGreen)
	fmt.Print("Welcome to ")
	color.Set(color.FgBlue)
	fmt.Println("TextHackventure!")
	color.Unset()
	fmt.Println("")
	fmt.Println("***************************")
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
			fmt.Println("look - look around the place you are currently in, if given no argument. If given an argument,  look at your iventory, item or feature in the room.")
			fmt.Println("go [north/south/east/west] - go in the choosed direction")
			fmt.Println("read [note/sign] - read an element of the place")
			fmt.Println("take [item] - take an item by it's name.")
			fmt.Println("inventory - list the inventory you have.")
		case "quit":
			fmt.Println("Goodbye, traveller...")
			os.Exit(0)
		case "go":
			if exit, ok := player.CurrentRoom.Exits[argument]; ok {
				fmt.Println("***************************")
				fmt.Println(exit.Description)
				fmt.Println("***************************")
				player.CurrentRoom = exit.Destination
				fmt.Println(player.CurrentRoom.Description)
				if player.CurrentRoom == &treasureRoom {
					color.Yellow("Congratulation, you have won TextHackventure!")
					os.Exit(0)
				}
			} else {
				fmt.Println("You cannot go this way!")
			}
		case "look":
			if len(fieldsCommand) == 1 {
				fmt.Println("You are in : " + player.CurrentRoom.Name)
				fmt.Println(player.CurrentRoom.Description)

				if len(player.CurrentRoom.Items) > 0 {
					fmt.Println("You also see:")
					for _, item := range player.CurrentRoom.Items {
						fmt.Printf(" - A %s\n", item.Name)
					}
				}
				return

			}
			target := fieldsCommand[1]
			if item, ok := player.Inventory[target]; ok {
				fmt.Println("In your inventory : ")
				fmt.Println(" * ", item.Name, " - ", item.Description)
				return
			}
			if item, ok := player.CurrentRoom.Items[target]; ok {
				fmt.Println("Items around you : ")
				fmt.Println(" * ", item.Name, " - ", item.Description)
				return
			}
			if feature, ok := player.CurrentRoom.Features[target]; ok {
				fmt.Println("Feature around you : ")
				fmt.Println(" * ", feature)
				return
			}
			if npc, ok := player.CurrentRoom.NPCs[target]; ok {
				fmt.Println("People around you : ")
				fmt.Println(" * ", npc.Description)
				return
			}
			fmt.Println("You don't see a '" + target + "' here.")
		case "read":
			if text, ok := player.CurrentRoom.Features[argument]; ok {
				fmt.Println(text)
			} else {
				fmt.Println("You cannot read this")
			}
		case "take":
			if item, ok := player.CurrentRoom.Items[argument]; ok {
				player.Inventory[argument] = item
				delete(player.CurrentRoom.Items, argument)
				fmt.Printf("You take the %s.\n", item.Name)
				if player.CurrentRoom == &hut {
					player.CurrentRoom.Description = "You stumble upon a vast clearing. A path continue toward the mountain on the east, and a small stream goes toward the west. In the center of the clearing, an old, visibly abandonned hut stands. It's thathched roof is smashed on quite a few place. As you enter it, you see and old, dusty table. A small note is placed on it. Leaning against the wall, there used to be a pickaxe, but you now took it."
				} else if player.CurrentRoom == &riverbank {
					player.CurrentRoom.Description = "The river rushes past. The spot on the bank where the sharp-edged flint was lying is now empty."
				}

			} else {
				fmt.Println("You cannot take this!")
			}
		case "inventory":
			fmt.Println("You are carrying: ")
			if len(player.Inventory) == 0 {
				fmt.Println("- Nothing")
			} else {
				for _, item := range player.Inventory {
					fmt.Println(" * ", item.Name, " - ", item.Description)
				}
			}
		case "use":
			switch argument {
			case "pickaxe":
				if _, hasPickaxe := player.Inventory["pickaxe"]; hasPickaxe {

					if player.CurrentRoom == &caveEntrance {
						fmt.Println("You swing the pickaxe and shatter the boulders! Slowly, you manage to create a passage big enough to let you pass. The tunnel continue toward the east, even further into the mountain...")
						caveEntrance.Exits["east"] = &Exit{
							Destination: &crystalCave,
							Description: "As you walk trought this thight tunnel, you hear the rhytmed sound of water drop falling from the ceiling. Since how long did no one came here? Where does the tunnels go? After a few minutes, you see a weird light. Suddenly, the tunnel ends in an enormous cave, lighten up by mysterious little lamps...",
						}
						caveEntrance.Description = "The big boulders that once blocked the way now lays in rubbles on the sides. The tunnel that goes to the center of the mountain, toward the east, is now fully usable. You can also go back on your step and go toward the west, in the direction of the clearing..."
					} else {
						fmt.Println("You cannot use the pickaxe here!")
					}

				} else {
					fmt.Println("You don't have a pickaxe")
				}
			case "flint":
				if _, hasFlint := player.Inventory["flint"]; hasFlint {
					if player.CurrentRoom == &guardianGate {
						fmt.Println("You strike the flint, and sparks fuses everywhere. Soon, a small cathes in the fire pit.")
						player.CurrentRoom.Description = "You stand on a wide, rocky plateau. In front of you stand a gigantic gate, made of an unknown, dark stone. A fearsome dragon, it's scales the color of obsidian, lies coiled before it. In the center of the plateau is a small, the then-unlit fire pit now is the home of a nice fire. It doesnt't seems to quite afraid the dragon tho..."
						player.CurrentRoom.Features["fire pit"] = "A small fire is now burning brightly in the pit"
					}
				} else {
					fmt.Println("You don't have a flint")
				}
			case "crystals":
				if _, hasCrystals := player.Inventory["crystals"]; hasCrystals {
					if player.CurrentRoom == &guardianGate {
						if strings.Contains(guardianGate.Features["fire pit"], "is now burning") {
							fmt.Println("You throw the malachite crystals into the fire. They erupt in a brilliant, magical flame! The dragon look at it with a petrified look. Terror gains him, and he flies off into the mountains. The gian gate slowly grinds open...")
							guardianGate.Description = "The much-feared dragon is gone, and the gate to the north stands wide open."
							guardianGate.Exits["north"] = &Exit{
								Destination: &treasureRoom,
								Description: "You step solonely trought the ancient gate, into the vault beyond...",
							}

						} else {
							fmt.Println("You need a fire lit first.")
						}
					}
				}
			default:
				fmt.Println("You can't use that.")

			}
		case "talk":
			if argument == "hermit" {
				if npc, ok := player.CurrentRoom.NPCs["hermit"]; ok {
					if !npc.TalkedTo {
						fmt.Println(npc.Dialogue)
						npc.TalkedTo = true
					} else {
						fmt.Println("'Leave me to my ghosts,' the old man mutters, refusing to look at you again.")
					}

				} else {
					fmt.Println("You cannot talk to that.")
				}
			}
		default:
			fmt.Println(cleanInput + " : command not found. Type 'help' to get a list of the commands.")
		}

	}

}
