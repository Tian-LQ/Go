package main

import (
	"firstProject/internal/grpc_class/pb/person"
	"fmt"
	"reflect"
)

func main() {
	var p = person.Person{
		Name:      "Mr.T",
		Age:       25,
		Sex:       person.Person_MAN,
		Test:      make([]string, 0, 4),
		TestMap:   make(map[string]string),
		TestOneOf: &person.Person_One{One: "person one"},
	}
	one, ok := p.TestOneOf.(*person.Person_One)
	if !ok {
		fmt.Println("The type of p.TestOneOf is not *person.Person_One")
		fmt.Printf("The type of p.TestOneOf is: %v", reflect.TypeOf(p.TestOneOf))
		return
	}
	fmt.Println(one)
	switch p.TestOneOf.(type) {
	case *person.Person_One:
		fmt.Println("select case *person.Person_One [switch]")
		one, _ := p.TestOneOf.(*person.Person_One)
		fmt.Printf("type of p.TestOneof: %v, value of p.TestOneof: %+v\n", reflect.TypeOf(one), one)
	case *person.Person_Two:
		fmt.Println("select case *person.Person_Two [switch]")
		two, _ := p.TestOneOf.(*person.Person_Two)
		fmt.Printf("type of p.TestOneof: %v, value of p.TestOneof: %+v\n", reflect.TypeOf(two), two)
	case *person.Person_Three:
		fmt.Println("select case *person.Person_Three [switch]")
		three, _ := p.TestOneOf.(*person.Person_Three)
		fmt.Printf("type of p.TestOneof: %v, value of p.TestOneof: %+v\n", reflect.TypeOf(three), three)
	}
}
