package task

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var count int32 = 0

// getMessage 读请求方法
func getMessage(id int) (string, error) {
	// 调用该方法的次数越多，请求耗时就越多
	atomic.AddInt32(&count, 1)
	time.Sleep(time.Duration(count) * time.Millisecond)
	fmt.Println("call getMessage once")
	return fmt.Sprintf("message: %d", id), nil
}

// singleflightGetMessage 使用singleflight的方式调用getMessage
func singleflightGetMessage(sg *singleflight.Group, id int) (string, error) {
	v, err, _ := sg.Do(fmt.Sprintf("%d", id), func() (interface{}, error) {
		return getMessage(id)
	})
	return v.(string), err
}
func TestSingleFlight(t *testing.T) {
	time.AfterFunc(time.Second, func() {
		atomic.AddInt32(&count, -count)
	})
	var (
		wg  sync.WaitGroup
		now = time.Now()
		n   = 1000
		sg  = &singleflight.Group{}
	)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			res, _ := singleflightGetMessage(sg, 1)
			//res, _ := getMessage(1)
			if res != "message: 1" {
				panic("err")
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("同时发起%d次请求，耗时：%s\n", n, time.Since(now))
}

func TestName(t *testing.T) {
	wg := sync.WaitGroup{}
	//wg.Add(1)
	//go func() {
	//	time.Sleep(time.Second)
	//	wg.Done()
	//}()
	//wg.Wait()
	wg.Wait()
	fmt.Println("1")
}

func TestPanic(t *testing.T) {
	go func() {
		//panic("err")
	}()
	go func() {
		panic("err")
	}()
	time.Sleep(time.Second)
}

func TestPanicWithSelect(t *testing.T) {
	go func() {
		go panic("err")
		select {}
	}()
	time.Sleep(time.Second)
}
