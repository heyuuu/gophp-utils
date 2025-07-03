package xmaps

import (
	"cmp"
	"slices"
)

// Keys 以 slice 形式返回所有键，不保证结果有序
// 功能等价于 slices.Collect(maps.Keys(m))，但不用 iter 性能更好
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k, _ := range m {
		r = append(r, k)
	}
	return r
}

// Values 以 slice 形式返回所有值，不保证结果有序
// 功能等价于 slices.Collect(maps.Values(m))，但不用 iter 性能更好
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// SortedKeys 以 slice 形式返回所有键，结果按默认规则排序
func SortedKeys[M ~map[K]V, K cmp.Ordered, V any](m M) []K {
	keys := Keys(m)
	slices.Sort(keys)
	return keys
}

// Combine 创建一个Map，以一个数组作为其键名，另一个数组作为其值
func Combine[Keys ~[]K, Values ~[]V, K comparable, V any](keys Keys, values Values) map[K]V {
	if len(keys) != len(values) {
		panic("slices.Combine: keys and values must have the same length")
	}

	result := make(map[K]V, len(keys))
	for i, key := range keys {
		result[key] = values[i]
	}
	return result
}

func GetOrElse[K comparable, V any, Map ~map[K]V](m Map, key K, defaultValue func() V) V {
	if value, exists := m[key]; exists {
		return value
	} else {
		return defaultValue()
	}
}

func GetOrPut[K comparable, V any, Map ~map[K]V](m Map, key K, defaultValue func() V) V {
	if value, exists := m[key]; exists {
		return value
	} else {
		value = defaultValue()
		if m != nil {
			m[key] = value
		}
		return value
	}
}
