package main

import "fmt"

var RepeatConfigMenu = Menu{
	Title: "Shuffle Config",
	Items: []MenuItem{
		{Label: "No Repeat", Action: handleNoRepeat},
		{Label: "Repeat All", Action: handleRepeatAll},
		{Label: "Repeat Only One", Action: handleRepeatOnlyOne},
		{Label: "Back", Action: handleBack},
	},
}

func handleNoRepeat() bool {
	Pl.NoRepeat()
	fmt.Println("\nRepeat is now set to No Repeat!")
	return false
}

func handleRepeatAll() bool {
	Pl.SetAllRepeat()
	fmt.Println("\nRepeat is now set to All Repeat!")
	return false
}

func handleRepeatOnlyOne() bool {
	Pl.SetOneRepeat()
	fmt.Println("\nRepeat is now set to Repeat Only One!")
	return false
}
