package gods

import "errors"

type List[T any] struct {
	head, tail *Node[T]
}

func (l *List[T]) Head() *Node[T] {
	return l.head
}

func (l *List[T]) Tail() *Node[T] {
	return l.tail
}

type Node[T any] struct {
	Value      T
	prev, next *Node[T]
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

var (
	ErrNilNode = errors.New("provided node cannot be nil")
)

func (l *List[T]) InsertAfter(anchor *Node[T], val T) {
	node := &Node[T]{
		Value: val,
	}

	if anchor == nil { //inserting at beginning of list
		node.prev = nil
		node.next = l.head

		l.head = node
	} else { // inserting after existing node
		node.prev = anchor
		node.next = anchor.next

		node.prev.next = node
	}
	if node.next != nil {
		node.next.prev = node
	} else {
		l.tail = node
	}
}

func (l *List[T]) InsertBefore(anchor *Node[T], val T) {
	node := &Node[T]{
		Value: val,
	}

	if anchor == nil { //inserting at end of list
		node.prev = l.tail
		node.next = nil

		l.tail = node
	} else { //inserting before existing node
		node.prev = anchor.prev
		node.next = anchor

		node.next.prev = node
	}
	if node.prev != nil {
		node.prev.next = node
	} else {
		l.head = node
	}
}

func (l *List[T]) Remove(node *Node[T]) error {
	if node == nil {
		return ErrNilNode
	}

	if node.prev == nil { //is at the head
		l.head = node.next
	} else {
		node.prev.next = node.next
	}

	if node.next == nil { // is at tail
		l.tail = node.prev
	} else {
		node.next.prev = node.prev
	}

	return nil
}
