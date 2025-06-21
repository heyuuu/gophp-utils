package mapx

import (
	"maps"
	"slices"
	"strconv"
	"testing"
)

func stdKeys[M ~map[K]V, K comparable, V any](m M) []K {
	return slices.Collect(maps.Keys(m))
}
func stdValues[M ~map[K]V, K comparable, V any](m M) []V {
	return slices.Collect(maps.Values(m))
}

func makeStubMap(size int) map[string]string {
	m := make(map[string]string, size)
	for i := 0; i < size; i++ {
		key := "key-" + strconv.Itoa(i)
		value := "value-" + strconv.Itoa(i)
		m[key] = value
	}
	return m
}

func Benchmark_Keys(b *testing.B) {
	benchmarks := []struct {
		name string
		m    map[string]string
	}{
		{
			name: "size 100",
			m:    makeStubMap(100),
		},
		{
			name: "size 10000",
			m:    makeStubMap(10000),
		},
		{
			name: "size 1000000",
			m:    makeStubMap(1000000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Keys(bm.m)
			}
		})
		b.Run(bm.name+"_std", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				stdKeys(bm.m)
			}
		})
	}
}

func Benchmark_Values(b *testing.B) {
	benchmarks := []struct {
		name string
		m    map[string]string
	}{
		{
			name: "size 100",
			m:    makeStubMap(100),
		},
		{
			name: "size 10000",
			m:    makeStubMap(10000),
		},
		{
			name: "size 1000000",
			m:    makeStubMap(1000000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Values(bm.m)
			}
		})
		b.Run(bm.name+"_std", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				stdValues(bm.m)
			}
		})
	}
}
