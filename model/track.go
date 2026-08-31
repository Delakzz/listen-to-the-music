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
