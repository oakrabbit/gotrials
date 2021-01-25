package idiomatic

import "fmt"

/**

For some tricks see : https://github.com/golang/go/wiki/SliceTricks
for more details on slice tricks
*/
func CreateAndPrintSlice() []int {
	myslice := make([]int, 10, 20)
	fmt.Println("Slice created: ", myslice)
	return myslice
}

func ChangeSliceValue(slice []int) {
	fmt.Println("slice before append", slice)
	slice[1] = 30
	// append returns a new slice Header much like copy
	// a new header and Pointr is returned therefore the append is not refletec in callee method variable
	slice = append(slice, 10)
	fmt.Println("slice after append", slice)

}
func ChangeSliceValueMuttable(slice *[]int) {
	fmt.Println("slice before append", slice)
	*slice = append(*slice, 10)
	fmt.Println("slice after append", slice)
}

func ViewCapacity(s []int) int {
	return cap(s)
}

func ViewLen(s []int) int {
	return len(s)
}

func ManipulateReslice() {
	slice := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println("Original Slice", slice)
	nslice := slice[2:]
	fmt.Println("Sub Slice", nslice)
	// change new value of sub slice
	nslice[1] = "z"
	fmt.Println("Sub Slice after value change", nslice)
	fmt.Println("Orginal Slice varible", slice)

}

func preppendSlice(a []int, b []int) []int {
	out := append(b[:], a[:]...)
	return out
}

func cutSlice(in []int, start int, end int) []int {
	// Just set end as the len of the slice. and cut to the end
	if start >= end {
		end = len(in)
	}

	in = append(in[:(start-1)], in[end:]...)
	return in
}

func deleteSliceItem(in []int, index int) []int {
	in = append(in[:index-1], in[index:]...)
	return in
}
