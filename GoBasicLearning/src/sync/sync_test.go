package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var Wait sync.WaitGroup
var Counter int = 0

func Routine(id int) {
	for i := 0; i < 2; i++ {
		Counter++
		time.Sleep(time.Second)
	}
	Wait.Done()
}

func TestSync(t *testing.T) {
	for i := 1; i <= 2; i++ {
		Wait.Add(1)
		go Routine(i)
	}
	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

func BenchmarkWrongAddFunction(b *testing.B) {
	b.ResetTimer()
	wg := sync.WaitGroup{}
	result := 0
	for i := 0; i < 10; i++ {
		go func(delta int) {
			wg.Add(1)
			defer wg.Done()
			result += delta
		}(i)
	}
	wg.Wait()
	fmt.Println(result)
	b.StopTimer()
}

// 其实这样使用wg.Add(1)也并不完全对(抛出问题, 引发思考)
func BenchmarkRightAddFunction(b *testing.B) {
	b.ResetTimer()
	wg := sync.WaitGroup{}
	result := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(delta int) {
			defer wg.Done()
			result += delta
		}(i)
	}
	wg.Wait()
	fmt.Println(result)
	b.StopTimer()
}
