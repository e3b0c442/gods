package gods

import "testing"

func checkStack[T comparable](st *Stack[T], ch []T) bool {
	if len(*st) != len(ch) {
		return false
	}

	for i := range *st {
		if (*st)[i] != ch[i] {
			return false
		}
	}

	return true
}
func TestStack(t *testing.T) {
	stack := NewStack[int]()
	check := []int{}

	t.Run("new stack", func(t *testing.T) {
		if !checkStack(stack, check) {
			t.Fatalf("got %#v, expected %#v", stack, check)
		}
	})

	t.Run("push 1", func(t *testing.T) {
		check = []int{1}
		stack.Push(1)

		if !checkStack(stack, check) {
			t.Fatalf("got %#v, expected %#v", stack, check)
		}
	})

	t.Run("peek 1", func(t *testing.T) {
		v, err := stack.Peek()
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		if v != 1 {
			t.Fatalf("got %d expected 1", v)
		}
		if !checkStack(stack, check) {
			t.Fatalf("got %#v, expected %#v", stack, check)
		}
	})

	t.Run("push 2", func(t *testing.T) {
		check = []int{1, 2}
		stack.Push(2)

		if !checkStack(stack, check) {
			t.Fatalf("got %#v, expected %#v", stack, check)
		}
	})

	t.Run("pop 2", func(t *testing.T) {
		check = []int{1}
		val, err := stack.Pop()
		if err != nil {
			t.Fatalf("got unexpected error %v", err)
		}
		if val != 2 {
			t.Fatalf("pop returned %d, expected 2", val)
		}
		if !checkStack(stack, check) {
			t.Fatalf("got %#v, expected %#v", stack, check)
		}
	})

	check = []int{}
	t.Run("pop 1", func(t *testing.T) {
		val, err := stack.Pop()
		if err != nil {
			t.Fatalf("got unexpected error %v", err)
		}
		if val != 1 {
			t.Fatalf("pop returned %d, expected 1", val)
		}
		if !checkStack(stack, check) {
			t.Fatalf("got %#v, expected %#v", stack, check)
		}
	})

	t.Run("peek empty", func(t *testing.T) {
		_, err := stack.Peek()
		if err == nil {
			t.Fatal("did not get expected error")
		}
		if !checkStack(stack, check) {
			t.Fatalf("got %#v, expected %#v", stack, check)
		}
	})

	t.Run("pop empty", func(t *testing.T) {
		val, err := stack.Pop()
		if err == nil {
			t.Fatal("expected error from empty stack")
		}
		if val != 0 {
			t.Fatal("empty stack pop should have returned zero value")
		}
		if !checkStack(stack, check) {
			t.Fatalf("got %#v, expected %#v", stack, check)
		}
	})
}
