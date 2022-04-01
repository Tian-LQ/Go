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
