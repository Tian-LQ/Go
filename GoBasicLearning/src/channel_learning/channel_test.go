package channel_learning

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/**********************************/
var strChan = make(chan string, 3)

func TestChannel1(t *testing.T) {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	// 接收操作
	go func() {
		<-syncChan1
		fmt.Println("Received a sync signal and wait a second... [receiver]")
		time.Sleep(time.Second)
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("Received:", elem, "[receiver]")
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan2 <- struct{}{}
	}()

	// 发送操作
	go func() {
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("Sent:", elem, "[sender]")
			if elem == "c" {
				syncChan1 <- struct{}{}
				fmt.Println("Sent a sync signal.[sender]")
			}
		}
		fmt.Println("Wait 2 seconds...[sender]")
		time.Sleep(time.Second * 2)
		close(strChan)
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2
}

/**********************************/
var mapChan = make(chan map[string]int, 1)

func TestChannel2(t *testing.T) {
	syncChan := make(chan struct{}, 2)
	// 接收操作
	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				elem["count"]++
			} else {
				break
			}
		}
		fmt.Println("Stopped.[receiver]")
		syncChan <- struct{}{}
	}()

	// 发送操作
	go func() {
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

/**********************************/
type Counter struct {
	count int
}

var newMapChan = make(chan map[string]*Counter, 1)

func TestChannel3(t *testing.T) {
	syncChan := make(chan struct{}, 2)
	// 接收操作
	go func() {
		for {
			if elem, ok := <-newMapChan; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("Stopped.[receiver]")
		syncChan <- struct{}{}
	}()

	// 发送操作
	go func() {
		countMap := map[string]*Counter{
			"count": &Counter{},
		}
		for i := 0; i < 5; i++ {
			newMapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap["count"])
		}
		close(newMapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

/**********************************/
func TestChannel4(t *testing.T) {
	dataChan := make(chan int, 5)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	// 接收操作
	go func() {
		<-syncChan1
		for {
			if elem, ok := <-dataChan; ok {
				fmt.Println("Received:", elem, "[receiver]")
			} else {
				break
			}
		}
		fmt.Println("Done. [receiver]")
		syncChan2 <- struct{}{}
	}()

	// 发送操作
	go func() {
		for i := 0; i < 5; i++ {
			dataChan <- i
			fmt.Printf("Sent: %d [sender]\n", i)
		}
		close(dataChan)
		syncChan1 <- struct{}{}
		fmt.Println("Done. [sender]")
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2
}

/**********************************/
var secondStrChan = make(chan string, 3)

// 接收操作
func receive(strChan <-chan string,
	syncChan1 <-chan struct{},
	syncChan2 chan<- struct{}) {
	<-syncChan1
	fmt.Println("Received a sync signal and wait a second... [receiver]")
	time.Sleep(time.Second)
	for {
		if elem, ok := <-strChan; ok {
			fmt.Println("Received:", elem, "[receiver]")
		} else {
			break
		}
	}
	fmt.Println("Stopped. [receiver]")
	syncChan2 <- struct{}{}
}

// 接收操作Plus
func receivePlus(strChan <-chan string,
	syncChan1 <-chan struct{},
	syncChan2 chan<- struct{}) {
	<-syncChan1
	fmt.Println("Received a sync signal and wait a second... [receiver]")
	time.Sleep(time.Second)
	for elem := range strChan {
		fmt.Println("Received:", elem, "[receiver]")
	}
	fmt.Println("Stopped. [receiver]")
	syncChan2 <- struct{}{}
}

// 发送操作
func send(strChan chan<- string,
	syncChan1 chan<- struct{},
	syncChan2 chan<- struct{}) {
	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		fmt.Println("Sent:", elem, "[sender]")
		if elem == "c" {
			syncChan1 <- struct{}{}
			fmt.Println("Sent a sync signal.[sender]")
		}
	}
	fmt.Println("Wait 2 seconds...[sender]")
	time.Sleep(time.Second * 2)
	close(strChan)
	syncChan2 <- struct{}{}
}

func TestChannelPlus(t *testing.T) {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go receive(secondStrChan, syncChan1, syncChan2)
	go send(secondStrChan, syncChan1, syncChan2)
	<-syncChan2
	<-syncChan2
}

/**********************************/
func TestSingleTrackChannel(t *testing.T) {
	var ok bool
	ch := make(chan int, 1)
	_, ok = interface{}(ch).(<-chan int)
	fmt.Println("[chan int] => [<-chan int]:", ok)
	_, ok = interface{}(ch).(chan<- int)
	fmt.Println("[chan int] => [chan<- int]:", ok)

	sendCh := make(chan<- int, 1)
	_, ok = interface{}(sendCh).(chan int)
	fmt.Println("[chan<- int] => [chan int]:", ok)

	receiveCh := make(<-chan int, 1)
	_, ok = interface{}(receiveCh).(chan int)
	fmt.Println("[<-chan int] => [chan int]:", ok)
}

func TestSelect(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	// 当所有case均满足时，随机命中case
	select {
	case ch1 <- 1:
		fmt.Println("ch1 receive value 1")
	case ch2 <- 2:
		fmt.Println("ch2 receive value 2")
	}
}

// 向一个未初始化(nil)的channel发送或接收元素均会引发运行时恐慌
func TestChannelNil(t *testing.T) {
	//var ch chan int
	//ch <- 1
	//val := <-ch
	//t.Log(val)
}

func TestChannelAndChannelBuffer(t *testing.T) {
	// 发送间隔时间
	sendingInterval := time.Second
	// 接收间隔时间
	receptionInterval := time.Second * 2
	// 采用非缓存通道时，其收发元素的速度总是与慢的那一方持平
	//intChan := make(chan int, 0)
	// 采用缓存通道时，接收元素的速度与发送元素的速度相互独立
	intChan := make(chan int, 5)

	// 发送Goroutine
	go func() {
		var ts0, ts1 int64
		for i := 0; i < 5; i++ {
			intChan <- i
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("Sent:", i)
			} else {
				fmt.Printf("Send: %d [interval: %d s]\n", i, ts1-ts0)
			}
			ts0 = time.Now().Unix()
			// 模拟发送元素的速度：1个/s
			time.Sleep(sendingInterval)
		}
		close(intChan)
	}()

	var ts0, ts1 int64
Loop:
	for {
		select {
		case v, ok := <-intChan:
			if !ok {
				break Loop
			}
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("Received:", v)
			} else {
				fmt.Printf("Received: %d [interval: %d]\n", v, ts1-ts0)
			}
		}
		ts0 = time.Now().Unix()
		// 模拟接收元素的速度：1个/s
		time.Sleep(receptionInterval)
	}
	fmt.Println("End.")
}

func TestSelected(t *testing.T) {
	t.Log("moment 1:", runtime.NumGoroutine())
	c := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		c <- 10086
	}()
	t.Log("moment 2:", runtime.NumGoroutine())
	select {
	case i := <-c:
		fmt.Println("Received form c:", i)
	case <-time.After(time.Second):
		t.Log("moment 3:", runtime.NumGoroutine())
		fmt.Println(<-c)
		t.Log("moment 4:", runtime.NumGoroutine())
	}
	t.Log("moment 5:", runtime.NumGoroutine())
}
