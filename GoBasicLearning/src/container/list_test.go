package container

import (
	"container/list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	l := list.List{}
	l.PushBack(1)
	t.Log(l.Back())
}

type MyInt int

type MyFunc func()

func Hello() {
	fmt.Println("Hello.")
}

type MyMap map[string]string

func TestTransfer(t *testing.T) {
	var a int = 1
	t.Logf("%T\n", a)
	var b MyInt = MyInt(a)
	t.Logf("%T\n", b)
	c := Hello
	t.Logf("%T\n", c)
	var d MyFunc = c
	t.Logf("%T\n", d)
	m1 := make(map[string]string)
	t.Logf("%T\n", m1)
	var m2 MyMap = m1
	t.Logf("%T\n", m2)
}

type Cat struct {
	name string
}

func (c Cat) String() string {
	return fmt.Sprintf("my name is %s.", c.name)
}

func (c Cat) Hello() {
	fmt.Println("Hello, I am a Cat.")
}

type Animal struct {
	Cat
}

func (a Animal) String() string {
	return fmt.Sprint("I am Animal.")
}

func TestStruct(t *testing.T) {
	a := Animal{
		Cat{name: "fuck you"},
	}
	fmt.Printf("%s\n", a)
	a.Hello()
}
