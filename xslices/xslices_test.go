package xslices

import (
	"reflect"
	"strconv"
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

func TestMap(t *testing.T) {
	type testCase struct {
		name      string
		slice     []int
		transform func(int) string
		want      []string
	}
	tests := []testCase{
		{
			name:      "Map nil slice",
			slice:     nil,
			transform: strconv.Itoa,
			want:      nil,
		},
		{
			name:      "Map empty slice",
			slice:     []int{},
			transform: strconv.Itoa,
			want:      nil,
		},
		{
			name:      "Map simple slice",
			slice:     []int{1, 2, 3},
			transform: strconv.Itoa,
			want:      []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.slice, tt.transform); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapInplace(t *testing.T) {
	type testCase struct {
		name      string
		slice     []int
		transform func(int) int
		wantRaw   []int
	}
	tests := []testCase{
		{
			name:      "Map nil slice",
			slice:     nil,
			transform: func(i int) int { return i * 2 },
			wantRaw:   nil,
		},
		{
			name:      "Map empty slice",
			slice:     []int{},
			transform: func(i int) int { return i * 2 },
			wantRaw:   []int{},
		},
		{
			name:      "Map simple slice",
			slice:     []int{1, 2, 3},
			transform: func(i int) int { return i * 2 },
			wantRaw:   []int{2, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw := tt.slice
			MapInplace(raw, tt.transform)
			if !reflect.DeepEqual(raw, tt.wantRaw) {
				t.Errorf("raw = %v, wantRaw %v", raw, tt.wantRaw)
			}
		})
	}
}
