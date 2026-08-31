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
