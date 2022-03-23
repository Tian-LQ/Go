package condition_loop

import (
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
