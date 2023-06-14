package main

import "errors"

type InMemoryAuth struct {
	users []User
}

func (a *InMemoryAuth) Register(u User) error {
	for _, user := range a.users {
		if user.Username == u.Username {
			return errors.New("This user is already exist")
		}
	}

	a.users = append(a.users, u)
	return nil
}

func (a *InMemoryAuth) Login(username, password string) bool {
	for _, user := range a.users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}
