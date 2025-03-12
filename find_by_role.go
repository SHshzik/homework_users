package main

import (
	"fmt"
	"log/slog"
	"user_workhome/user"
)

func findByRole(service user.Service) {
	var role string
	fmt.Println("Type role: ")
	_, err := fmt.Scanln(&role)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	valid := user.ValidRole(role)
	if !valid {
		fmt.Println("Wrong user role")
		return
	}

	users := service.FindByRole(role)
	fmt.Println(users)
}
