package extension_reuse

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(name string) {
	p.Speak()
	fmt.Println(" ", name)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	fmt.Print("Wang!")
}

func (d *Dog) SpeakTo(name string) {
	d.p.SpeakTo(name)
}

// 命名嵌套类型
type Cat struct {
	Pet
}

func (c *Cat) Speak() {
	fmt.Print("Miao~")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("tianlq")
}

func TestCat(t *testing.T) {
	cat := new(Cat)
	// 下面两种用法是等价的
	cat.SpeakTo("xusq")
	cat.Pet.SpeakTo("xusq")
	cat.Speak()
}
