package mapx

// Keys returns the keys of the map m.
// The keys will be in an indeterminate order.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k, _ := range m {
		r = append(r, k)
	}
	return r
}

// Values returns the values of the map m.
// The values will be in an indeterminate order.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
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
