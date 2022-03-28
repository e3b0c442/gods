package gods

import (
	"errors"
	"testing"
)

func checkQueue[T comparable](q *Queue[T], ch []T) error {
	return checkList(&q.List, ch)
}

func TestQueue(t *testing.T) {
	queue := NewQueue[int]()

	t.Run("empty", func(t *testing.T) {
		check := []int{}

		if err := checkQueue(queue, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("push one", func(t *testing.T) {
		queue.Push(1)
		check := []int{1}

		if err := checkQueue(queue, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("push two", func(t *testing.T) {
		queue.Push(2)
		check := []int{1, 2}

		if err := checkQueue(queue, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("pop one", func(t *testing.T) {
		val, err := queue.Pop()
		if err != nil {
			t.Fatalf("unexpected error popping populated queue: %v", err)
		}
		if val != 1 {
			t.Fatalf("expected value popped, expected 1, got %d", val)
		}
		check := []int{2}

		if err := checkQueue(queue, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("pop two", func(t *testing.T) {
		val, err := queue.Pop()
		if err != nil {
			t.Fatalf("unexpected error popping populated queue: %v", err)
		}
		if val != 2 {
			t.Fatalf("expected value popped, expected 2, got %d", val)
		}
		check := []int{}

		if err := checkQueue(queue, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("pop empty", func(t *testing.T) {
		_, err := queue.Pop()
		if !errors.Is(err, ErrQueueEmpty) {
			t.Fatalf("expected empty queue error, got %v", err)
		}
	})
}
