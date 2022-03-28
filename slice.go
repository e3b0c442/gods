package gods

//Equal returns true if the two slices are have the same contents. The slices must be the same type and the type must be
//comparable.
func Equal[T comparable](l, r []T) bool {
	if len(l) != len(r) {
		return false
	}
	for i := range l {
		if l[i] != r[i] {
			return false
		}
	}
	return true
}

//Reverse reverses a slice with O(n) complexity and O(1) memory.
func Reverse[T any](sl []T) {
	if len(sl) == 0 {
		return
	}
	for i := 0; i < len(sl)/2; i++ {
		sl[i], sl[len(sl)-1-i] = sl[len(sl)-1-i], sl[i]
	}
}
