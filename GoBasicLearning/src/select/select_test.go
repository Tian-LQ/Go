package _select

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func AsyncService() chan string {
	//retCh := make(chan string)
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("return result.")
		// 此时通过改用channel buffer的方式，解决了channel写入后阻塞的问题
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	select {
	case retCh := <-AsyncService():
		t.Log(retCh)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}

func TestSelectAndFor(t *testing.T) {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	close(intChan)
	syncChan := make(chan struct{}, 1)
	// 这里之所以用标签Loop是因为break在select中只能跳出当前select语句块
	go func() {
	Loop:
		for {
			select {
			case e, ok := <-intChan:
				if !ok {
					fmt.Println("End.")
					break Loop
				}
				fmt.Printf("Received: %v\n", e)
			}
		}
		syncChan <- struct{}{}
	}()
	<-syncChan
}
