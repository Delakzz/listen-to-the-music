package main

import (
	"fmt"
	"strings"
)

type Playlist struct {
	ArrayList
	name     string
	queue    *Queue
	shuffled bool
	playing  bool
	repeat   int // 0 for no repeat. 1 for all repeat. 2 for only one repeat.
}

// Playlist returns the list that holds the playlist of tracks.
func (p *Playlist) Playlist() []Track {
	return p.queue.arrayList
}

// NotEmpty checks if the playlist is not empty.
func (p *Playlist) NotEmpty() bool {
	return p.queue.IsEmpty()
}

// PlaylistName returns the name of the playlist
func (p *Playlist) PlaylistName() string {
	return p.name
}

// IsPlaying checks if this playlist is currently being played.
func (p *Playlist) IsPlaying() bool {
	return p.playing
}

// IsAllRepeat checks if queued playlist is configured to play all in repeat.
func (p *Playlist) IsAllRepeat() bool {
	if p.repeat == 1 {
		return true
	}
	return false
}

// IsOneRepeat checks if queued playlist is configured to only play one track in repeat.
func (p *Playlist) IsOneRepeat() bool {
	if p.repeat == 2 {
		return true
	}
	return false
}

// IsNoRepeat checks if queued playlist has repeat disabled.
func (p *Playlist) IsNoRepeat() bool {
	if p.repeat == 0 {
		return true
	}
	return false
}

// IsShuffled checks if the queued playlist is shuffled or not
func (p *Playlist) IsShuffled() bool {
	return p.shuffled
}

// SetAllRepeat enables all repeat for this playlist.
func (p *Playlist) SetAllRepeat() {
	p.repeat = 1
}

// SetOneRepeat enables only one repeat for this playlist.
func (p *Playlist) SetOneRepeat() {
	p.repeat = 2
}

// NoRepeat disables repeat for this playlist
func (p *Playlist) NoRepeat() {
	p.repeat = 0
}

// setRepeat sets repeat based on given value.
func (p *Playlist) setRepeat(val int) {
	p.repeat = val
}

// EnbaleShuffle enables the shuffle on this playlist.
func (p *Playlist) EnableShuffle() {
	p.shuffled = true
}

// DisableShuffle disables the shuffle on this playlist.
func (p *Playlist) DisableShuffle() {
	p.shuffled = false
}

// setShuffle sets shuffle into the specified value.
func (p *Playlist) setShuffle(val bool) {
	p.shuffled = val
}

// AddToQueue adds a track to the queue of this playlist.
func (p *Playlist) AddToQueue(track Track) {
	p.queue.Enqueue(track)
}

// createPlaylistQueue creats a queue of tracks from a given index from the playlist to the last index, exclusively.
func (p *Playlist) createPlaylistQueue(start, stop int) {
	for _, track := range p.arrayList[start:stop] {
		p.queue.Enqueue(track)
	}
}

// AddToPlaylist adds a track to the playlist.
func (p *Playlist) AddToPlaylist(track Track) {
	p.AddLast(track)
}

// Play plays the playlist from a given track.
func (p *Playlist) Play(track Track) {
	start := p.GetIndexOfTrack(track)
	if start != -1 {
		p.createPlaylistQueue(start, p.GetSize())
	}
}

// StopPlay stops this playlist from playing. Reseets all playlist configurations.
func (p *Playlist) StopPlay() {
	p.queue.Purge()
	p.DisableShuffle()
	p.NoRepeat()
	p.playing = false
}

// PlayingTrack gets the tracj currently being played by the queue of this playlist.
func (p *Playlist) PlayingTrack() string {
	return p.queue.GetHead().String()
}

// GetIndexOfTrack get the index of a specified track on this playlist. -1 if not found.
func (p *Playlist) GetIndexOfTrack(track Track) int {
	for i, val := range p.arrayList {
		if val == track {
			return i
		}
	}
	return -1
}

// NextTrack plays the next track on queue. Stops playing if no more track left.
func (p *Playlist) NextTrack() {
	if p.queue.GetSize() == 1 {
		p.StopPlay()
	}

	track := p.queue.Dequeue()

	if p.IsAllRepeat() {
		p.queue.Enqueue(*track)
	}
}

// PlaylistQueue shows the queue currently being played from this playlist.
func (p *Playlist) PlaylistQueue() string {
	if p.queue.IsEmpty() {
		return "This queue is empty."
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Currently playing: %s", p.queue.GetHead().String()))
	sb.WriteString("<---- Next on Queue ---->")
	for i := 1; i < p.queue.size; i++ {
		sb.WriteString(fmt.Sprintf("%v\n", p.queue.arrayList[i].String()))
	}
	sb.WriteString("<---- End of Queue ---->")
	return sb.String()
}

func (p *Playlist) String() string {
	return p.ArrayList.String()
}
