package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 默认零值为空字符串""
	s = "hello"
	t.Log(len(s))
	t.Log(s[1])
	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))
}

func TestStringUTF8(t *testing.T) {
	s := "田"
	t.Logf("the length of s: %d\n", len(s)) // byte数
	c := []rune(s)
	t.Logf("the length of c: %d\n", len(c))
	t.Logf("田's unicode: %x\n", c[0])
	t.Logf("田's UTF8: %x\n", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	// string的forr遍历是对于rune而言的而非byte
	for _, c := range s {
		t.Logf("%[1]c %[1]d\n", c)
	}
	for i := 0; i < len(s); i++ {
		t.Logf("s[%d]: %d\n", i, s[i])
	}
}

func TestStringFn(t *testing.T) {
	str := "A,B,C"
	parts := strings.Split(str, ",")
	t.Log(parts)
	newStr := strings.Join(parts, "|")
	t.Log(newStr)
}

func TestStringConvert(t *testing.T) {
	str := strconv.Itoa(97)
	t.Logf("str: %v\n", str)
	strNumber := "99999"
	if result, err := strconv.Atoi(strNumber); err == nil {
		t.Log(result + 1)
	}
}
