package main

var MainQueue = CreateQueue()

type Queue struct {
	ArrayList
}

// CreateQueue creates an empty queue backed by an ArrayList.
func CreateQueue() *Queue {
	return &Queue{
		ArrayList: *CreateArrayList(),
	}
}

// GetHead gets the first track on this queue.
func (q *Queue) GetHead() *Track {
	if q.IsEmpty() {
		return nil
	}
	return q.Get(0)
}

// GetTail gets the last track on this queue.
func (q *Queue) GetTail() *Track {
	if q.IsEmpty() {
		return nil
	}
	return q.Get(q.GetSize() - 1)
}

// Enqueue adds a new track at the end of queue.
func (q *Queue) Enqueue(track Track) {
	q.AddLast(track)
}

// Dequeue removes and returns the track at the front of the queue, or nil if empty.
func (q *Queue) Dequeue() *Track {
	if q.IsEmpty() {
		return nil
	}
	head := *q.Get(0)
	q.RemoveAtIndex(0)
	return &head
}

// Purge removes all tracks on this queue.
func (q *Queue) Purge() {
	q.Clear()
}
