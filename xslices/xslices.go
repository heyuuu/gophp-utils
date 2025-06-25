package xslices

// Filter 过滤符合条件的元素生成新的slice
// 与 slices.DeleteFunc() 不同，xslices.Filter 会产生新的slices而不是原地修改
func Filter[S ~[]E, E any](s S, predicate func(E) bool) S {
	if len(s) == 0 {
		return nil
	}

	result := make(S, 0, len(s))
	for _, item := range s {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func Map[S ~[]E1, E1 any, E2 any](s S, transform func(E1) E2) []E2 {
	if len(s) == 0 {
		return nil
	}

	result := make([]E2, len(s))
	for i, v := range s {
		result[i] = transform(v)
	}
	return result
}

// LastIndex 类似 slices.Index()，但是是反向开始查找的
func LastIndex[S ~[]E, E comparable](s S, v E) int {
	for i := len(s) - 1; i >= 0; i-- {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// LastIndexFunc 类似 slices.IndexFunc()，但是是反向开始查找的
func LastIndexFunc[S ~[]E, E any](s S, f func(E) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return i
		}
	}
	return -1
}

// Every 判断切片中是否所有元素都符合条件
func Every[S ~[]E, E any](s S, predicate func(E) bool) bool {
	for _, v := range s {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// Any 判断切片中是否有任一元素符合条件
func Any[S ~[]E, E any](s S, predicate func(E) bool) bool {
	for _, v := range s {
		if predicate(v) {
			return true
		}
	}
	return false
}
