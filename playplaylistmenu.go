package main

import (
	"fmt"
	"math/rand"
)

var PlayPlaylistMenu = Menu{
	Title: "Play Setup",
	Items: []MenuItem{
		{Label: "Shuffle", Action: handleShuffle},
		{Label: "Repeat", Action: handleRepeat},
		{Label: "Play Directly!", Action: handlePlay},
		{Label: "Back", Action: handleBack},
	},
}

func handleShuffle() bool {
	ShuffleConfigMenu.Run()
	return true
}

func handleRepeat() bool {
	RepeatConfigMenu.Run()
	return true
}

func handlePlay() bool {
	Pl.playing = true
	Pl.createPlaylistQueue(0, Pl.ArrayList.GetSize())

	// shuffle bitch
	if Pl.shuffled {
		rand.Shuffle(Pl.queue.GetSize(), func(i, j int) {
			Pl.queue.arrayList[i], Pl.queue.arrayList[j] = Pl.queue.arrayList[j], Pl.queue.arrayList[i]
		})
	}

	fmt.Println("\nPlaylist is now being played!")
	return false
}
