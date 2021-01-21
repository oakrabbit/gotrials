package variatic

import (
	"fmt"
	model "github.com/gotests/models"
)

func Printusers(users ...*model.User) {
	for i, u := range users {
		fmt.Println("u: ", i, u)
	}
}
