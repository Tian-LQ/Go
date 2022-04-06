package sync

import (
	"fmt"
	"testing"
)

// TODO 阻塞式Goroutine任务池
type BlockTaskPool struct {
	// 用一个channel来控制数量
	concurrent chan struct{}
}

func NewBlockTaskPool(maxCnt int) *BlockTaskPool {
	return &BlockTaskPool{
		concurrent: make(chan struct{}, maxCnt),
	}
}

func (tp *BlockTaskPool) Do(action func()) {
	tp.concurrent <- struct{}{}
	go func() {
		action()
		<-tp.concurrent
	}()
}

// TODO 带缓存的Goroutine任务池
type CacheBlockTaskPool struct {
	// 用一个channel来控制数量
	concurrent chan struct{}
	// 用一个channel来缓存任务task
	queue chan func()
}

// 通常queueSize大于maxCnt
func NewCacheBlockTaskPool(maxCnt int, queueSize int) *CacheBlockTaskPool {
	return &CacheBlockTaskPool{
		concurrent: make(chan struct{}, maxCnt),
		queue:      make(chan func(), queueSize),
	}
}

// 当存在maxCnt个task正在工作(未退出)，且存在queueSize个task缓存时，该方法会阻塞
func (tp *CacheBlockTaskPool) Do(action func()) {
	tp.queue <- action
}

func (tp *CacheBlockTaskPool) Start() {
	go func() {
		// 最多会有maxCnt个TaskGoroutine启动且阻塞
		for {
			tp.concurrent <- struct{}{}
			go func() {
				tk := <-tp.queue
				tk()
				<-tp.concurrent
			}()
		}
	}()
}

func TestCacheBlockTaskPool(t *testing.T) {
	tp := NewCacheBlockTaskPool(5, 10)
	tp.Start()
	tp.Do(func() {
		fmt.Println("do some calculate task")
	})
}
