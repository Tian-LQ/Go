package csp_concurrent

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working one something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

// fmt.Println方法需要等待service方法返回之后才能打印结果，再去执行otherTask方法
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
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
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
}
