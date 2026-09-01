package main

import (
	"fmt"
	"strings"
)

type ArrayList struct {
	arrayList []Track
	size      int
}

// CreateArrayList creats an array list with default of capacity 50.
func CreateArrayList(capacity ...int) *ArrayList {
	cap := 50
	if len(capacity) > 0 {
		cap = capacity[0]
	}
	return &ArrayList{
		arrayList: make([]Track, cap),
		size:      0,
	}
}

// incrementSize add 1 to the size of ArrayList.
func (a *ArrayList) incrementSize() {
	a.size++
}

// decrementSize lessen the size of ArrayList by 1.
func (a *ArrayList) decrementSize() {
	a.size--
}

// AddFirst adds a new track to the front, making it the topmost of first element.
func (a *ArrayList) AddFirst(track Track) {
	a.arrayList = append([]Track{track}, a.arrayList...)
	a.incrementSize()
}

// AddLast adds a new track to the end, making it the bottommost or last element.
func (a *ArrayList) AddLast(track Track) {
	a.arrayList = append(a.arrayList, track)
	a.incrementSize()
}

// Add adds a new track to the specified index.
func (a *ArrayList) Add(track Track, idx int) {
	a.arrayList = append(a.arrayList[:idx], append([]Track{track}, a.arrayList[idx:]...)...)
	a.incrementSize()
}

// RemoveAtIndex removes the element at a specified index.
func (a *ArrayList) RemoveAtIndex(idx int) Track {
	target := a.arrayList[idx]
	a.arrayList = append(a.arrayList[:idx], a.arrayList[idx+1:]...)
	a.decrementSize()
	return target
}

// Remove removes the first occurence of track from the ArrayList. Returns the removed track, or none if not found.
func (a *ArrayList) Remove(track Track) Track {
	var target Track
	for i, t := range a.arrayList {
		if t == track {
			target = t
			a.RemoveAtIndex(i)
			break
		}
	}
	return target
}

// Get gets the element at the specified index.
func (a *ArrayList) Get(idx int) Track {
	return a.arrayList[idx]
}

// GetSize returns the size of this ArrayList.
func (a *ArrayList) GetSize() int {
	return a.size
}

// String returns a string representation of the ArrayList.
func (a *ArrayList) String() string {
	if a.size == 0 {
		return "This list is empty..."
	}
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < a.size; i++ {
		sb.WriteString(fmt.Sprintf("%v", a.arrayList[i]))
		if i+i != a.size {
			sb.WriteString(",")
		}
	}
	sb.WriteString("]")
	return sb.String()
}
