package la

import "testing"

func TestXor(t *testing.T) {
	tests := []struct {
		name string
		a    bool
		b    bool
		want bool
	}{
		{"", false, false, false},
		{"", false, true, true},
		{"", true, false, true},
		{"", true, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Xor(tt.a, tt.b); got != tt.want {
				t.Errorf("Xor() = %v, want %v", got, tt.want)
			}
		})
	}
}
