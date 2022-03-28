package gods

import (
	"fmt"
	"testing"
)

func checkList[T comparable](ls *List[T], ch []T) error {
	if len(ch) == 0 {
		if ls.Front() == nil && ls.Back() == nil {
			return nil
		}
		return fmt.Errorf("expected empty list, found list with value")
	}
	node := ls.Front()
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

	node = ls.Back()
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

	t.Run("empty", func(t *testing.T) {
		check := []int{}

		if err := checkList(list, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("empty insert front", func(t *testing.T) {
		list.InsertFront(2)
		check := []int{2}

		if err := checkList(list, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("remove head to empty", func(t *testing.T) {
		if err := list.Remove(list.Front()); err != nil {
			t.Fatal(err)
		}
		check := []int{}
		if err := checkList(list, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("empty insert back", func(t *testing.T) {
		list.InsertBack(2)
		check := []int{2}

		if err := checkList(list, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("remove tail to empty", func(t *testing.T) {
		if err := list.Remove(list.Back()); err != nil {
			t.Fatal(err)
		}
		check := []int{}
		if err := checkList(list, check); err != nil {
			t.Fatal(err)
		}
	})

	list.InsertBefore(nil, 3)
	t.Run("populated insert front", func(t *testing.T) {
		list.InsertFront(2)
		check := []int{2, 3}

		if err := checkList(list, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("populated insert back", func(t *testing.T) {
		list.InsertBack(4)
		check := []int{2, 3, 4}

		if err := checkList(list, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("insert before head", func(t *testing.T) {
		list.InsertBefore(list.Front(), 1)
		check := []int{1, 2, 3, 4}

		if err := checkList(list, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("insert after tail", func(t *testing.T) {
		list.InsertAfter(list.Back(), 5)
		check := []int{1, 2, 3, 4, 5}

		if err := checkList(list, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("remove one after head", func(t *testing.T) {
		two := list.Front().Next()
		if err := list.Remove(two); err != nil {
			t.Fatal(err)
		}
		check := []int{1, 3, 4, 5}

		if err := checkList(list, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("remove one before tail", func(t *testing.T) {
		four := list.Back().Prev()
		if err := list.Remove(four); err != nil {
			t.Fatal(err)
		}
		check := []int{1, 3, 5}

		if err := checkList(list, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("insert after head", func(t *testing.T) {
		list.InsertAfter(list.Front(), 2)
		check := []int{1, 2, 3, 5}

		if err := checkList(list, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("insert before tail", func(t *testing.T) {
		list.InsertBefore(list.Back(), 4)
		check := []int{1, 2, 3, 4, 5}

		if err := checkList(list, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("remove head", func(t *testing.T) {
		if err := list.Remove(list.Front()); err != nil {
			t.Fatal(err)
		}
		check := []int{2, 3, 4, 5}

		if err := checkList(list, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("remove tail", func(t *testing.T) {
		if err := list.Remove(list.Back()); err != nil {
			t.Fatal(err)
		}
		check := []int{2, 3, 4}

		if err := checkList(list, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("insert head after remove", func(t *testing.T) {
		list.InsertBefore(list.Front(), 1)
		check := []int{1, 2, 3, 4}

		if err := checkList(list, check); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("insert tail after remove", func(t *testing.T) {
		list.InsertAfter(list.Back(), 5)
		check := []int{1, 2, 3, 4, 5}

		if err := checkList(list, check); err != nil {
			t.Fatal(err)
		}
	})

	for i := 0; i < 5; i++ {
		err := list.Remove(list.Front())
		if err != nil {
			t.Fatalf("list empty on iter %d", i)
		}
	}

	t.Run("remove from empty", func(t *testing.T) {
		err := list.Remove(list.Back())
		if err == nil {
			t.Fatal("empty list should have errored")
		}
		check := []int{}

		if err := checkList(list, check); err != nil {
			t.Fatal(err)
		}
	})

}
