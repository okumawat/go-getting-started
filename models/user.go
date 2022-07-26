package models

import "errors"

type User struct {
	Id    int
	Name  string
	Email string
}

var (
	users  []*User
	nextId int = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.Id != 0 {
		return User{}, errors.New("User shouldn't contain id")
	}
	u.Id = nextId
	nextId++
	users = append(users, &u)
	return u, nil
}

func GetUserById(id int) (User, error) {
	for _, user := range users {
		if user.Id == id {
			return *user, nil
		}
	}
	return User{}, errors.New("id not found")
}
