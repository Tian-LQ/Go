package data_type

import (
	"math"
	"testing"
)

type MyInt int64

func TestType(t *testing.T) {
	// a, b, c三个变量都不能进行相互转换
	var a int = 1
	var b int64 = 1
	var c MyInt = 1
	t.Logf("var a's type: %T\n", a)
	t.Logf("var b's type: %T\n", b)
	t.Logf("var c's type: %T\n", c)
	// 预定义值
	t.Logf("math.MaxInt64: %d\n", math.MaxInt64)
	t.Logf("math.MaxFloat64: %f\n", math.MaxFloat64)
	t.Logf("math.MaxUint32: %d\n", math.MaxUint32)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Logf("var a's type: %T\n", a)
	t.Logf("var aPtr's type: %T\n", aPtr)
	vec := []int{1, 2, 3}
	t.Logf("var vec's type: %T\n", vec)
	vecFirstNumberPtr := &vec[0]
	t.Logf("var vecFirstNumberPtr's type: %T\n", vecFirstNumberPtr)
}

func TestString(t *testing.T) {
	var str string
	t.Logf("var str's type: %T\n", str)
	t.Logf("var str's value: [%s]\n", str)
	t.Logf("var str's length: %d\n", len(str))
}
