package main

import "os"

var MainMenu = Menu{
	Title: "Main",
	Items: []MenuItem{
		{Label: "Exit", Action: handleExit},
	},
}

// func handleAddTrack()

// func handleViewPlaylist()

// func handlePlayPlaylist()

func handleExit() {
	os.Exit(0)
}
