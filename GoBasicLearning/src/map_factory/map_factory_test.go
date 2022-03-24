package map_factory

import "testing"

func TestMapWithFuncValue(t *testing.T) {
	m := make(map[int]func(op int) int)
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

// bool类型的默认零值为: false
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	t.Log(mySet[1])
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is existing\n", n)
	} else {
		t.Logf("%d is not existing\n", n)
	}
	mySet[3] = true
	t.Logf("mySet's length: %d\n", len(mySet))
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Logf("%d is existing\n", n)
	} else {
		t.Logf("%d is not existing\n", n)
	}
}
