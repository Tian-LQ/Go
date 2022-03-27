package task_cancel

import (
	"fmt"
	"testing"
	"time"
)

// 当cancelChannel未关闭且为空时会return false
// 当cancelChannel当中存在结构体时，会return true并接收一个结构体
// 当cancelChannel关闭时会直接return true
func isCancelled(cancelChannel chan struct{}) bool {
	select {
	case <-cancelChannel:
		return true
	default:
		return false
	}
}

// 取消一个任务
func cancelFirst(cancelChannel chan struct{}) {
	cancelChannel <- struct{}{}
}

// 取消全部任务
func cancelSecond(cancelChannel chan struct{}) {
	close(cancelChannel)
}

func TestCancel(t *testing.T) {
	cancelChannel := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Printf("Number %d task Cancelled\n", i)
		}(i, cancelChannel)
	}
	//cancelFirst(cancelChannel)
	cancelSecond(cancelChannel)
	time.Sleep(time.Second * 1)
}
