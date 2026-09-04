package main

import "fmt"

var IsPlayingMenu = Menu{
	Title: "Playing Config",
	Items: []MenuItem{
		{Label: "Stop Playing", Action: handleStop},
		{Label: "View Queue", Action: handleViewQueue},
		{Label: "Next Track", Action: handleNextTrack},
		{Label: "Go Back", Action: handleBack},
	},
}

func handleStop() bool {
	Pl.StopPlay()
	fmt.Println("The playlist stopped playing!")
	return false
}

func handleViewQueue() bool {
	fmt.Println(Pl.PlaylistQueue())
	ReadString("Press enter to continue")
	return true
}

func handleNextTrack() bool {
	Pl.NextTrack()
	return true
}
