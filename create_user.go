package main

import (
	"fmt"
	"log/slog"
	"user_workhome/user"
)

func createUser(service user.Service) {
	var (
		name  string
		email string
		role  string
	)
	fmt.Println("Type user name, email and role separated by space: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	_, err = fmt.Scan(&email)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	_, err = fmt.Scan(&role)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	valid := user.ValidRole(role)
	if !valid {
		fmt.Println("Wrong user role")
		return
	}
	newUser, err := service.CreateUser(name, email, role)
	if err != nil {
		fmt.Println("Something went wrong")
		slog.Error(err.Error())
		return
	}
	fmt.Println("User successfully created")
	fmt.Println(newUser)
}
