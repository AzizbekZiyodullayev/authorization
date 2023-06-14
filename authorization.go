package main

type Auth interface {
	Register(user User) error
	Login(username, password string) bool
}
