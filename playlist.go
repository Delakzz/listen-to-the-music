package main

type Playlist struct {
	ID     int
	Name   string
	Tracks []Track
}

// AddTrack appends a track to the playlist.
func (p *Playlist) AddTrack(track Track) {
	p.Tracks = append(p.Tracks, track)
}

// RemoveTrack removes the first matching track from the playlist.
func (p *Playlist) RemoveTrack(track Track) {
	for i, t := range p.Tracks {
		if t == track {
			p.Tracks = append(p.Tracks[:i], p.Tracks[i+1:]...)
			break
		}
	}
}

// GetTracks returns all tracks in the playlist.
func (p *Playlist) GetTracks() []Track {
	return p.Tracks
}

// GetTrackCount returns the number of tracks in the playlist.
func (p *Playlist) GetTrackCount() int {
	return len(p.Tracks)
}

// ClearTracks removes all tracks from the playlist.
func (p *Playlist) ClearTracks() {
	p.Tracks = []Track{}
}

// HasTrack reports whether the track exists in the playlist.
func (p *Playlist) HasTrack(track Track) bool {
	for _, t := range p.Tracks {
		if t == track {
			return true
		}
	}
	return false
}

// GetTrackByTitle returns the first track matching the given title.
func (p *Playlist) GetTrackByTitle(title string) *Track {
	for _, t := range p.Tracks {
		if t.Title == title {
			return &t
		}
	}
	return nil
}

// GetTrackByArtist returns all tracks by the given artist.
func (p *Playlist) GetTrackByArtist(artist string) []Track {
	var tracks []Track
	for _, t := range p.Tracks {
		if t.Artist == artist {
			tracks = append(tracks, t)
		}
	}
	return tracks
}

// GetTrackByIndex returns the track at the given index.
func (p *Playlist) GetTrackByIndex(index int) *Track {
	if index >= 0 && index < len(p.Tracks) {
		return &p.Tracks[index]
	}
	return nil
}
