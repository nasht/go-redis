package array

func Pop[T any](a []T) (T, []T) {
	if len(a) == 0 {
		var empty T
		return empty, a
	}
	return a[len(a)-1], a[:len(a)-1]
}

func Shift[T any](a []T) (T, []T) {
	if len(a) == 0 {
		var empty T
		return empty, a
	}
	return a[0], a[1:]
}