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
