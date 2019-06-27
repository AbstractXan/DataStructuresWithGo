package queue

import l "list"

// Queue implements a basic queue
// Enqueue / Dequeue
type Queue struct {
	list l.List
}

// Enqueue enqueues a value
func (q *Queue) Enqueue(value int) {
	q.list.InsertAt(value, 1)
}

// Dequeue dequeues a value
func (q *Queue) Dequeue() int {
	dequeued := q.Back()
	q.list.DeleteLast()
	return dequeued
}

// Front returns value at front
func (q *Queue) Front() int {
	return q.list.GetVal(1)
}

// Back returns value at back
func (q *Queue) Back() int {
	return q.list.GetVal(q.list.Length())
}
