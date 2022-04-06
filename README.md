# GoDS

This is a repository of data structures built with the new generics functionality in Go 1.18.

This is not necessarily meant for public consumption, but rather as a learning exercise. That said, if you find use in these, mores the better.

Structures are built on top of idiomatic Go types; e.g. slices rather than arrays since array capacities are part of the type and it is not possible to make an open-ended generic-length array, and use idiomatic error returns.

# Structures

## List

`List` implements a doubly-linked list. It is created with the `NewList[T]()` function, which returns a pointer to a new, empty list. Items in the list are the `Node` type, which encapsulates the value.

### Node

`Node represents a linked list node. It has the following methods:

- `func (n *Node[T]) Next() *Node[T]`: returns the node behind the receiver, or nil if the receiver is the back of the list.
- `func (n *Node[T]) Prev() *Node[T]`: returns the node in front of the receiver, or nil if the recevier is the front of the list.

`List` supports the following operations:

- `func (*List[T]) InsertBefore(*Node[T], T)`: adds a value before the provided node in the list. If the node pointer is nil, adds the value to the back of the list.
- `func (*List[T]) InsertAfter(*Node[T], T)`: adds a value after the provided node in the list. If the node pointer is nil, adds the value to the front of the list.
- `func (*List[T]) InsertFront(T)`: adds a value to the front of the list. Equivalent to `InsertAfter(nil, T)`
- `func (*List[T]) InsertBack(T)`: adds a value to the back of the list. Equivalent to `InsertBefore(nil, T)`
- `func (*List[T]) Remove(*Node[T]) error`: removes a node from the list, repointing the surrounding nodes accordingly. Returns `ErrNilNode` if the provided node pointer is nil.

## Stack

`Stack` implements a generic LIFO stack data structure backed by a slice. It is created with the `NewStack[T]()` function, which returns a pointer to a new, empty stack. It has supports following operations:

- `func (*Stack[T]) Push(T)`: push an item onto the top of the stack
- `func (*Stack[T]) Peek() (T, error)`: return the value on top of the stack without removing it. Returns `ErrStackEmpty` if the stack is empty.
- `func (*Stack[T]) Pop() (T, error)`: remove the value on top of the stack and return it. Returns `ErrStackEmpty` if the stack is empty.

## Queue

`Queue` implements a generic FIFO queue data structure backed by a `List`. It is created with the `NewQueue[T]()` function, which returns a pointer to a new, empty queue. It supports the following operations:

- `func (*Queue[T]) Push(T)`: pushes an item to the back of the queue.
- `func (*Queue[T]) Peek() (T, error)`: returns the item at the front of the queue, or `ErrQueueEmpty` if the queue is empty.
- `func (*Queue[T]) Pop() (T, error)`: remove the value at the front of the queue and return it. Returns `ErrQueueEmpty` if the queue is empty.

## Deque

`Deque` implements a generic double-ended queue backed by a slice. `Deque` attempts to be efficient and uses the slice as a ring buffer to avoid excessive reallocations. It is created with the `NewDeque[T]()` function, which returns a pointer to a new, empty deque with a minimum capacity of 8. The minimum capacity can be customized through the use of `NewDequeWithCap[T](int)`. Passing a capacity less than 1 will panic. `Deque` supports the following operations:

- `func (*Deque[T]) PushBack(T)`: pushes an item to the back of the deque.
- `func (*Deque[T]) PushFront(T)`: pushes an item to the front of the deque.
- `func (*Deque[T]) PopBack() (T, error)`: remove the value at the back of the deque and return it. Returns `ErrDequeEmpty` if the deque is empty.
- `func (*Deque[T]) PopFront() (T, error)`: remove the value at the back of the deque and return it. Returns `ErrDequeEmpty` if the deque is empty.
- `func (*Deque[T]) PeekBack() (T, error)`: return the value at the back of the deque without removing it. Returns `ErrDequeEmpty` if the deque is empty.
- `func (*Deque[T]) PeekFront() (T, error)`: return the value at the front of the deque without removing it. Returns `ErrDequeEmpty` if th deque is empty.
