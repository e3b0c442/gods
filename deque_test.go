package gods

import (
	"fmt"
	"testing"
)

func checkDeque[T comparable](check, control *Deque[T]) error {
	if !Equal(check.buf, control.buf) {
		return fmt.Errorf("unexpected deque buffer: got %#v expected %#v", check.buf, control.buf)
	}
	if check.front != control.front {
		return fmt.Errorf("unexpected deque front: got %d expected %d", check.front, control.front)
	}
	if check.back != control.back {
		return fmt.Errorf("unexpected deque back: got %d expected %d", check.back, control.back)
	}
	if check.len != control.len {
		return fmt.Errorf("unexpected deque len: got %d expected %d", check.len, control.len)
	}
	if check.cap != control.cap {
		return fmt.Errorf("unexpected deque cap: got %d expected %d", check.cap, control.cap)
	}
	return nil
}

func TestDeque(t *testing.T) {
	d := NewDeque[int]()

	buf := make([]int, defaultCap)
	control := &Deque[int]{
		buf:   buf,
		front: -1,
		back:  -1,
		len:   0,
		cap:   defaultCap,
	}

	t.Run("new default", func(t *testing.T) {
		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	d = NewDequeWithCap[int](1)
	buf = make([]int, 1)
	control.buf = buf
	control.len = 0
	control.cap = 1
	t.Run("new cap 1", func(t *testing.T) {
		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("pushback 1", func(t *testing.T) {
		d.PushBack(1)
		control.buf[0] = 1
		control.front = 0
		control.back = 0
		control.len = 1
		control.cap = 1

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("peekback 1", func(t *testing.T) {
		v, err := d.PeekBack()
		if err != nil {
			t.Fatal(err)
		}
		if v != 1 {
			t.Fatalf("peekback got %d expected 1", v)
		}
		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("peekfront 1", func(t *testing.T) {
		v, err := d.PeekFront()
		if err != nil {
			t.Fatal(err)
		}
		if v != 1 {
			t.Fatalf("peekfront got %d expected 1", v)
		}
		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popback 1", func(t *testing.T) {
		v, err := d.PopBack()
		if err != nil {
			t.Fatal(err)
		}
		if v != 1 {
			t.Fatalf("popback got %d expected 1", v)
		}
		control.front = -1
		control.back = -1
		control.len = 0

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popback empty", func(t *testing.T) {
		_, err := d.PopBack()
		if err != ErrDequeEmpty {
			t.Fatalf("did not get expected empty deque error, got %v", err)
		}

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("peekback empty", func(t *testing.T) {
		_, err := d.PeekBack()
		if err != ErrDequeEmpty {
			t.Fatalf("did not get expected empty deque error, got %v", err)
		}

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("peekfront empty", func(t *testing.T) {
		_, err := d.PeekFront()
		if err != ErrDequeEmpty {
			t.Fatalf("did not get expected empty deque error, got %v", err)
		}

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("shift 2", func(t *testing.T) {
		d.PushFront(2)
		control.buf[0] = 2
		control.front = 0
		control.back = 0
		control.len = 1
		control.cap = 1

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("peekback 2", func(t *testing.T) {
		v, err := d.PeekBack()
		if err != nil {
			t.Fatal(err)
		}
		if v != 2 {
			t.Fatalf("peekback got %d expected 2", v)
		}
		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("peekfront 2", func(t *testing.T) {
		v, err := d.PeekFront()
		if err != nil {
			t.Fatal(err)
		}
		if v != 2 {
			t.Fatalf("peekfront got %d expected 2", v)
		}
		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popfront 2", func(t *testing.T) {
		v, err := d.PopFront()
		if err != nil {
			t.Fatal(err)
		}
		if v != 2 {
			t.Fatalf("popfront got %d expected 2", v)
		}
		control.front = -1
		control.back = -1
		control.len = 0

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popfront empty", func(t *testing.T) {
		_, err := d.PopFront()
		if err != ErrDequeEmpty {
			t.Fatalf("did not get expected empty deque error, got %v", err)
		}

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("pushback 20", func(t *testing.T) {
		d.PushBack(20)
		control.buf[0] = 20
		control.front = 0
		control.back = 0
		control.len = 1
		control.cap = 1

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("pushfront 10 (expand)", func(t *testing.T) {
		d.PushFront(10)
		control.buf = []int{20, 10}
		control.front = 1
		control.back = 0
		control.len = 2
		control.cap = 2

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("peekfront 10", func(t *testing.T) {
		v, err := d.PeekFront()
		if err != nil {
			t.Fatal(err)
		}
		if v != 10 {
			t.Fatalf("peekfront got %d expected 10", v)
		}
		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("peekback 20", func(t *testing.T) {
		v, err := d.PeekBack()
		if err != nil {
			t.Fatal(err)
		}
		if v != 20 {
			t.Fatalf("peekback got %d expected 20", v)
		}
		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popback 20", func(t *testing.T) {
		v, err := d.PopBack()
		if err != nil {
			t.Fatal(err)
		}
		if v != 20 {
			t.Fatalf("popback got %d expected 20", v)
		}
		control.front = 1
		control.back = 1
		control.len = 1
		control.cap = 2

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("pushback 30", func(t *testing.T) {
		d.PushBack(30)
		control.buf[0] = 30
		control.front = 1
		control.back = 0
		control.len = 2
		control.cap = 2

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popfront 10", func(t *testing.T) {
		v, err := d.PopFront()
		if err != nil {
			t.Fatal(err)
		}
		if v != 10 {
			t.Fatalf("popback got %d expected 10", v)
		}
		control.front = 0
		control.back = 0
		control.len = 1
		control.cap = 2

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("pushfront 20", func(t *testing.T) {
		d.PushFront(20)
		control.buf[1] = 20
		control.front = 1
		control.back = 0
		control.len = 2
		control.cap = 2

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("pushfront 10 (expand again)", func(t *testing.T) {
		d.PushFront(10)
		control.buf = []int{20, 30, 10}
		control.front = 2
		control.back = 1
		control.len = 3
		control.cap = 3

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("peekfront 3 10", func(t *testing.T) {
		v, err := d.PeekFront()
		if err != nil {
			t.Fatal(err)
		}
		if v != 10 {
			t.Fatalf("peekfront got %d expected 10", v)
		}
		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("peekback 3 30", func(t *testing.T) {
		v, err := d.PeekBack()
		if err != nil {
			t.Fatal(err)
		}
		if v != 30 {
			t.Fatalf("peekback got %d expected 30", v)
		}
		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popback 30", func(t *testing.T) {
		v, err := d.PopBack()
		if err != nil {
			t.Fatal(err)
		}
		if v != 30 {
			t.Fatalf("popback got %d expected 20", v)
		}
		control.front = 2
		control.back = 0
		control.len = 2
		control.cap = 3

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("peekfront 2 10", func(t *testing.T) {
		v, err := d.PeekFront()
		if err != nil {
			t.Fatal(err)
		}
		if v != 10 {
			t.Fatalf("peekfront got %d expected 10", v)
		}
		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("peekback 2 20", func(t *testing.T) {
		v, err := d.PeekBack()
		if err != nil {
			t.Fatal(err)
		}
		if v != 20 {
			t.Fatalf("peekback got %d expected 20", v)
		}
		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popfront 10 (contract)", func(t *testing.T) {
		v, err := d.PopFront()
		if err != nil {
			t.Fatal(err)
		}
		if v != 10 {
			t.Fatalf("popback got %d expected 10", v)
		}
		control.buf = []int{20, 0}
		control.front = 0
		control.back = 0
		control.len = 1
		control.cap = 2

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("peekback 1 20", func(t *testing.T) {
		v, err := d.PeekBack()
		if err != nil {
			t.Fatal(err)
		}
		if v != 20 {
			t.Fatalf("peekback got %d expected 20", v)
		}
		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("peekfront 1 20", func(t *testing.T) {
		v, err := d.PeekFront()
		if err != nil {
			t.Fatal(err)
		}
		if v != 20 {
			t.Fatalf("peekfront got %d expected 20", v)
		}
		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popback 1 20", func(t *testing.T) {
		v, err := d.PopBack()
		if err != nil {
			t.Fatal(err)
		}
		if v != 20 {
			t.Fatalf("popback got %d expected 20", v)
		}
		control.front = -1
		control.back = -1
		control.len = 0
		control.cap = 2

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("pushback 8", func(t *testing.T) {
		d.PushBack(8)
		control.buf[0] = 8
		control.front = 0
		control.back = 0
		control.len = 1
		control.cap = 2

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("pushfront 5", func(t *testing.T) {
		d.PushFront(5)
		control.buf[1] = 5
		control.front = 1
		control.back = 0
		control.len = 2
		control.cap = 2

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("pushback 13", func(t *testing.T) {
		d.PushBack(13)
		control.buf = []int{5, 8, 13}
		control.front = 0
		control.back = 2
		control.len = 3
		control.cap = 3

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("pushfront 3", func(t *testing.T) {
		d.PushFront(3)
		control.buf = []int{5, 8, 13, 0, 3}
		control.front = 4
		control.back = 2
		control.len = 4
		control.cap = 5

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("pushback 21", func(t *testing.T) {
		d.PushBack(21)
		control.buf[3] = 21
		control.front = 4
		control.back = 3
		control.len = 5
		control.cap = 5

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("pushfront 2", func(t *testing.T) {
		d.PushFront(2)
		control.buf = []int{3, 5, 8, 13, 21, 0, 0, 2}
		control.front = 7
		control.back = 4
		control.len = 6
		control.cap = 8

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("pushback 34", func(t *testing.T) {
		d.PushBack(34)
		control.buf[5] = 34
		control.front = 7
		control.back = 5
		control.len = 7
		control.cap = 8

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("pushfront 1", func(t *testing.T) {
		d.PushFront(1)
		control.buf[6] = 1
		control.front = 6
		control.back = 5
		control.len = 8
		control.cap = 8

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("pushback 55", func(t *testing.T) {
		d.PushBack(55)
		control.buf = []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 0, 0, 0, 0}
		control.front = 0
		control.back = 8
		control.len = 9
		control.cap = 13

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popfront 1", func(t *testing.T) {
		v, err := d.PopFront()
		if err != nil {
			t.Fatal(err)
		}
		if v != 1 {
			t.Fatalf("popback got %d expected 1", v)
		}
		control.front = 1
		control.back = 8
		control.len = 8
		control.cap = 13

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popfront 2", func(t *testing.T) {
		v, err := d.PopFront()
		if err != nil {
			t.Fatal(err)
		}
		if v != 2 {
			t.Fatalf("popback got %d expected 2", v)
		}
		control.front = 2
		control.back = 8
		control.len = 7
		control.cap = 13

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popfront 3", func(t *testing.T) {
		v, err := d.PopFront()
		if err != nil {
			t.Fatal(err)
		}
		if v != 3 {
			t.Fatalf("popback got %d expected 3", v)
		}
		control.front = 3
		control.back = 8
		control.len = 6
		control.cap = 13

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popfront 5 (contract)", func(t *testing.T) {
		v, err := d.PopFront()
		if err != nil {
			t.Fatal(err)
		}
		if v != 5 {
			t.Fatalf("popback got %d expected 5", v)
		}
		control.buf = []int{8, 13, 21, 34, 55, 0, 0, 0}
		control.front = 0
		control.back = 4
		control.len = 5
		control.cap = 8

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popback 55", func(t *testing.T) {
		v, err := d.PopBack()
		if err != nil {
			t.Fatal(err)
		}
		if v != 55 {
			t.Fatalf("popback got %d expected 55", v)
		}
		control.front = 0
		control.back = 3
		control.len = 4
		control.cap = 8

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popback 34 (contract)", func(t *testing.T) {
		v, err := d.PopBack()
		if err != nil {
			t.Fatal(err)
		}
		if v != 34 {
			t.Fatalf("popback got %d expected 34", v)
		}
		control.buf = []int{8, 13, 21, 0, 0}
		control.front = 0
		control.back = 2
		control.len = 3
		control.cap = 5

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popfront 8 (contract)", func(t *testing.T) {
		v, err := d.PopFront()
		if err != nil {
			t.Fatal(err)
		}
		if v != 8 {
			t.Fatalf("popback got %d expected 8", v)
		}
		control.buf = []int{13, 21, 0}
		control.front = 0
		control.back = 1
		control.len = 2
		control.cap = 3

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popback 21 (contract)", func(t *testing.T) {
		v, err := d.PopBack()
		if err != nil {
			t.Fatal(err)
		}
		if v != 21 {
			t.Fatalf("popback got %d expected 21", v)
		}
		control.buf = []int{13, 0}
		control.front = 0
		control.back = 0
		control.len = 1
		control.cap = 2

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popfront 13", func(t *testing.T) {
		v, err := d.PopFront()
		if err != nil {
			t.Fatal(err)
		}
		if v != 13 {
			t.Fatalf("popback got %d expected 13", v)
		}
		control.front = -1
		control.back = -1
		control.len = 0
		control.cap = 2

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popback empty fib", func(t *testing.T) {
		_, err := d.PopBack()
		if err != ErrDequeEmpty {
			t.Fatalf("did not get expected empty deque error, got %v", err)
		}

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("popfront empty fib", func(t *testing.T) {
		_, err := d.PopFront()
		if err != ErrDequeEmpty {
			t.Fatalf("did not get expected empty deque error, got %v", err)
		}

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("peekback empty fib", func(t *testing.T) {
		_, err := d.PeekBack()
		if err != ErrDequeEmpty {
			t.Fatalf("did not get expected empty deque error, got %v", err)
		}

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("peekfront empty fib", func(t *testing.T) {
		_, err := d.PeekFront()
		if err != ErrDequeEmpty {
			t.Fatalf("did not get expected empty deque error, got %v", err)
		}

		if err := checkDeque(d, control); err != nil {
			t.Fatal(err)
		}
	})
}
