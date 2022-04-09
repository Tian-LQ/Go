package _interface

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() string {
	return fmt.Sprintf(`fmt.Println("Hello World")`)
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

// 自定义类型
type MyFunction func(op int) int

func timeSpent(inner MyFunction) MyFunction {
	return func(n int) int {
		start := time.Now()
		res := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return res
	}
}

func Square(n int) int {
	time.Sleep(time.Second * 1)
	return n * n
}

func TestFunc(t *testing.T) {
	fn := timeSpent(Square)
	t.Log(fn(10))
	t.Logf("Square's type: %T\n", Square)
}

type Pet interface {
	Name() string
	Say() string
}

type Cat struct {
	name string
}

type Dog struct {
	name string
}

func (d Dog) Name() string {
	return d.name
}

func (c Cat) Name() string {
	return c.name
}

func (c Cat) Say() string {
	return "miao~"
}

func (c *Cat) SetName(newName string) {
	c.name = newName
}

func TestInterfacePet(t *testing.T) {
	//cat := Cat{name: "Tom"}
	//fmt.Printf("%+v\n", cat)
	//var pet Pet = cat
	//fmt.Printf("%+v\n", pet)
	//cat.SetName("Jerry")
	//fmt.Printf("%+v\n", cat)
	//fmt.Printf("%+v\n", pet)

	// 做个实验
	var cat1 *Cat
	cat2 := cat1
	var pet Pet = cat2
	if cat1 == nil {
		fmt.Println("cat1 is nil")
	}
	if cat2 == nil {
		fmt.Println("cat2 is nil")
	}
	if pet == nil {
		fmt.Println("pet is nil")
	}
	if pet == cat2 {
		fmt.Println("pet == cat2")
	}
	pet.Say()
	// 只要我们把一个有类型的nil赋给接口变量
	// 这个变量的值就一定不会是那个真正的nil(iface实例)
	//fmt.Printf("%+v\n", cat1)
	//fmt.Printf("%+v\n", cat2)
	fmt.Printf("%T\n", pet)
	fmt.Println(reflect.TypeOf(pet))

	// 声明一个接口但不做初始化，或者直接对其赋值字面量nil
	// 此时这个接口变量才是真正的nil
	var pet2 Pet = nil
	if pet2 == nil {
		fmt.Println("pet2 is nil")
	}
}

type MyStruct struct {
}

func (s *MyStruct) Hello() {
	fmt.Println("Hello World!")
}

func NewMyStruct() MyStruct {
	return MyStruct{}
}

func TestMyStruct(t *testing.T) {
	//NewMyStruct().Hello()
}

func TestGoroutine(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestGoroutineOrderlyExecute(t *testing.T) {
	var count uint32 = 0
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestGoroutineOrderlyExecuteAnother(t *testing.T) {
	//ch := make(chan struct{})
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		fmt.Println(i)
	//		ch <- struct{}{}
	//	}(i)
	//	<-ch
	//}
}
