package gods

import "errors"

//Stack represents an abstract stack with push, pop, and peek operations.
type Stack[T any] []T

//NewStack returns a poniter to an initialized stack with zero elements.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

//ErrStackEmpty is returned when a pop or peek operation is performed on an empty stack.
var ErrStackEmpty = errors.New("cannot pop from an empty stack")

//Push adds an item to the top of the stack.
func (st *Stack[T]) Push(item T) {
	*st = append(*st, item)
}

//Pop removes the item on top of the stack and returns it. If the stack is empty, ErrStackEmpty is returned.
func (st *Stack[T]) Pop() (T, error) {
	if len(*st) == 0 {
		var t T
		return t, ErrStackEmpty
	}
	it := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return it, nil
}

//Peek returns the item at the top of the stack without removing it. If the stack is empty, ErrStackEmpty is returned.
func (st *Stack[T]) Peek() (T, error) {
	if len(*st) == 0 {
		var t T
		return t, ErrStackEmpty
	}

	return (*st)[len(*st)-1], nil
}
