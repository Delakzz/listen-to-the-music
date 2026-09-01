package main

type Track struct {
	Title  string
	Artist string
}

// GetTitle returns the track's title.
func (t *Track) GetTitle() string {
	return t.Title
}

// GetArtist returns the track's artist.
func (t *Track) GetArtist() string {
	return t.Artist
}

// SetTitle sets the track's title.
func (t *Track) SetTitle(title string) {
	t.Title = title
}

// SetArtist sets the track's artist.
func (t *Track) SetArtist(artist string) {
	t.Artist = artist
}

// IsEqual reports whether two tracks have the same title and artist.
func (t *Track) IsEqual(other Track) bool {
	return t.Title == other.Title && t.Artist == other.Artist
}

// String returns a human-readable "title by artist" representation.
func (t *Track) String() string {
	return t.Title + " by " + t.Artist
}
