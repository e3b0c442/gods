package gods

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

func Reverse[T any](sl []T) {
	if len(sl) == 0 {
		return
	}
	for i := 0; i < len(sl)/2; i++ {
		sl[i], sl[len(sl)-1-i] = sl[len(sl)-1-i], sl[i]
	}
}
