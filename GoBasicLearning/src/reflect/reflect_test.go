package reflect

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func CheckType(obj interface{}) {
	t := reflect.TypeOf(obj)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	case reflect.String:
		fmt.Println("String")
	default:
		fmt.Println("Unknown")
	}
}

func TestTypeAndValue(t *testing.T) {
	var v int64 = 10
	t.Log(reflect.TypeOf(v), reflect.ValueOf(v))
	t.Log(reflect.ValueOf(v).Type())
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(&f)
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{
		EmployeeID: "1",
		Name:       "Mike",
		Age:        30,
	}
	t.Logf("Name: value(%[1]v), Type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age:", e)
}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	//t.Log(a == b)	// invalid operation: a == b (map can only be compared to nil)
	t.Logf("a == b? %v\n", reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{3, 2, 1}
	t.Logf("s1 == s2? %v\n", reflect.DeepEqual(s1, s2))
	t.Logf("s1 == s3? %v\n", reflect.DeepEqual(s1, s3))
}

func fillBySetting(st interface{}, settings map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to the struct type")
		}
	}
	if settings == nil {
		return errors.New("settings is nil")
	}
	var (
		field reflect.StructField
		ok    bool
	)
	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 40}
	e := Employee{}
	if err := fillBySetting(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Logf("e: %+v\n", e)

	c := new(Customer)
	if err := fillBySetting(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Logf("c: %+v\n", *c)
}

func TestReflect(t *testing.T) {
	e := &Employee{}
	t.Log(reflect.TypeOf(e).Kind())
	t.Log(reflect.TypeOf(e).Elem().Kind())
}
