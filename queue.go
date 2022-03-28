package gods

import "errors"

//Queue represents an abstract FIFO queue structure with push and pop operations.
type Queue[T any] struct {
	List[T]
}

//NewQueue returns an initialized queue with zero elements.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

//ErrQueueEmpty is returned when a pop operation is performed on an empty queue.
var ErrQueueEmpty = errors.New("cannot pop from an empty queue")

//Push pushes the provided value to the end of the queue
func (q *Queue[T]) Push(item T) {
	q.InsertBack(item)
}

//Pop returns the item in the front of the queue, or an error if the queue is empty.
func (q *Queue[T]) Pop() (T, error) {
	var item T
	if q.Front() == nil {
		return item, ErrQueueEmpty
	}
	item = q.Front().Value
	q.Remove(q.Front())
	return item, nil
}
