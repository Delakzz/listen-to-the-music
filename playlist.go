package model

type Playlist struct {
	ID     int
	Name   string
	Tracks []Track
}

func (p *Playlist) AddTrack(track Track) {
	p.Tracks = append(p.Tracks, track)
}

func (p *Playlist) RemoveTrack(track Track) {
	for i, t := range p.Tracks {
		if t == track {
			p.Tracks = append(p.Tracks[:i], p.Tracks[i+1:]...)
			break
		}
	}
}

func (p *Playlist) GetTracks() []Track {
	return p.Tracks
}

func (p *Playlist) GetTrackCount() int {
	return len(p.Tracks)
}

func (p *Playlist) ClearTracks() {
	p.Tracks = []Track{}
}

func (p *Playlist) HasTrack(track Track) bool {
	for _, t := range p.Tracks {
		if t == track {
			return true
		}
	}
	return false
}

func (p *Playlist) GetTrackByTitle(title string) *Track {
	for _, t := range p.Tracks {
		if t.Title == title {
			return &t
		}
	}
	return nil
}

func (p *Playlist) GetTrackByArtist(artist string) []Track {
	var tracks []Track
	for _, t := range p.Tracks {
		if t.Artist == artist {
			tracks = append(tracks, t)
		}
	}
	return tracks
}

func (p *Playlist) GetTrackByIndex(index int) *Track {
	if index >= 0 && index < len(p.Tracks) {
		return &p.Tracks[index]
	}
	return nil
}
