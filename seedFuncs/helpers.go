package seedfuncs

import (
	"errors"

	"github.com/billc-dev/tuango-go/ent"
)

func getUserById(users []*ent.User, id string) (*ent.User, error) {
	for _, u := range users {
		if *u.Username == id {
			return u, nil
		}
	}
	return &ent.User{}, errors.New("could not find user")
}
