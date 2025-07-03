package xslices

// Filter 过滤符合条件的元素生成新的slice
// 与 slices.FilterInplace() 实现不同，本函数会分配新内存而不会影响旧 slice
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

	if len(result) == 0 {
		return nil
	}
	return result
}

// FilterInplace 过滤符合条件的元素生成新的slice
// 与 slices.Filter() 实现不同，本函数会修改原 slice 而不会分配新内存
// 功能类似于 slices.DeleteFunc() 的，差别在 predicate 函数的结果是反向的
func FilterInplace[S ~[]E, E any](s S, predicate func(E) bool) S {
	if len(s) == 0 {
		return nil
	}

	result := s[0:0:len(s)] // 复用原 slice 内存
	for _, item := range s {
		if predicate(item) {
			result = append(result, item)
		}
	}
	clear(s[len(result):]) // 重置未使用的内存，与 slices.DeleteFunc() 中的逻辑类似

	if len(result) == 0 {
		return nil
	}
	return result
}

// Map 为切片的每个值应用回调函数并将结果作为新的切片
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

// MapInplace 为切片的每个值应用回调函数并设置到原位置
// 与 Map 不同，此函数会直接修改切片而不会申请新内存
func MapInplace[S ~[]E, E any](s S, transform func(E) E) {
	for i, v := range s {
		s[i] = transform(v)
	}
}

// Diff 计算切片的差集，不修改原切片
// 保留所有在 s 中存在且未在 others 里出现的元素，元素顺序不变
func Diff[S ~[]E, E comparable](s S, others ...S) S {
	if len(s) == 0 {
		return nil
	}

	keep := make(map[E]struct{}, len(s))
	for _, item := range s {
		keep[item] = struct{}{}
	}

	for _, other := range others {
		for _, item := range other {
			delete(keep, item)
			if len(keep) == 0 {
				return nil
			}
		}
		if len(keep) == 0 {
			return nil
		}
	}

	result := make(S, 0, len(s)) // 注意此处不可用 len(keep)，因为会有重复值
	for _, item := range s {
		if _, ok := keep[item]; ok {
			result = append(result, item)
		}
	}
	return result
}

func Unique[S ~[]E, E comparable](s S) S {
	result := make(S, 0, len(s))
	seen := make(map[E]struct{}, len(s))

	for _, item := range s {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			result = append(result, item)
		}
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
