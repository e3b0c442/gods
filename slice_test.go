package gods

import "testing"

func TestEqual(t *testing.T) {

	for _, test := range []struct {
		Label string
		L     []int
		R     []int
		Equal bool
	}{
		{
			Label: "empty",
			L:     []int{},
			R:     []int{},
			Equal: true,
		},
		{
			Label: "differing sizes",
			L:     []int{},
			R:     []int{1},
			Equal: false,
		},
		{
			Label: "same size different members",
			L:     []int{1, 2},
			R:     []int{2, 1},
			Equal: false,
		},
		{
			Label: "equal single",
			L:     []int{1},
			R:     []int{1},
			Equal: true,
		},
		{
			Label: "equal multiple",
			L:     []int{1, 2},
			R:     []int{1, 2},
			Equal: true,
		},
	} {
		t.Run(test.Label, func(t *testing.T) {
			if Equal(test.L, test.R) != test.Equal {
				t.Fatalf("L: %#v R: %#v", test.L, test.R)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	for _, test := range []struct {
		Label string
		In    []int
		Out   []int
	}{
		{
			Label: "empty",
			In:    []int{},
			Out:   []int{},
		},
		{
			Label: "one",
			In:    []int{1},
			Out:   []int{1},
		},
		{
			Label: "two",
			In:    []int{1, 2},
			Out:   []int{2, 1},
		},
		{
			Label: "three",
			In:    []int{1, 2, 3},
			Out:   []int{3, 2, 1},
		},
	} {
		t.Run(test.Label, func(t *testing.T) {
			orig := make([]int, len(test.In))
			copy(orig, test.In)

			Reverse(test.In)
			if !Equal(test.In, test.Out) {
				t.Fatalf("got %#v from orig %#v, expected %#v", test.In, orig, test.Out)
			}
		})
	}
}
