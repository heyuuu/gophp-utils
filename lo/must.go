package lo

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Must1[T any](value1 T, err error) T {
	if err != nil {
		panic(err)
	}
	return value1
}

func Must2[T1 any, T2 any](value1 T1, value2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return value1, value2
}
