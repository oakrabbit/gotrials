package main

import (
	"fmt"
)


type User struct{
	name string
	lastName string
	birth string
	address string
}

type UserBuilder struct{
	user *User
}



func NewUser() *UserBuilder {
	return &UserBuilder{&User{}}

}
func (u *UserBuilder)Name(n string) *UserBuilder{
	u.user.name = n
	//fmt.Println("user address", u)
	return u
}

func (u *UserBuilder)LastName(last string) *UserBuilder{
	u.user.lastName=last
	return u
}
func (u *UserBuilder)Birth(b string) UserBuilder{
	fmt.Println("user address", &u)
	u.user.birth = b
	return *u
}
func (u *UserBuilder) Address(a string) UserBuilder{
	//fmt.Println("user address", &u)
	u.user.address = a
	return *u
}

func (u *UserBuilder) Build() User{
	//fmt.Println("user address", &u)

	return *u.user
}

func main() {
	fmt.Println("builder pattern")
	user := NewUser().
		Name("alex").
		Address("street A")
	user.LastName("ultimo nome")

	user2:= NewUser()
	user2.Name("edu")

	fmt.Println("builder created : ", user.Build())
	fmt.Println("builder created : ", user2.Build())
}
