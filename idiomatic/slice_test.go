package slice

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
