package xstrings

import (
	"reflect"
	"testing"
)

func TestToUpper(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"", "", ""},
		{"", "Simple WORD", "SIMPLE WORD"},
		{"", "_wORD", "_WORD"},
		{"", "用户のiD", "用户のID"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUpper(tt.arg); got != tt.want {
				t.Errorf("ToUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToLower(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"", "", ""},
		{"", "Simple WORD", "simple word"},
		{"", "_wORD", "_word"},
		{"", "用户のiD", "用户のid"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLower(tt.arg); got != tt.want {
				t.Errorf("ToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"", "", ""},
		{"", "Simple WORD", "Simple word"},
		{"", "word", "Word"},
		{"", "WORD", "Word"},
		{"", "_wORD", "_word"},
		{"", "用户のiD", "用户のid"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Capitalize(tt.arg); got != tt.want {
				t.Errorf("Capitalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperFirst(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"", "", ""},
		{"", "Simple WORD", "Simple WORD"},
		{"", "word", "Word"},
		{"", "WORD", "WORD"},
		{"", "_wORD", "_wORD"},
		{"", "用户のiD", "用户のiD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpperFirst(tt.arg); got != tt.want {
				t.Errorf("UpperFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerFirst(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"", "", ""},
		{"", "Simple WORD", "simple WORD"},
		{"", "word", "word"},
		{"", "WORD", "wORD"},
		{"", "_wORD", "_wORD"},
		{"", "用户のiD", "用户のiD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerFirst(tt.arg); got != tt.want {
				t.Errorf("LowerFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareFold(t *testing.T) {
	tests := []struct {
		name string
		arg1 string
		arg2 string
		want int
	}{
		{"", "Word", "woRD", 0},
		{"", "BBB", "BBB0", -1},
		{"", "BBB", "BBA0", 1},
		{"", "中文id", "中文ID", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareFold(tt.arg1, tt.arg2); got != tt.want {
				t.Errorf("CompareFold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualFold(t *testing.T) {
	tests := []struct {
		name string
		arg1 string
		arg2 string
		want bool
	}{
		{"", "Word", "woRD", true},
		{"", "BBB", "BBB0", false},
		{"", "BBB", "BBA0", false},
		{"", "中文id", "中文ID", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EqualFold(tt.arg1, tt.arg2); got != tt.want {
				t.Errorf("EqualFold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasPrefixFold(t *testing.T) {
	tests := []struct {
		name string
		arg1 string
		arg2 string
		want bool
	}{
		{"", "Word", "woRD", true},
		{"", "Word000", "woRD", true},
		{"", "BBB", "BBB0", false},
		{"", "中文idxxx", "中文ID", true},
		{"", "xxx中文id", "中文ID", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPrefixFold(tt.arg1, tt.arg2); got != tt.want {
				t.Errorf("HasPrefixFold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasSuffixFold(t *testing.T) {
	tests := []struct {
		name string
		arg1 string
		arg2 string
		want bool
	}{
		{"", "Word", "woRD", true},
		{"", "Word000", "woRD", false},
		{"", "000Word", "woRD", true},
		{"", "BBB", "0BBB", false},
		{"", "中文idxxx", "中文ID", false},
		{"", "xxx中文id", "中文ID", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasSuffixFold(tt.arg1, tt.arg2); got != tt.want {
				t.Errorf("HasSuffixFold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitWords(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want []string
	}{
		{"", "", nil},
		{"", " word ", []string{"word"}},
		{"", "word", []string{"word"}},
		{"", "Simple word", []string{"Simple", "word"}},
		{"", "Simple-word", []string{"Simple", "word"}},
		{"", "Simple_word", []string{"Simple", "word"}},
		{"", "with01number", []string{"with", "01", "number"}},
		{"", "HTTPServer", []string{"HTTP", "Server"}},
		{"", "用户のId", []string{"用户の", "Id"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitWords(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllCases(t *testing.T) {
	type want struct {
		CamelCase          string
		PascalCase         string
		SnakeCase          string
		ScreamingSnakeCase string
		KebabCase          string
		ScreamingKebabCase string
	}

	tests := []struct {
		name string
		arg  string
		want want
	}{
		{"", "", want{}},
		{"", " ", want{}},
		{"", "Simple word", want{
			CamelCase:          "simpleWord",
			PascalCase:         "SimpleWord",
			SnakeCase:          "simple_word",
			ScreamingSnakeCase: "SIMPLE_WORD",
			KebabCase:          "simple-word",
			ScreamingKebabCase: "SIMPLE-WORD",
		}},
		{"", "wordWith01number", want{
			CamelCase:          "wordWith01Number",
			PascalCase:         "WordWith01Number",
			SnakeCase:          "word_with_01_number",
			ScreamingSnakeCase: "WORD_WITH_01_NUMBER",
			KebabCase:          "word-with-01-number",
			ScreamingKebabCase: "WORD-WITH-01-NUMBER",
		}},
		{"", "HTTPServer", want{
			CamelCase:          "httpServer",
			PascalCase:         "HttpServer",
			SnakeCase:          "http_server",
			ScreamingSnakeCase: "HTTP_SERVER",
			KebabCase:          "http-server",
			ScreamingKebabCase: "HTTP-SERVER",
		}},
		{"", "用户のid", want{
			CamelCase:          "用户のId",
			PascalCase:         "用户のId",
			SnakeCase:          "用户の_id",
			ScreamingSnakeCase: "用户の_ID",
			KebabCase:          "用户の-id",
			ScreamingKebabCase: "用户の-ID",
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// CamelCase
			if got := CamelCase(tt.arg); got != tt.want.CamelCase {
				t.Errorf("CamelCase() = %v, want %v", got, tt.want.CamelCase)
			}
			// PascalCase
			if got := PascalCase(tt.arg); got != tt.want.PascalCase {
				t.Errorf("PascalCase() = %v, want %v", got, tt.want.PascalCase)
			}
			// SnakeCase
			if got := SnakeCase(tt.arg); got != tt.want.SnakeCase {
				t.Errorf("SnakeCase() = %v, want %v", got, tt.want.SnakeCase)
			}
			// SnakeCase
			if got := ScreamingSnakeCase(tt.arg); got != tt.want.ScreamingSnakeCase {
				t.Errorf("ScreamingSnakeCase() = %v, want %v", got, tt.want.ScreamingSnakeCase)
			}
			// KebabCase
			if got := KebabCase(tt.arg); got != tt.want.KebabCase {
				t.Errorf("KebabCase() = %v, want %v", got, tt.want.KebabCase)
			}
			// ScreamingKebabCase
			if got := ScreamingKebabCase(tt.arg); got != tt.want.ScreamingKebabCase {
				t.Errorf("ScreamingKebabCase() = %v, want %v", got, tt.want.ScreamingKebabCase)
			}
		})
	}
}
