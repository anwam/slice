package slice

func Map[T comparable](items []T, cb func(t T) T) (result []T) {
	if len(items) == 0 {
		return items
	}
	for _, v := range items {
		result = append(result, cb(v))
	}
	return result
}

func Reduce[T, K any](s []T, fn func(K, T) K, initial K) (acc K) {
	acc = initial
	for _, v := range s {
		acc = fn(acc, v)
	}
	return acc
}

func Filter[T any](s []T, fn func(T) T) []T {
	var result []T
	for _, v := range s {
		result = append(result, fn(v))
	}
	return result
}
