package ascii

// ascii 相关函数
// 提供功能类似 c 标准库 ctype.h

const (
	MaxAscii = 0x7f
)

func IsAscii[T byte | rune](c T) bool {
	return 0 <= c && c <= MaxAscii
}

func IsLower[T byte | rune](c T) bool {
	return 'a' <= c && c <= 'z'
}

func IsUpper[T byte | rune](c T) bool {
	return 'A' <= c && c <= 'Z'
}

func IsAlpha[T byte | rune](c T) bool {
	return IsLower(c) || IsUpper(c)
}

func IsDigit[T byte | rune](c T) bool {
	return '0' <= c && c <= '9'
}

func IsAlphaNum[T byte | rune](c T) bool {
	return IsAlpha(c) || IsDigit(c)
}

func IsXDigit[T byte | rune](c T) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

func ParseXDigit[T byte | rune](c T) (byte, bool) {
	if c >= '0' && c <= '9' {
		return byte(c - '0'), true
	} else if c >= 'A' && c <= 'F' {
		return byte(c - 'A' + 10), true
	} else if c >= 'a' && c <= 'f' {
		return byte(c - 'a' + 10), true
	} else {
		return 0, false
	}
}

func IsControl[T byte | rune](c T) bool {
	return (0 <= c && c <= 0x1f) || c == 0x7f
}

func IsSpace[T byte | rune](c T) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\v' || c == '\f' || c == '\r'
}

func ToLower[T byte | rune](c T) T {
	if IsUpper(c) {
		return c - 'A' + 'a'
	}
	return c
}

func ToUpper[T byte | rune](c T) T {
	if IsLower(c) {
		return c - 'a' + 'A'
	}
	return c
}
