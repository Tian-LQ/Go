package variable_constant

import (
	"fmt"
	"testing"
)

func Fibonacci(index int) int {
	former, former_of_former := 1, 1
	// 这个地方也可以考虑用加减法/异或法不引入临时变量tem
	var tem int
	for i := 2; i < index; i++ {
		tem = former + former_of_former
		former_of_former = former
		former = tem
	}
	return former
}

func ExchangeFirst(val1 *int, val2 *int) {
	*val1 = *val1 + *val2
	*val2 = *val1 - *val2
	*val1 = *val1 - *val2
}

func ExchangeSecond(val1 *int, val2 *int) {
	*val1 = *val1 ^ *val2
	*val2 = *val1 ^ *val2
	*val1 = *val1 ^ *val2
}

func ExchangeThird(val1 *int, val2 *int) {
	tem := *val1
	*val1 = *val2
	*val2 = tem
}

func ExchangeGo(val1 *int, val2 *int) {
	*val1, *val2 = *val2, *val1
}

// 快速设置连续值
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// Open = 1, Close = 2, Pending = 4
const (
	Open = 1 << iota
	Close
	Pending
)

func TestFirstTry(t *testing.T) {
	fmt.Printf("The 6th number of Fibonacci is: %d\n", Fibonacci(6))
	fmt.Println("*************************************************")
	a, b := 1, 2
	fmt.Printf("Before ExchangeFirst a and b ==> a : %d, b : %d\n", a, b)
	ExchangeFirst(&a, &b)
	fmt.Printf("After ExchangeFirst a and b ==> a : %d, b : %d\n", a, b)
	fmt.Println("*************************************************")
	c, d := 3, 4
	fmt.Printf("Before ExchangeSecond c and d ==> c : %d, d : %d\n", c, d)
	ExchangeSecond(&c, &d)
	fmt.Printf("After ExchangeSecond c and d ==> c : %d, d : %d\n", c, d)
	fmt.Println("*************************************************")
	e, f := 5, 6
	fmt.Printf("Before ExchangeThird e and f ==> e : %d, f : %d\n", e, f)
	ExchangeThird(&e, &f)
	fmt.Printf("After ExchangeThird e and f ==> e : %d, f : %d\n", e, f)
	fmt.Println("*************************************************")
	h, i := 7, 8
	fmt.Printf("Before ExchangeGo h and i ==> h : %d, i : %d\n", h, i)
	ExchangeGo(&h, &i)
	fmt.Printf("After ExchangeGo h and i ==> h : %d, i : %d\n", h, i)
	fmt.Println("*************************************************")
	fmt.Printf("Monday: %d, Tuesday: %d, Wednesday: %d, Thursday: %d, Friday: %d, Saturday: %d, Sunday: %d\n", Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
	fmt.Printf("Open: %d, Close: %d, Pending: %d\n", Open, Close, Pending)
}
