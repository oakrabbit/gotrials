package main

import (
	"fmt"
	"testing"
)

func Test_node_clone(t *testing.T) {
	no := node{"no1","description"}

	no2 := no.clone()
	no2.Name ="no2"
	fmt.Println("original",no)
	fmt.Println("proto ",no2)
	t.Log("OK")
}
