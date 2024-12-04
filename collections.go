package collections

func Any[T any](iterable []T, pred func(T) bool) (bool, int) {
	for idx, value := range iterable {
		if pred(value) {
			return true, idx
		}
	}
	return false, -1
}

func All[T any](iterable []T, pred func(T) bool) (bool, int) {
	b, idx := Any(iterable, func(arg T) bool { return !pred(arg) })
	return !b, idx
}

func Count[T any](iterable []T, pred func(T) bool) int {
	count := 0
	for _, v := range iterable {
		if pred(v) {
			count++
		}
	}
	return count
}
