package gods

import (
	"errors"
	"math"
)

const (
	defaultCap = 8
)

//Deque represents an abstract double-ended queue supporting push, pop, and peek operations on both ends.
type Deque[T any] struct {
	buf                        []T
	front, back, len, cap, min int
}

//expand expands the slice backing the deque using the golden ratio. Yes, this is a personal quirk and not the most
//efficient way to expand.
func (d *Deque[T]) expand() {
	newCap := int(math.Round(float64(d.cap) * math.Phi))
	newBuf := make([]T, newCap)
	i := 0
	for j := d.front; j < d.front+d.len; j++ {
		newBuf[i] = d.buf[j%d.cap]
		i++
	}
	d.buf = newBuf
	d.cap = newCap
	d.front = 0
	d.back = d.len - 1
}

//contract contracts the slice backing the deque using the golden ratio. Yes, this is a personal quirk and not the most
//efficient way to contract.
func (d *Deque[T]) contract() {
	newCap := int(math.Round(float64(d.cap) / math.Phi))
	if newCap < d.min {
		return
	}
	newBuf := make([]T, newCap)
	i := 0
	for j := d.front; j < d.front+d.len; j++ {
		newBuf[i] = d.buf[j%d.cap]
		i++
	}
	d.buf = newBuf
	d.cap = newCap
	d.front = 0
	d.back = d.len - 1
}

//NewDeque returns a pointer to an initialized deque with zero elements and a minimum capacity of 8.
func NewDeque[T any]() *Deque[T] {
	return NewDequeWithCap[T](defaultCap)
}

//NewDequeWithCap returns a pointer to an initialized deque with zero elements and a minimum capacity of `cap`. If cap
//is less than 1, NewDequeWithCap will panic.
func NewDequeWithCap[T any](cap int) *Deque[T] {
	if cap < 1 {
		panic("attempt to create deque with capacity < 1")
	}
	deque := &Deque[T]{
		cap:   cap,
		min:   cap,
		front: -1,
		back:  -1,
	}
	deque.buf = make([]T, cap)
	return deque
}

//ErrDequeEmpty is returned when a pop or peek operation is performed on an empty deque.
var ErrDequeEmpty = errors.New("cannot pop from an empty deque")

//PeekBack returns the item at the back of the deque without removing it. Returns `ErrDequeEmpty` if the deque is empty.
func (d *Deque[T]) PeekBack() (T, error) {
	if d.len == 0 {
		var t T
		return t, ErrDequeEmpty
	}
	return d.buf[d.back], nil
}

//PeekFront returns the item at the front of the deque without removing it. Returns `ErrDequeEmpty` if the deque is
//empty.
func (d *Deque[T]) PeekFront() (T, error) {
	if d.len == 0 {
		var t T
		return t, ErrDequeEmpty
	}
	return d.buf[d.front], nil
}

//PushBack pushes the provided value to the back of the deque.
func (d *Deque[T]) PushBack(val T) {
	if d.len == d.cap {
		d.expand()
	}

	if d.len == 0 {
		d.front = 0
		d.back = 0
	} else {
		d.back = (d.back + 1) % d.cap
	}
	d.buf[d.back] = val
	d.len += 1
}

//PushFront pushes the provided value to the front of the deque.
func (d *Deque[T]) PushFront(val T) {
	if d.len == d.cap {
		d.expand()
	}

	if d.len == 0 {
		d.front = 0
		d.back = 0
	} else {
		d.front = (d.front - 1) % d.cap
		if d.front < 0 {
			d.front = d.cap + d.front
		}
	}
	d.buf[d.front] = val
	d.len += 1
}

//PopBack removes the value at the back of the deque and returns it. Returns `ErrDequeEmpty` if the deque is empty.
func (d *Deque[T]) PopBack() (T, error) {
	if d.len == 0 {
		var t T
		return t, ErrDequeEmpty
	}
	last := d.buf[d.back]
	if d.len == 1 {
		d.front = -1
		d.back = -1
	} else {
		d.back = (d.back - 1) % d.cap
		if d.back < 0 {
			d.back = d.cap + d.back
		}
	}
	d.len--
	if d.len > 0 && float64(d.cap)/float64(d.len) > (1.0+math.Sqrt2) {
		d.contract()
	}
	return last, nil
}

//PopFront removes the value at the front of the deque and returns it. Returns `ErrDequeEmpty` if the deque is empty.
func (d *Deque[T]) PopFront() (T, error) {
	if d.len == 0 {
		var t T
		return t, ErrDequeEmpty
	}

	first := d.buf[d.front]
	if d.len == 1 {
		d.front = -1
		d.back = -1
	} else {
		d.front = (d.front + 1) % d.cap
	}
	d.len--
	if d.len > 0 && float64(d.cap)/float64(d.len) > (1.0+math.Sqrt2) {
		d.contract()
	}
	return first, nil
}
