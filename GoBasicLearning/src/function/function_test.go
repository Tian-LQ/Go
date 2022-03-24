package function

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ReturnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(100)
}

func Square(n int) int {
	time.Sleep(time.Second * 1)
	return n * n
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		res := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return res
	}
}

func TestFunc(t *testing.T) {
	a, b := ReturnMultiValues()
	t.Logf("a: %d, b: %d\n", a, b)
	fn := timeSpent(Square)
	t.Log(fn(10))
}

func Sum(ops ...int) int {
	result := 0
	for _, op := range ops {
		result += op
	}
	return result
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	//panic("error")
}
