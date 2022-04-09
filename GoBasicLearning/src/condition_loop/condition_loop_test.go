package condition_loop

import (
	"fmt"
	"sync"
	"testing"
)

func TestWhileLoop(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Logf("this is %dth time of the condition_loop\n", i+1)
	}
}

func TestCondition(t *testing.T) {
	m := sync.Map{}
	m.Store("name", "tianlq")
	if val, ok := m.Load("name"); ok {
		t.Logf("[key: value] => [%s: %s]\n", "tianlq", val)
	}
}

func TestSwitchMultipleCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Logf("%d: Even number\n", i)
		case 1, 3:
			t.Logf("%d: Odd number\n", i)
		default:
			t.Logf("%d: Not between 0-3\n", i)
		}
	}
}

// TODO 此时case后的condition表达式结果便限制为必须是布尔值
func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Logf("%d: Even number\n", i)
		case i%2 == 1:
			t.Logf("%d: Odd number\n", i)
		default:
			t.Logf("%d: Unknown\n", i)
		}
	}
}

func TestSwitchCase1(t *testing.T) {
	//value := [...]int8{0, 1, 2, 3, 4, 5, 6}
	//switch 1 + 3 {
	//case value[0], value[1]:
	//	fmt.Println("0 or 1")
	//case value[2], value[3]:
	//	fmt.Println("2 or 3")
	//case value[4], value[5], value[6]:
	//	fmt.Println("4 or 5 or 6")
	//}
}

func TestSwitchCase2(t *testing.T) {
	value := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value[2] {
	case 0, 1:
		fmt.Println("0 or 1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}
}

func TestSwitchCase3(t *testing.T) {
	// 绕过编译器
	value := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value[2] {
	case value[0], value[1], value[2]:
		fmt.Println("0 or 1")
	case value[2], value[3], value[4]:
		fmt.Println("2 or 3")
	case value[4], value[5], value[6]:
		fmt.Println("4 or 5 or 6")
	}
}

type myUInt16 uint16

func TestTypeSwitchCase(t *testing.T) {
	value6 := interface{}(byte(127))
	switch t := value6.(type) {
	case uint16, myUInt16:
		fmt.Println("uint8 or uint16")
	case byte:
		fmt.Println("byte")
	default:
		fmt.Printf("unsupported type: %T\n", t)
	}
}
