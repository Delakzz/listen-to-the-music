package main

// var MENUS = map[string]map[int]string{
// 	"main": {
// 		1: "Add Track to Playlist",
// 		2: "View Playlist",
// 		3: "Play Playlist",
// 		4: "Exit",
// 	},
// 	"play": {
// 		1: "Shuffle",
// 		2: "Repeat",
// 		3: "Play Directly",
// 	},
// 	"shuffle": {
// 		1: "Enable",
// 		2: "Disable",
// 	},
// 	"repeat": {
// 		0: "No Repeat",
// 		1: "Repeat All",
// 		2: "Repeat Only One",
// 	},
// 	"stop": {
// 		1: "Stop Playing",
// 		2: "View Queue",
// 		3: "Next Track",
// 		4: "Go Back",
// 	},
// }

var Pl = Playlist{
	name:     "My Playlist",
	shuffled: false,
	queue:    &Queue{},
	repeat:   0,
	playing:  false,
}

func main() {
	t1 := Track{Title: "Life Puzzle", Artist: "Arthur Nery"}
	t2 := Track{Title: "I Gotchu", Artist: "Arthur Nery"}
	t3 := Track{Title: "Tahanan", Artist: "Adie"}
	t4 := Track{Title: "Paraluman", Artist: "Adie"}
	t5 := Track{Title: "Who Knows", Artist: "Daniel Caesar"}
	t6 := Track{Title: "Pag-Ibig ay Kanibalismo II", Artist: "fitterkarma"}

	Pl.AddToPlaylist(t1)
	Pl.AddToPlaylist(t2)
	Pl.AddToPlaylist(t3)
	Pl.AddToPlaylist(t4)
	Pl.AddToPlaylist(t5)
	Pl.AddToPlaylist(t6)

	// fmt.Println(pl.String())

	MainMenu.Run()
}
