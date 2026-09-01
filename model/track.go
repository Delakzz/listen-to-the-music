package model

type Track struct {
	Title  string
	Artist string
}

func (t *Track) GetTitle() string {
	return t.Title
}

func (t *Track) GetArtist() string {
	return t.Artist
}

func (t *Track) SetTitle(title string) {
	t.Title = title
}

func (t *Track) SetArtist(artist string) {
	t.Artist = artist
}

func (t *Track) IsEqual(other Track) bool {
	return t.Title == other.Title && t.Artist == other.Artist
}

func (t *Track) String() string {
	return t.Title + " by " + t.Artist
}
