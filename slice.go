package slice

func Map[T, K any](s []T, fn func(T) K) (result []K) {
	if len(s) == 0 {
		return result
	}
	for _, v := range s {
		result = append(result, fn(v))
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

func Filter[T any](s []T, fn func(T) bool) []T {
	var result []T
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
