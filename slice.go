package slice

func Map[T comparable](items []T, cb func(t T) T) []T {
	if len(items) == 0 {
		return items
	}
	var result []T
	for _, v := range items {
		result = append(result, cb(v))
	}
	return result
}
