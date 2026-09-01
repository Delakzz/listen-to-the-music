package main

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
