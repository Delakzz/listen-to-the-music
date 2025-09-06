package main

import "fmt"

var MENUS = map[string]map[int]string{
	"main": {
		1: "Add Track to Playlist",
		2: "View Playlist",
		3: "Play Playlist",
		4: "Exit",
	},
	"play": {
		1: "Shuffle",
		2: "Repeat",
		3: "Play Directly",
	},
	"shuffle": {
		1: "Enable",
		2: "Disable",
	},
	"repeat": {
		0: "No Repeat",
		1: "Repeat All",
		2: "Repeat Only One",
	},
	"stop": {
		1: "Stop Playing",
		2: "View Queue",
		3: "Next Track",
		4: "Go Back",
	},
}

func showMenu(menu string) {
	if _, ok := MENUS[menu]; !ok {
		fmt.Println("Invalid menu string!")
		return
	}

	for i := 1; i <= len(MENUS[menu]); i++ {
		fmt.Printf("[%d] %s\n", i, MENUS[menu][i])
	}
}

func main() {
	showMenu("main")
}
