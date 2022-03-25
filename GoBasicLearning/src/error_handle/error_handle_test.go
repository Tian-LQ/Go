package error_handle

import (
	"errors"
	"fmt"
	"testing"
)

var LessThanOneError = errors.New("input param index should not less than 1")
var LargerThanOneHundredError = errors.New("input param index should not larger than 100")

func Fibonacci(index int) (int, error) {
	if index < 1 {
		return 0, LessThanOneError
	}
	if index > 100 {
		return 0, LargerThanOneHundredError
	}
	former, former_of_former := 1, 1
	for i := 2; i < index; i++ {
		former, former_of_former = former_of_former, former
		former = former + former_of_former
	}
	return former, nil
}

func TestGetFibonacci(t *testing.T) {
	val, err := Fibonacci(-10)
	if err != nil {
		switch {
		case err == LessThanOneError:
			fmt.Println("Wrong Input! It is less than 1.")
		case err == LargerThanOneHundredError:
			fmt.Println("Wrong Input! It is larger than 100.")
		default:
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Printf("The 10th Fibonacci number is: %d\n", val)
}
