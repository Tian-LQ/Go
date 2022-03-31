package unit_test

import (
	"fmt"
	"testing"
)

func Square(i int) int {
	return i * i
}

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := Square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is: %d, result is: %d, expected is: %d\n", i, ret, expected[i])
		}
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("End")
}

func TestFatalInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Error")
	fmt.Println("End")
}
