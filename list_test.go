package gods

import (
	"fmt"
	"testing"
)

func checkList[T comparable](ls *List[T], ch []T) error {
	if len(ch) == 0 {
		if ls.Head() == nil && ls.Tail() == nil {
			return nil
		}
		return fmt.Errorf("expected empty list, found list with value")
	}
	node := ls.Head()
	for _, v := range ch {
		if node == nil {
			return fmt.Errorf("list ended unexpectedly, expected value %v", v)
		} else if node.Value != v {
			return fmt.Errorf("got %v expected %v", node.Value, v)
		}
		node = node.Next()
	}
	if node != nil {
		return fmt.Errorf("expected end of list, got %#v", node.Value)
	}

	node = ls.Tail()
	Reverse(ch)
	for _, v := range ch {
		if node == nil {
			return fmt.Errorf("reverse list ended unexpectedly, expected value %v", v)
		} else if node.Value != v {
			return fmt.Errorf("reverse got %v expected %v", node.Value, v)
		}
		node = node.prev
	}
	if node != nil {
		return fmt.Errorf("reverse expected end of list, got %#v", node.Value)
	}
	return nil
}
func TestList(t *testing.T) {
	list := NewList[int]()
	check := []int{}

	if err := checkList(list, check); err != nil {
		t.Fatal(err)
	}

	list.InsertAfter(nil, 2)
	check = []int{2}

	if err := checkList(list, check); err != nil {
		t.Fatal(err)
	}

	list.InsertBefore(list.Head(), 1)
	check = []int{1, 2}
	if err := checkList(list, check); err != nil {
		t.Fatal(err)
	}

	list.InsertBefore(nil, 3)
	check = []int{1, 2, 3}
	if err := checkList(list, check); err != nil {
		t.Fatal(err)
	}

	two := list.Head().Next()
	if err := list.Remove(two); err != nil {
		t.Fatal(err)
	}
	check = []int{1, 3}
	if err := checkList(list, check); err != nil {
		t.Fatal(err)
	}

	if err := list.Remove(list.Tail()); err != nil {
		t.Fatal(err)
	}
	check = []int{1}
	if err := checkList(list, check); err != nil {
		t.Fatal(err)
	}

	if err := list.Remove(list.Head()); err != nil {
		t.Fatal(err)
	}
	check = []int{}
	if err := checkList(list, check); err != nil {
		t.Fatal(err)
	}

	err := list.Remove(list.Tail())
	if err == nil {
		t.Fatal("empty list should have errored")
	}
	if err := checkList(list, check); err != nil {
		t.Fatal(err)
	}

}
