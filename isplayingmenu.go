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
	IsPlaying = false
	fmt.Println("The playlist stopped playing!")
	return true
}

func handleViewQueue() bool {
	return true
}

func handleNextTrack() bool {
	return true
}
