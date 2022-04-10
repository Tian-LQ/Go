package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Message 类型
type Message string

// Greeter 结构体
type Greeter struct {
	Message Message
}

// Event 结构体
type Event struct {
	Greeter Greeter
}

// NewMessage Message的构造函数
func NewMessage() Message {
	return Message("Hi there!")
}

// NewGreeter Greeter构造函数
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// NewEvent Event构造函数
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (g Greeter) Greet() Message {
	return g.Message
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// 使用wire前
//func main() {
//	message := NewMessage()
//	greeter := NewGreeter(message)
//	event := NewEvent(greeter)
//
//	event.Start()
//}

type Slice []int

func (A Slice) Append(value int) {
	A = append(A, value)
}

// 使用wire后
func main() {
	//event := InitializeEvent()
	//event.Start()
	mSlice := make(Slice, 10, 20)
	mSlice.Append(5)
	mSliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&mSlice))
	mSliceHeader.Len = 11
	fmt.Println(mSlice)
}
