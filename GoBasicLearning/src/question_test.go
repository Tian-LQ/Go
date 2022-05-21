package main

import "testing"

func function(s []int) int {
	i := 0
	for ; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			break
		}
	}
	if i == len(s)-1 {
		return 0
	}
	return i + 1
}

func TestFunctionName(t *testing.T) {
	s := []int{1, 2, 3}
	t.Log(function(s))
}

func function2(s []int) bool {
	if s == nil || len(s) <= 2 {
		return true
	}
	m := make(map[int]bool)
	min, max := s[0], s[0]
	for _, val := range s {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
		m[val] = true
	}
	x := (max - min) % (len(s) - 1)
	if x != 0 {
		return false
	}
	n := (max - min) / (len(s) - 1)
	for i := 0; i < len(s); i++ {
		_, ok := m[min+i*n]
		if !ok {
			return false
		}
	}
	return true
}

func TestFunction2Name(t *testing.T) {
	s := []int{1}
	t.Log(function2(s))

}
