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

	if !checkStack(stack, check) {
		t.Fatalf("got %#v, expected %#v", stack, check)
	}

	check = []int{1}
	stack.Push(1)

	if !checkStack(stack, check) {
		t.Fatalf("got %#v, expected %#v", stack, check)
	}

	check = []int{1, 2}
	stack.Push(2)

	if !checkStack(stack, check) {
		t.Fatalf("got %#v, expected %#v", stack, check)
	}

	check = []int{1}
	val, err := stack.Pop()
	if err != nil {
		t.Fatalf("got unexpected error %v", err)
	}
	if val != 2 {
		t.Fatalf("pop returned %d, expected 2", val)
	}

	check = []int{}
	val, err = stack.Pop()
	if err != nil {
		t.Fatalf("got unexpected error %v", err)
	}
	if val != 1 {
		t.Fatalf("pop returned %d, expected 1", val)
	}

	val, err = stack.Pop()
	if err == nil {
		t.Fatal("expected error from empty stack")
	}
	if val != 0 {
		t.Fatal("empty stack pop should have returned zero value")
	}
}
