package user

import (
	"errors"
	"sync/atomic"
)

var idCounter int32 = 0

func RegisterUser(u User) (User, error) {
	if u.Username == "" {
		return User{}, errors.New("username required")
	}

	for _, user := range GetAllUsers() {
		if user.Email == u.Email {
			return User{}, errors.New("email telah digunakan")
		}
	}

	newID := int(atomic.AddInt32(&idCounter, 1))
	u.ID = newID

	CreateData(u)
	return u, nil
}
