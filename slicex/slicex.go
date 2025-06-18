package slicex

func Filter[T any, Slice ~[]T](s Slice, predicate func(T) bool) Slice {
	if len(s) == 0 {
		return nil
	}

	result := make(Slice, 0, len(s))
	for _, item := range s {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func Map[T any, Slice ~[]T, R any](s Slice, transform func(T) R) []R {
	if len(s) == 0 {
		return nil
	}

	result := make([]R, len(s))
	for i, v := range s {
		result[i] = transform(v)
	}
	return result
}
