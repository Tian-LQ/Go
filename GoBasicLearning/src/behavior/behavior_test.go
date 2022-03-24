package behavior

import (
	"fmt"
	"testing"
)

type Employee struct {
	Id   int
	Age  int
	Name string
}

func (e *Employee) String() string {
	return fmt.Sprintf("[ Id:%d | Age:%d | Name: %s ]", e.Id, e.Age, e.Name)
}

func TestCreateEmployeeObj(t *testing.T) {
	e1 := Employee{0, 25, "tianlq"}
	e2 := Employee{
		Id:   1,
		Age:  25,
		Name: "xusq",
	}
	e3 := new(Employee)
	t.Log(e1)
	t.Log(e2)
	t.Logf("e1's type: %T\n", e1)
	t.Logf("e3's type: %T\n", e3)
	t.Log(e1.String())
}
