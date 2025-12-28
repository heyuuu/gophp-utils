package xstrings

import (
	"strconv"
	"testing"
)

func TestTrimIndent(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			arg:  "\tline1\n" + "\t\tline2\n",
			want: "line1\n" + "\tline2",
		},
		{
			arg: `
				line1
					line2
						line3
			`,
			want: "line1\n\tline2\n\t\tline3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimIndent(tt.arg); got != tt.want {
				t.Errorf("TrimIndent() = %v, want %v", strconv.Quote(got), strconv.Quote(tt.want))
			}
		})
	}
}
