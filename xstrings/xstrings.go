package xstrings

import (
	"github.com/heyuuu/gophp-utils/ascii"
	"strings"
)

func IsBlank(s string) bool {
	for _, r := range s {
		if !ascii.IsSpace(r) {
			return false
		}
	}
	return true
}

func PadLeft(s string, size int, pad byte) string {
	if len(s) >= size {
		return s
	}

	buf := make([]byte, size)
	padSize := size - len(s)
	copy(buf[:padSize], s)
	for i := 0; i < padSize; i++ {
		buf[i] = pad
	}
	return string(buf)
}

func PadRight(s string, size int, pad byte) string {
	if len(s) >= size {
		return s
	}

	buf := make([]byte, size)
	copy(buf, s)
	for i := len(s); i < size; i++ {
		buf[i] = pad
	}
	return string(buf)
}

// LastCut
// 类似 strings.Cut()，但是是从字符串尾部反向开始查找的
func LastCut(s string, sep string) (before, after string, found bool) {
	if i := strings.LastIndex(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}
