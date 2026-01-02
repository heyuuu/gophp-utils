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

// Join
// strings.Join() 的别名占位
func Join(elems []string, sep string) string {
	return strings.Join(elems, sep)
}

// JoinFunc
// 类似 strings.Join()，但是支持非 string 类型列表+转换函数
func JoinFunc[T any](elems []T, sep string, transform func(T) string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return transform(elems[0])
	}

	var buf strings.Builder
	for i, elem := range elems {
		if i > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(transform(elem))
	}
	return buf.String()
}

func ReverseJoin(elems []string, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}

	size := (len(elems) - 1) * len(sep)
	for _, elem := range elems {
		size += len(elem)
	}

	var buf strings.Builder
	buf.Grow(size)
	for i := len(elems) - 1; i >= 0; i-- {
		buf.WriteString(elems[i])
		if i > 0 {
			buf.WriteString(sep)
		}
	}
	return buf.String()
}
