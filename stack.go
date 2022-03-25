package gods

import "errors"

//Stack represents an abstract stack with push and pop operations.
type Stack[T any] []T

//NewStack returns an initialized stack with zero elements.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

//Push adds an item to the top of the stack.
func (st *Stack[T]) Push(item T) {
	*st = append(*st, item)
}

//ErrStackEmpty is returned when a pop operation is performed on an empty stack.
var ErrStackEmpty = errors.New("cannot pop from an empty stack")

//Pop removes the item on top of the stack and returns it. If the stack is empty, ErrStackEmpty is returned.
func (st *Stack[T]) Pop() (T, error) {
	if len(*st) < 1 {
		var t T
		return t, ErrStackEmpty
	}
	it := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return it, nil
}
