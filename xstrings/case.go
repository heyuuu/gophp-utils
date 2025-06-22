package xstrings

// 本文件内是处理字符串大小写相关的函数。
// 这里有几个通用约定:
// - 大小写转换只处理 ASCII 范围内的字符，对之外的字符保持不变。与标准库函数(strings.ToLower/strings.ToUpper等)逻辑不同。
// - 分词时将非 ASCII 字符作为单独类型，不同 ASCII 间不做区分。e.g. "用户のID" 会分词为 "用户の" + "ID"

import (
	"unsafe"
)

var (
	byteIsLower [256]bool
	byteIsUpper [256]bool
	byteIsAlpha [256]bool
	byteToLower [256]byte
	byteToUpper [256]byte
)

func init() {
	// init byteIsLower / byteIsUpper / byteIsAlpha
	for c := byte('a'); c <= byte('z'); c++ {
		byteIsLower[c] = true
		byteIsAlpha[c] = true
	}
	for c := byte('A'); c <= byte('Z'); c++ {
		byteIsUpper[c] = true
		byteIsAlpha[c] = true
	}

	// init byteToLower / byteToUpper
	for i := 0; i <= 255; i++ {
		byteToLower[i] = byte(i)
		byteToUpper[i] = byte(i)
	}
	for c := byte('a'); c <= byte('z'); c++ {
		byteToUpper[c] = c - 'a' + 'A'
	}
	for c := byte('A'); c <= byte('Z'); c++ {
		byteToLower[c] = c - 'A' + 'a'
	}
}

// unsafeBytesToString 字符切片直接转string，要求调用方保证字符切片不再修改
func unsafeBytesToString(s []byte) string {
	return unsafe.String(unsafe.SliceData(s), len(s))
}

// ToUpper 字符串转小写
// 与标准库 strings.ToUpper() 的区别是，它不处理除英文字母外的其他unicode字母
func ToUpper(s string) string {
	var buf []byte
	for i, c := range []byte(s) {
		rc := byteToUpper[c]
		if rc != c {
			// 懒初始化，只在字符变更时才构建新字符串
			if buf == nil {
				buf = []byte(s)
			}
			buf[i] = rc
		}
	}
	if buf == nil {
		return s
	}
	return unsafeBytesToString(buf)
}

// ToLower 字符串转小写
// 与标准库 strings.ToLower() 的区别是，它不处理除英文字母外的其他unicode字母
func ToLower(s string) string {
	var buf []byte
	for i, c := range []byte(s) {
		rc := byteToLower[c]
		if rc != c {
			// 懒初始化，只在字符变更时才构建新字符串
			if buf == nil {
				buf = []byte(s)
			}
			buf[i] = rc
		}
	}
	if buf == nil {
		return s
	}
	return unsafeBytesToString(buf)
}

// Capitalize 首字母大写，其他字母小写
func Capitalize(s string) string {
	var buf []byte
	for i, c := range []byte(s) {
		var rc byte
		if i == 0 {
			rc = byteToUpper[c]
		} else {
			rc = byteToLower[c]
		}

		if rc != c {
			// 懒初始化，只在字符变更时才构建新字符串
			if buf == nil {
				buf = []byte(s)
			}
			buf[i] = rc
		}
	}
	if buf == nil {
		return s
	}
	return string(buf)
}

// UpperFirst 首字母大写
func UpperFirst(s string) string {
	if s == "" || !byteIsLower[s[0]] {
		return s
	}
	return string(append([]byte{s[0] - 'a' + 'A'}, s[1:]...))
}

// LowerFirst 首字母小写
func LowerFirst(s string) string {
	if s == "" || !byteIsUpper[s[0]] {
		return s
	}
	return string(append([]byte{s[0] - 'A' + 'a'}, s[1:]...))
}

// Compare 比较字符串，忽略大小写
// 与标准库 strings.Compare() 的区别是，它不处理除英文字母外的其他unicode字母
func CompareFold(s1 string, s2 string) int {
	l := min(len(s1), len(s2))
	for i := 0; i < l; i++ {
		c1, c2 := byteToLower[s1[i]], byteToLower[s2[i]]
		if c1 == c2 {
			continue
		} else if c1 < c2 {
			return -1
		} else {
			return 1
		}
	}
	if l < len(s1) {
		return 1
	}
	if l < len(s2) {
		return -1
	}
	return 0
}

// EqualFold 比较字符串是否相等，忽略大小写
// 与标准库 strings.EqualFold() 的区别是，它不处理除英文字母外的其他unicode字母
func EqualFold(s1 string, s2 string) bool {
	return len(s1) == len(s2) && CompareFold(s1, s2) == 0
}

// HasPrefixFold 判断字符串是否有指定前缀，忽略大小写
func HasPrefixFold(s string, prefix string) bool {
	return len(s) >= len(prefix) && CompareFold(s[:len(prefix)], prefix) == 0
}

// HasSuffixFold 判断字符串是否有指定后缀，忽略大小写
func HasSuffixFold(s string, suffix string) bool {
	return len(s) >= len(suffix) && CompareFold(s[len(s)-len(suffix):], suffix) == 0
}

const (
	stateSeparator = iota // start or ' ' or '-' or '_'
	stateLower
	stateUpper
	stateDigit
	stateOthers
)

var byteStates [256]int

func init() {
	for i := 0; i < len(byteStates); i++ {
		byteStates[i] = stateOthers
	}
	for i := 'A'; i <= 'Z'; i++ {
		byteStates[i] = stateUpper
	}
	for i := 'a'; i <= 'z'; i++ {
		byteStates[i] = stateLower
	}
	for i := '0'; i <= '9'; i++ {
		byteStates[i] = stateDigit
	}
	for _, c := range []byte(" -_") {
		byteStates[c] = stateSeparator
	}
}

// splitWords 分词，返回字符串切分后的单词列表
func splitWords(s string) []string {
	var result []string
	var state int = stateSeparator
	var wordStart int = 0
	for i, c := range []byte(s) {
		nextState := byteStates[c]
		if state == nextState {
			continue
		}

		// AAAaa 形式会拆分为 "AA" + "Aaa"
		if state == stateUpper && (nextState == stateLower) {
			if wordStart < i-1 {
				result = append(result, s[wordStart:i-1])
			}

			state = nextState
			wordStart = i - 1
		} else {
			if state != stateSeparator {
				result = append(result, s[wordStart:i])
			}

			state = nextState
			wordStart = i
		}

	}
	if state != stateSeparator {
		result = append(result, s[wordStart:])
	}
	return result
}

// commonCase 通用的字符串case处理函数
// @param s 原字符串
// @param sep 分隔符
// @param caseHandler 单词case处理函数
func commonCase(s string, sep string, caseHandler func(i int, word []byte)) string {
	// 分词
	words := splitWords(s)
	if len(words) == 0 {
		return ""
	}

	// 预计算结果字符串尺寸
	size := 0
	for i, word := range words {
		if i > 0 {
			size += len(sep)
		}
		size += len(word)
	}

	//
	buf := make([]byte, 0, size)
	for i, word := range words {
		// 添加分隔符
		if i > 0 {
			buf = append(buf, sep...)
		}

		// 非英文单词不处理大小写
		if !byteIsAlpha[word[0]] {
			buf = append(buf, word...)
			continue
		}

		// 英文单词处理大小写
		wordStart := len(buf)
		buf = append(buf, word...)
		caseHandler(i, buf[wordStart:])
	}
	return unsafeBytesToString(buf)
}

// CamelCase 驼峰命名法(又称小驼峰命名法)，e.g. "userName"
func CamelCase(s string) string {
	return commonCase(s, "", func(wordIndex int, word []byte) {
		for charIndex, c := range word {
			if charIndex == 0 && wordIndex > 0 {
				word[charIndex] = byteToUpper[c]
			} else {
				word[charIndex] = byteToLower[c]
			}
		}
	})
}

// PascalCase 帕斯卡命名法(又称大驼峰命名法)，e.g. "UserName"
func PascalCase(s string) string {
	return commonCase(s, "", func(wordIndex int, word []byte) {
		for charIndex, c := range word {
			if charIndex == 0 {
				word[charIndex] = byteToUpper[c]
			} else {
				word[charIndex] = byteToLower[c]
			}
		}
	})
}

// SnakeCase 蛇型命名法，e.g. "user_name"
func SnakeCase(s string) string {
	return commonCase(s, "_", func(wordIndex int, word []byte) {
		for charIndex, c := range word {
			word[charIndex] = byteToLower[c]
		}
	})
}

// ScreamingSnakeCase 大蛇型命名法，e.g. "USER_NAME"
func ScreamingSnakeCase(s string) string {
	return commonCase(s, "_", func(wordIndex int, word []byte) {
		for charIndex, c := range word {
			word[charIndex] = byteToUpper[c]
		}
	})
}

// KebabCase 烤串式命名法，e.g. "user-name"
func KebabCase(s string) string {
	return commonCase(s, "-", func(wordIndex int, word []byte) {
		for charIndex, c := range word {
			word[charIndex] = byteToLower[c]
		}
	})
}

// KebabCase 大烤串式命名法，e.g. "USER-NAME"
func ScreamingKebabCase(s string) string {
	return commonCase(s, "-", func(wordIndex int, word []byte) {
		for charIndex, c := range word {
			word[charIndex] = byteToUpper[c]
		}
	})
}
