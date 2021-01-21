package variatic

import (
	"fmt"
	model "github.com/gotrials/models"
)

func Printusers(users ...*model.User) {
	for i, u := range users {
		fmt.Println("u: ", i, u)
	}
}
