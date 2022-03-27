package concurrent

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second)
	t.Logf("counter: %d\n", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	mutex := sync.Mutex{}
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			// 在defer块中释放锁资源
			defer func() {
				mutex.Unlock()
			}()
			mutex.Lock()
			counter++
		}()
	}
	// 此行代码是为了等待所有协程执行完毕
	time.Sleep(time.Second)
	t.Logf("counter: %d\n", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			// 在defer块中释放锁资源
			defer func() {
				mutex.Unlock()
			}()
			mutex.Lock()
			counter++
			wg.Done()
		}()
	}
	// 通过WaitGroup的方式来等待所有协程执行完毕
	wg.Wait()
	t.Logf("counter: %d\n", counter)
}
