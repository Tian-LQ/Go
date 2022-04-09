package array_slice

import (
	"fmt"
	"testing"
)

func TestArrayInit(t *testing.T) {
	var arr [3]int
	t.Logf("arr's type: %T, arr: %v\n", arr, arr)
	arr1 := [4]int{1, 2, 3, 4}
	t.Logf("arr1's type: %T, arr1: %v\n", arr1, arr1)
	arr2 := [4]int{1, 2}
	t.Logf("arr2's type: %T, arr2: %v\n", arr2, arr2)
	arr3 := [...]int{1, 2, 3, 4, 5}
	t.Logf("arr3's type: %T, arr3: %v\n", arr3, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		t.Logf("arr[%d]: %d\n", i, arr[i])
	}
}

func TestArrayRange(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	for i, val := range arr {
		t.Logf("arr[%d]: %d\n", i, val)
	}
	t.Logf("arr[1:2] type: %T\n", arr[1:2])
}

func TestArraySubSlice(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	subSlice := arr[:3]
	t.Logf("arr's type: %T, arr's value: %v\n", arr, arr)
	t.Logf("subSlice's type: %T, subSlice's value: %v\n", subSlice, subSlice)
}

func TestSliceInit(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	t.Logf("slice's length: %d, slice's capacity: %d, slice's value: %v\n", len(slice), cap(slice), slice)
	slice = append(slice, 6)
	t.Logf("slice's length: %d, slice's capacity: %d, slice's value: %v\n", len(slice), cap(slice), slice)
	slice1 := make([]int, 4)
	t.Logf("slice1's length: %d, slice1's capacity: %d, slice1's value: %v\n", len(slice1), cap(slice1), slice1)
	slice2 := make([]int, 0, 4)
	t.Logf("slice2's length: %d, slice2's capacity: %d, slice2's value: %v\n", len(slice2), cap(slice2), slice2)
}

func TestSliceGrowing(t *testing.T) {
	slice := make([]int, 0)
	for i := 0; i < 10; i++ {
		slice = append(slice, i)
		t.Logf("len: %d, cap: %d, val: %v\n", len(slice), cap(slice), slice)
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	t.Logf("year => len: %d, cap: %d, val: %v\n", len(year), cap(year), year)
	summer := year[5:8]
	t.Logf("summer => len: %d, cap: %d, val: %v\n", len(summer), cap(summer), summer)
	autumn := year[8:11]
	t.Logf("autumn => len: %d, cap: %d, val: %v\n", len(autumn), cap(autumn), autumn)
	summer[0] = "Unknown"
	t.Log("after change summer[0]:")
	t.Logf("summer => len: %d, cap: %d, val: %v\n", len(summer), cap(summer), summer)
	t.Logf("year => len: %d, cap: %d, val: %v\n", len(year), cap(year), year)
}

func TestSlice(t *testing.T) {
	slice := make([]int, 0, 4)
	fmt.Printf("slice's address: %p, slice's len: %d, slice's cap: %d\n", slice, len(slice), cap(slice))
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	fmt.Printf("slice's address: %p, slice's len: %d, slice's cap: %d\n", slice, len(slice), cap(slice))
	slice = append(slice, 5)
	fmt.Printf("slice's address: %p, slice's len: %d, slice's cap: %d\n", slice, len(slice), cap(slice))
}

func TestForRangeArray(t *testing.T) {
	array := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex := len(array) - 1
	for i, e := range array {
		if i == maxIndex {
			array[0] += e
		} else {
			array[i+1] += e
		}
	}
	fmt.Println(array)
}

func TestForRangeSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	maxIndex := len(slice) - 1
	for i, e := range slice {
		if i == maxIndex {
			slice[0] += e
		} else {
			slice[i+1] += e
		}
	}
	fmt.Println(slice)
}
