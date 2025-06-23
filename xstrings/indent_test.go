package xstrings

import (
	"testing"
)

func TestTrimIndent(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			"",
			"\tline1\n" + "\t\tline2\n",
			"line1\n" + "\tline2\n",
		},
		{
			"",
			`
			line1
				line2
					line3
`,
			`
line1
	line2
		line3
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimIndent(tt.arg); got != tt.want {
				t.Errorf("TrimIndent() = %v, want %v", got, tt.want)
			}
		})
	}
}
