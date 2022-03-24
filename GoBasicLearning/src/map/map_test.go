package _map

import "testing"

func TestMapInit(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Logf("m1[\"one\"]: %d\n", m1["one"])
	t.Logf("m1's len: %d\n", len(m1))
	t.Logf("m1's type: %T\n", m1)
	m2 := make(map[string]int, 8)
	t.Logf("m2's len: %d\n", len(m2))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	// 如果访问map当中不存在的key值，会返回value类型默认零值
	t.Log(m1[1])
	if val, ok := m1[1]; ok {
		t.Logf("m[1]: %d\n", val)
	} else {
		t.Log("m1 don't have key: 1")
	}
	m2 := map[int]string{}
	t.Logf("m2[1]: [%s]\n", m2[1])
}

func TestMapTravel(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, val := range m1 {
		t.Logf("m1[\"%s\"]: %d\n", key, val)
	}
}
