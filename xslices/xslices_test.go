package xslices

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type testCase struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      []int
		wantRaw   []int
	}
	tests := []testCase{
		{
			name:      "Filter nil slice",
			slice:     nil,
			predicate: func(i int) bool { return true },
			want:      nil,
			wantRaw:   nil,
		},
		{
			name:      "Filter empty slice",
			slice:     nil,
			predicate: func(i int) bool { return true },
			want:      nil,
			wantRaw:   nil,
		},
		{
			name:      "Filter to empty slice",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return false },
			want:      nil,
			wantRaw:   []int{1, 2, 3, 4, 5},
		},
		{
			name:      "base odd",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i%2 == 1 },
			want:      []int{1, 3, 5},
			wantRaw:   []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw := tt.slice
			got := Filter(raw, tt.predicate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(raw, tt.wantRaw) {
				t.Errorf("raw = %v, wantRaw %v", raw, tt.wantRaw)
			}
		})
	}
}

func TestFilterInplace(t *testing.T) {
	type testCase struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      []int
		wantRaw   []int
	}
	tests := []testCase{
		{
			name:      "Filter nil slice",
			slice:     nil,
			predicate: func(i int) bool { return true },
			want:      nil,
			wantRaw:   nil,
		},
		{
			name:      "Filter empty slice",
			slice:     nil,
			predicate: func(i int) bool { return true },
			want:      nil,
			wantRaw:   nil,
		},
		{
			name:      "Filter to empty slice",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return false },
			want:      nil,
			wantRaw:   []int{0, 0, 0, 0, 0},
		},
		{
			name:      "base odd",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i%2 == 1 },
			want:      []int{1, 3, 5},
			wantRaw:   []int{1, 3, 5, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw := tt.slice
			got := FilterInplace(raw, tt.predicate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterInplace() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(raw, tt.wantRaw) {
				t.Errorf("raw = %v, wantRaw %v", raw, tt.wantRaw)
			}
		})
	}
}
