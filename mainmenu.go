package main

import (
	"fmt"
	"os"
)

var MainMenu = Menu{
	Title: "Main",
	Items: []MenuItem{
		{Label: "Add Track to Playlist", Action: handleAddTrack},
		{Label: "View Playlist", Action: handleViewPlaylist},
		{Label: "Play Playlist", Action: handlePlayPlaylist},
		{Label: "Exit", Action: handleExit},
	},
}

func handleAddTrack() bool {
	title := ReadString("Title of the track")
	artist := ReadString("Track Artist")
	track := Track{Title: title, Artist: artist}
	MainQueue.Enqueue(track)
	fmt.Println("Track has been added to the queue!")
	return true
}

func handleViewPlaylist() bool {
	fmt.Println()
	showHeader("My Playlist")
	fmt.Println(MainQueue.String())
	showHeader("End")
	scanner.Scan()
	return true
}

func handlePlayPlaylist() bool {
	PlayPlaylistMenu.Run()
	return true
}

func handleExit() bool {
	os.Exit(0)
	return true
}
