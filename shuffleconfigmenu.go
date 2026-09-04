package main

import "fmt"

var ShuffleConfigMenu = Menu{
	Title: "Shuffle Config",
	Items: []MenuItem{
		{Label: "Enable", Action: handleEnableShuffle},
		{Label: "Disable", Action: handleDisableShuffle},
		{Label: "Back", Action: handleBack},
	},
}

func handleEnableShuffle() bool {
	Pl.EnableShuffle()
	fmt.Println("\nShuffle is enabled!")
	return false
}

func handleDisableShuffle() bool {
	Pl.DisableShuffle()
	fmt.Println("\nShuffle is disabled!")
	return false
}
