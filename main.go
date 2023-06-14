package main

import "fmt"

func main() {
	InMemoryAuth := InMemoryAuth{
		users: []User{},
	}
	for {
		var answer string
		fmt.Println("1. Login \n2.Register\n3.Exit")
		fmt.Scan(&answer)
		if answer == "1" {
			var username, password string
			fmt.Printf("Enter your username: ")
			fmt.Scan(&username)
			fmt.Printf("Enter your password: ")
			fmt.Scan(&password)
			if InMemoryAuth.Login(username, password) {
				fmt.Println("Login successful")
			} else {
				fmt.Println("Uncorrect username or password !")
			}
		} else if answer == "2" {
			var username, password string
			fmt.Printf("Enter your username: ")
			fmt.Scan(&username)
			fmt.Printf("Enter your password: ")
			fmt.Scan(&password)
			user := User{
				Username: username,
				Password: password,
			}
			err := InMemoryAuth.Register(user)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Registered successfully")
			}
		} else {
			break
		}
	}
}
