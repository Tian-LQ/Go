package _interface

import (
	"fmt"
	"testing"
	"time"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() string {
	return fmt.Sprintf(`fmt.Println("Hello World")`)
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

// 自定义类型
type MyFunction func(op int) int

func timeSpent(inner MyFunction) MyFunction {
	return func(n int) int {
		start := time.Now()
		res := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return res
	}
}

func Square(n int) int {
	time.Sleep(time.Second * 1)
	return n * n
}

func TestFunc(t *testing.T) {
	fn := timeSpent(Square)
	t.Log(fn(10))
	t.Logf("Square's type: %T\n", Square)
}
