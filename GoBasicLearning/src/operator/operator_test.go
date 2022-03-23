package operator

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

// TODO Go当中数组比较前提是<数组元素类型且长度相同>
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Logf("a == b: %v\n", a == b)
	t.Logf("a == d: %v\n", a == d)
	t.Logf("array a's type: %T, lengh: %d\n", a, len(a))
	t.Logf("array c's type: %T, lengh: %d\n", c, len(c))
}

// 通过按位置零运算符来实现权限控制
func TestBitClear(t *testing.T) {
	a := 7 // 0111
	a = a &^ Readable
	a = a &^ Executable
	t.Logf("var a's value(Binary): %b\n", a)
	t.Logf("var a's Readable: %v\n", a&Readable == Readable)
	t.Logf("var a's Writable: %v\n", a&Writable == Writable)
	t.Logf("var a's Executable: %v\n", a&Executable == Executable)
}
