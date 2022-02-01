package main

import (
	"fmt"
	"sync"
)

type single struct{
	name string
	schema string
	host string
}
var instance *single
var once sync.Once
func GetInstance() *single{

	once.Do(func(){
		instance = &single{
			name:   "",
			schema: "",
			host:   "",
		}
		fmt.Println("instance now created", instance)

	})
	return instance
}


func SingleInstance(){
	for i := 1; i <=5; i++{
			go func() {
				getInstance := GetInstance()
				fmt.Println("instance now created", &getInstance)
			}()
	}
}
