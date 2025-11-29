package user

import "errors"

// import User "simple-crud/core/user/user"

var users []User = []User{}

func CreateData(user User) {
	users = append(users, user)
}

func GetAllUsers() []User {
	return users
}

func GetUserById(id int) (*User, error) {
	for i := range users {
		if users[i].ID == id {
			return &users[i], nil
		}
	}
	return nil, errors.New("data not found")
}

func DeleteUser(id int) (User, error) {
	for i, user := range users {
		if user.ID == id {
			user := user
			users = append(users[:i], users[i+1:]...)
			return user, nil
		}
	}

	return User{}, errors.New("data not found")
}

func doEditUser(username string, email string, u *User) {
	u.Username = username
	u.Email = email
}

func EditUser(username string, email string, id int) error {
	user, err := GetUserById(id)

	if err == nil {
		doEditUser(username, email, user)
		return nil
	}

	return err
}
