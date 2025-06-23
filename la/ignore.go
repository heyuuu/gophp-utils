package la

func Ignore(err error) {
}

func Ignore1[T1 any](v1 T1, err error) T1 {
	return v1
}

func Ignore2[T1 any, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	return v1, v2
}

func Ignore3[T1 any, T2 any, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	return v1, v2, v3
}
