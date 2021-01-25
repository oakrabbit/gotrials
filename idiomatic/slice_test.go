package idiomatic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateSlice(t *testing.T) {
	slice := CreateAndPrintSlice()
	fmt.Println("slice is:", slice)
	t.Log("type =", reflect.TypeOf(slice).Elem())

}

// test with value copy to method to append
// no change to the value should happen
func TestSliceCapacity(t *testing.T) {
	slice := CreateAndPrintSlice()
	ChangeSliceValue(slice)
	t.Log("Slice with added valye:", slice)
	capacity := cap(slice)
	t.Log("capacity", capacity)

	if capacity != 20 {
		t.Fail()
	}
}

func TestSliceCapacityWithAddedCapacity(t *testing.T) {
	slice := CreateAndPrintSlice()
	ChangeSliceValueMuttable(&slice)
	t.Log("Slice with added valye:", slice)
	capacity := cap(slice)
	t.Log("capacity", capacity)
	// because cap is defined as 10, and new value is appended capacitydobles its value to 20.
	if capacity != 20 {
		t.Fail()
	}

}

func TestSliceLenisTen(t *testing.T) {
	slice := CreateAndPrintSlice()
	size := len(slice)
	if size != 10 {
		t.Fail()
	}

}

func TestManipulatingSlice(t *testing.T) {
	ManipulateReslice()
}

func TestPreppendSlice(t *testing.T) {
	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}

	out := preppendSlice(sliceA, sliceB)
	t.Log("Slice out", out)

	validationSlice := []int{4, 5, 6, 1, 2, 3}

	if !reflect.DeepEqual(out, validationSlice) {
		t.Fail()
	}
}

func TestCutSlice(t *testing.T) {
	// this will not do GC on sliceA.
	sliceA := []int{1, 2, 3, 4, 5, 6}
	out := cutSlice(sliceA, 3, 2)
	t.Log(out)
	validationSlice := []int{1, 2}
	if !reflect.DeepEqual(out, validationSlice) {
		t.Fail()
	}
	// reassign all
	validationSlice = []int{1, 2, 5, 6}
	out = cutSlice(sliceA, 3, 4)
	t.Log(out)
	if !reflect.DeepEqual(out, validationSlice) {
		t.Fail()
	}

}

func TestDeleteSliceItem(t *testing.T) {
	sliceA := []int{1, 2, 3, 4, 5, 6}
	out := deleteSliceItem(sliceA, 3)
	sliceA = out
	t.Log(&sliceA)

}
