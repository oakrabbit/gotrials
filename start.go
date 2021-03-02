package main

import (
	"fmt"
	model "github.com/gotrials/models"
	"github.com/gotrials/variatic"
	"github.com/gotrials/cmd"
)

// set a global constant
const Host string = "localhost"

func main() {

	cmd.Execute()

	//stuff to delete

	// create new user struct
	// method 1:
	user1 := &model.User{Name: "user 1"}
	fmt.Println("user 1:", user1)

	// new user using the builder interface
	builder := &model.User{}
	user2 := builder.SetName("user2").SetFullname("lastname").Build()
	fmt.Println("user 2:", user2)
	// update user 2 to user 3 name
	changeName(&user2)
	fmt.Println("user 3:", user2)
	// print using a variadic function
	variatic.Printusers(user1, &user2)
}

func changeName(user *model.User) {
	user.Name = "user 3"
}
