package time_channel

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(time.Second * 5)
	fmt.Printf("Present time: %v.\n", time.Now())
	// 当定时器到期时，会向timer.C发送一个绝对到期时间
	expirationTime := <-timer.C
	fmt.Printf("Expiration time: %v.\n", expirationTime)
	// timer.Stop()的结果为false，表示定时器此时已经过期
	fmt.Printf("Stop timer: %v.\n", timer.Stop())
}

func TestTimerForTimeout(t *testing.T) {
	intChan := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
		intChan <- 1
	}()
	//select {
	//case e := <-intChan:
	//	fmt.Println("Received: ", e)
	//case <-time.NewTimer(time.Millisecond * 500).C:
	//	fmt.Println("Time out!")
	//}

	// time.After()为超时提供的便捷方式
	// 与上面的select超时等价
	select {
	case e := <-intChan:
		fmt.Println("Received: ", e)
	case <-time.After(time.Millisecond * 500):
		fmt.Println("Time out!")
	}
}

func TestForAndSelect(t *testing.T) {
	intChan := make(chan int, 1)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			intChan <- i
		}
		close(intChan)
	}()

	timeout := time.Millisecond * 500
	var timer *time.Timer
	for {
		if timer == nil {
			timer = time.NewTimer(timeout)
		} else {
			timer.Reset(timeout)
		}
		select {
		case e, ok := <-intChan:
			if !ok {
				fmt.Println("End.")
				return
			} else {
				fmt.Printf("Received :%v.\n", e)
			}
		case <-timer.C:
			fmt.Println("Timeout!")
		}
	}
}

func SayHello() {
	fmt.Println("Hello, Mr.Tian!")
}

func TestTimeAfterFunc(t *testing.T) {
	// 在指定时间后会执行相应的方法(并不会向timer.C的通道内发送元素)
	timer := time.AfterFunc(time.Second, SayHello)
	//<-timer.C		// 这条语句会永远阻塞，因为定时器并不会向其发送元素
	time.Sleep(time.Second * 2)
	// 此时定时器已经停止，timer.Stop()返回false
	fmt.Printf("timer stop: %v.\n", timer.Stop())
}

func Test(t *testing.T) {
	intChan := make(chan int, 1)
	ticker := time.NewTicker(time.Second)
	go func() {
		for _ = range ticker.C {
			select {
			case intChan <- 1:
			case intChan <- 2:
			case intChan <- 3:
			}
		}
		fmt.Println("End. [sender]")
	}()
	var sum int
	for e := range intChan {
		fmt.Printf("Received: %v\n", e)
		sum += e
		if sum > 10 {
			//ticker.Stop()
			fmt.Printf("Got: %v\n", sum)
			break
		}
	}
	fmt.Println("End. [receiver]")
}
