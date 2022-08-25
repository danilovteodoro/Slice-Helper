package collections

func MapArray[T, R any](items []T, f func(t T) R) (out []R) {
	for _, item := range items {
		out = append(out, f(item))
	}
	return
}

func FilterArray[T any](items []T, f func(t T) bool) (out []T) {
	for _, item := range items {
		if f(item) {
			out = append(out, item)
		}
	}
	return out
}

func Some[T any](items []T, f func(t T) bool) bool {
	for _, item := range items {
		if f(item) {
			return true
		}
	}
	return false
}

func Every[T any](items []T, f func(t T) bool) bool {
	for _, item := range items {
		if !f(item) {
			return false
		}
	}
	return true
}

func Find[T any](items []T, f func(t T) bool) *T {
	for _, item := range items {
		if f(item) {
			return &item
		}
	}
	return nil
}

func Group[T, R any](items []T, f func(t T, r R) R) (out R) {

	for _, item := range items {
		out = f(item, out)
	}

	return
}
