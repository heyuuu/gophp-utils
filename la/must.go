package la

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Must1[T1 any](v1 T1, err error) T1 {
	Must(err)
	return v1
}

func Must2[T1 any, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	Must(err)
	return v1, v2
}

func Must3[T1 any, T2 any, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	Must(err)
	return v1, v2, v3
}
