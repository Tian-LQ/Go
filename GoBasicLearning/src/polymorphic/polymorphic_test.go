package polymorphic

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return fmt.Sprintf(`fmt.Println("Hello, World!")`)
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloWorld() string {
	return fmt.Sprintf(`System.out.Println("Hello, World!")`)
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T: %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphicProgrammer(t *testing.T) {
	writeFirstProgram(&GoProgrammer{})
	writeFirstProgram(&JavaProgrammer{})
}

func DoSomething(obj interface{}) {
	if val, ok := obj.(int); ok {
		fmt.Printf("obj's type: int, obj's value: %d\n", val)
		return
	}
	if val, ok := obj.(string); ok {
		fmt.Printf("obj's type: string, obj's value: %s\n", val)
		return
	}
	fmt.Println("Unknown type!")
}

func DoSomethingPlus(obj interface{}) {
	switch obj.(type) {
	case int:
		fmt.Printf("obj's type: int, obj's value: %d\n", obj)
	case string:
		fmt.Printf("obj's type: string, obj's value: %s\n", obj)
	default:
		fmt.Println("Unknown type!")
	}
}

func TestEmptyInterface(t *testing.T) {
	DoSomething(1)
	DoSomething("Hello, World")
	DoSomething(1.5)
	t.Log("------------------------")
	DoSomethingPlus(1)
	DoSomethingPlus("Hello, World")
	DoSomethingPlus(1.5)
}
