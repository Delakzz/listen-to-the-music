package main

import "fmt"

var PlayPlaylistMenu = Menu{
	Title: "Play Setup",
	Items: []MenuItem{
		{Label: "Shuffle", Action: handleShuffle},
		{Label: "Repeat", Action: handleRepeat},
		{Label: "Play Directly!", Action: handlePlay},
		{Label: "Back", Action: handleBack},
	},
}

var IsPlaying = false
var CurrentQueue Queue

func handleShuffle() bool {
	showHeader("Shuffle Config")
	return true
}

func handleRepeat() bool {
	showHeader("Repeat Config")
	return true
}

func handlePlay() bool {
	IsPlaying = true
	// copy(MainQueue.arrayList, CurrentQueue)
	fmt.Println("Playlist is now being played!")
	return true
}
