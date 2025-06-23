package xstrings

import (
	"github.com/heyuuu/gophp-utils/ascii"
	"strings"
)

func indentWidth(s string) int {
	for i, c := range []byte(s) {
		if !ascii.IsSpace(c) {
			return i
		}
	}
	return len(s)
}

// TrimIndent 去除多行字符串的前置缩进，多用于代码中大段文本的美化表示
// 去除缩进长度为所有非空白行缩进长度的最小值；空白行缩进长度不够时置空
func TrimIndent(s string) string {
	lines := strings.Split(s, "\n")

	// 计算共同缩进长度
	commonIndent := -1
	for _, line := range lines {
		// 跳过空白行
		if IsBlank(line) {
			continue
		}

		// 非空白行计算共同缩进长度
		indent := indentWidth(line)
		if commonIndent < 0 {
			commonIndent = indent
		} else {
			commonIndent = min(commonIndent, indent)
		}
	}

	// 无需修改缩进时，直接返回原值
	if commonIndent <= 0 {
		return s
	}

	// 逐行修改
	for i, line := range lines {
		if commonIndent < len(line) {
			lines[i] = line[commonIndent:]
		} else {
			lines[i] = ""
		}
	}
	return strings.Join(lines, "\n")
}
