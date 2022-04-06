package gods

import "errors"

//Queue represents an abstract FIFO queue structure with push, pop, and peek operations.
type Queue[T any] struct {
	List[T]
}

//NewQueue returns a pointer to an initialized queue with zero elements.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

//ErrQueueEmpty is returned when a pop operation is performed on an empty queue.
var ErrQueueEmpty = errors.New("cannot pop from an empty queue")

//Push pushes the provided value to the end of the queue
func (q *Queue[T]) Push(item T) {
	q.InsertBack(item)
}

//Pop removes and returns the item in the front of the queue, or `ErrQueueEmpty` if the queue is empty.
func (q *Queue[T]) Pop() (T, error) {
	var item T
	if q.Front() == nil {
		return item, ErrQueueEmpty
	}
	item = q.Front().Value
	q.Remove(q.Front())
	return item, nil
}

//Peek returns the item in the front of the queue without removing it, or `ErrQueueEmpty` if the queue is empty.
func (q *Queue[T]) Peek() (T, error) {
	if q.Front() == nil {
		var item T
		return item, ErrQueueEmpty
	}

	return q.Front().Value, nil
}
