package main

import "fmt"

type iPrototype interface {
	clone() iPrototype
	show()
}

type node struct{
	Name string
	Description string
}

func (n *node)clone() *node {
	return &node{}
}


func exec() {
	no := node{"no1","description"}

	no2 := no.clone()
	no2.Name ="no2"
	fmt.Println("original",no)
	fmt.Println("proto ",no2)

}