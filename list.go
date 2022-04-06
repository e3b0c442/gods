package gods

import "errors"

//List represents a doubly-linked list.
type List[T any] struct {
	front, back *Node[T]
}

//Front returns the Node at the front of the list, or nil if the list is empty.
func (l *List[T]) Front() *Node[T] {
	return l.front
}

//Back returns the Node at the back of the list, or nil if the list is empty.
func (l *List[T]) Back() *Node[T] {
	return l.back
}

//Node represents a node in a linked list. It has a value and pointers to the next and previous items in the list.
type Node[T any] struct {
	Value      T
	prev, next *Node[T]
}

//Next returns the next Node in the List.
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

//Prev returns the previous Node in the List.
func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

//NewList returns an initialized list with zero elements.
func NewList[T any]() *List[T] {
	return &List[T]{}
}

//ErrNilNode is returned when nil is provided but a valid Node pointer is required.
var ErrNilNode = errors.New("provided node cannot be nil")

//InsertAfter inserts a Node in the List after the provided anchor. If anchor is nil, the value is inserted at the front
//of the list.
func (l *List[T]) InsertAfter(anchor *Node[T], val T) {
	node := &Node[T]{
		Value: val,
	}

	if anchor == nil { //inserting at beginning of list
		node.prev = nil
		node.next = l.front

		l.front = node
	} else { // inserting after existing node
		node.prev = anchor
		node.next = anchor.next

		node.prev.next = node
	}
	if node.next != nil {
		node.next.prev = node
	} else {
		l.back = node
	}
}

//InsertFront inserts a Node at the front of the List. It is equivalent to `InsertAfter(nil, value)`.
func (l *List[T]) InsertFront(val T) {
	l.InsertAfter(nil, val)
}

//InsertBefore inserts a Node in the List before the provided anchor. If anchor is nil, the value is inserted at the
//back of the List.
func (l *List[T]) InsertBefore(anchor *Node[T], val T) {
	node := &Node[T]{
		Value: val,
	}

	if anchor == nil { //inserting at end of list
		node.prev = l.back
		node.next = nil

		l.back = node
	} else { //inserting before existing node
		node.prev = anchor.prev
		node.next = anchor

		node.next.prev = node
	}
	if node.prev != nil {
		node.prev.next = node
	} else {
		l.front = node
	}
}

//InsertBack inserts a Node at the back of the List. It is equivalent to `InsertBefore(nil, value)`.
func (l *List[T]) InsertBack(val T) {
	l.InsertBefore(nil, val)
}

//Remove removes the provided Node from the list. If the list is empty, `ErrNilNode` is returned.
func (l *List[T]) Remove(node *Node[T]) error {
	if node == nil {
		return ErrNilNode
	}

	if node.prev == nil { //is at the head
		l.front = node.next
	} else {
		node.prev.next = node.next
	}

	if node.next == nil { // is at tail
		l.back = node.prev
	} else {
		node.next.prev = node.prev
	}

	return nil
}
