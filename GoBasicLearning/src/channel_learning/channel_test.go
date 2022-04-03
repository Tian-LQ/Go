package channel_learning

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"sync"
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

/************************************************/

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

/************************************************/
// Go Concurrency Pattern: Moving On
func Query(input []string) string {
	ch := make(chan string)
	for _, s := range input {
		go func(str string) {
			select {
			case ch <- str:
				fmt.Printf("ready for send %s. [sender]\n", str)
			default:
				fmt.Printf("send failed: %s. [sender]\n", str)
			}
		}(s)
	}
	fmt.Println("ready for receive. [receiver]")
	return <-ch
}

func TestChannelForMovingOn(t *testing.T) {
	input := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
	t.Log(Query(input))
}

/************************************************/
// Go Concurrency Pattern: Time Out

func TestChannelForTimeout(t *testing.T) {
	timeout := make(chan struct{}, 1)
	// 启动了一个类似定时器的G
	go func() {
		time.Sleep(time.Second * 1)
		timeout <- struct{}{}
	}()
	ch := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 2)
		ch <- 1019
		fmt.Println("send value. [sender]")
	}()

	select {
	case <-ch:
		fmt.Println("receive value. [receiver]")
	case <-timeout:
		fmt.Println("time out")
	}
}

/************************************************/
// Go Concurrency Pattern: Pipeline

// 定义入站通道，通过入站通道从上游接收元素
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

// Use channel buffer
func genWithChannelBuffer(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	// 由于非阻塞的out <- num操作，因此无需创建新的goroutine
	for _, num := range nums {
		out <- num
	}
	close(out)
	return out
}

// 对入站通道当中的数据进行某些处理，生成新值并存入出站通道
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// 通过出站通道向下游发送值
func TestPipeline(t *testing.T) {
	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

// 由于中间过程sq方法的入站和出站通道具有相同类型，因此我们可以组合多次
func TestPipelinePlus(t *testing.T) {
	// Set up the pipeline and consume the output.
	for num := range sq(sq(gen(2, 3))) {
		fmt.Println(num) // 16 then 81
	}
}

/************************************************/
// Go Concurrency Pattern: Fan-out, Fan-in
// Fan-out: 多个方法可以从同一个channel当中接收数据直到该channel关闭
// Fan-in: 一个方法可以从多个input channel当中读取数据并且将其复用至一个单独的channel，直到他们都关闭

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func TestFanOutAndFanIn(t *testing.T) {
	// Fan-out:
	// Distribute the sq work across two goroutines that both read from in.
	in := gen(2, 3)
	c1 := sq(in)
	c2 := sq(in)

	// Fan-in
	// Consume the merged output from c1 and c2
	for n := range merge(c1, c2) {
		fmt.Printf("receive value : %d. [receiver]\n", n) // 4 then 9, or 9 then 4
	}
}

// TODO 管道的模式
// 1.阶段在完成所有发送操作后关闭其出站通道
// 2.阶段不断从入站通道接收值，直到这些通道关闭
// 总而言之：入站channel先关闭，紧跟着出站channel再关闭

// TODO 思考
// 实际管道使用过程中，每个阶段并不是接收所有入站值，接收方可能只需要部分值即可取得进展
// 亦或是某个阶段会因为入站值表示某种错误而提前退出，在这样的情况下则会导致发送这些值的
// goroutine无限期阻塞，从而产生泄露
// 我们通过尝试修改入站channel为buffer channel的形式来解决阻塞的问题，但是这很脆弱
// 因此我们需要为下游阶段提供一种方法，能够告知上游发送方"我已经停止接收输入"

/************************************************/
// Go Concurrency Pattern: Explicit cancellation

func sqPlus(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
		close(out)
	}()
	return out
}

func mergePlus(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func TestExplicitCancellation(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	// Fan-out:
	// Distribute the sq work across two goroutines that both read from in.
	in := gen(2, 3)
	c1 := sqPlus(done, in)
	c2 := sqPlus(done, in)

	// Fan-in
	// Consume the merged output from c1 and c2
	out := mergePlus(done, c1, c2)
	fmt.Println(<-out)
}

// TODO 使用管道建议
// 1.阶段在完成所有发送操作后关闭其出站通道
// 2.阶段不断从入站通道接收值，直到这些通道关闭或发送方被解除阻止

/************************************************/
// Pipeline : Digesting a tree

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
	c := make(chan result)
	errc := make(chan error, 1)
	go func() {
		var wg sync.WaitGroup
		err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			wg.Add(1)
			go func() {
				data, err := ioutil.ReadFile(path)
				select {
				case c <- result{path, md5.Sum(data), err}:
				case <-done:
				}
			}()
			select {
			case <-done:
				return errors.New("walk canceled")
			default:
				return nil
			}
		})
		go func() {
			wg.Wait()
			close(c)
		}()
		errc <- err
	}()
	return c, errc
}

func MD5All(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

	c, errc := sumFiles(done, root)
	m := make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}
	if err := <-errc; err != nil {
		return nil, err
	}
	return m, nil
}

func TestMD5(t *testing.T) {
	if m, err := MD5All("C:\\Users\\田磊泉\\Desktop\\Go_Learning_File\\学习笔记思维导图"); err != nil {
		fmt.Printf("%v\n", m)
	}
}

func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)
	go func() {
		defer close(paths)
		errc <- filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path:
			case <-done:
				return errors.New("walk canceled")
			}
			return nil
		})
	}()
	return paths, errc
}

func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
	for path := range paths {
		data, err := ioutil.ReadFile(path)
		select {
		case c <- result{path, md5.Sum(data), err}:
		case <-done:
			return
		}
	}
}

func MD5AllPlus(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)
	paths, errc := walkFiles(done, root)
	if err := <-errc; err != nil {
		return nil, err
	}

	c := make(chan result)
	var wg sync.WaitGroup
	const numDigesters = 20
	wg.Add(numDigesters)
	for i := 0; i < numDigesters; i++ {
		go func() {
			digester(done, paths, c)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(c)
	}()

	m := make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}

	if err := <-errc; err != nil {
		return nil, err
	}

	return m, nil
}

func TestMD5Plus(t *testing.T) {

}
