package thread_groutine

import (
	"fmt"
	"testing"
)

func Delay() {
	fns := make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		fns = append(fns, func() {
			fmt.Printf("hello, this is : %d\n", i)
		})
		fns[i]()
	}

	fmt.Println("**************************************")

	for _, fn := range fns {
		fn()
	}
}

func TestGoroutine1(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func TestGoroutine2(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

func TestClosure(t *testing.T) {
	Delay()
}
