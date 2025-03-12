package main

import (
	"fmt"
	"log/slog"
	"strconv"
	"user_workhome/user"
)

func findUser(service user.Service) {
	var sId string
	fmt.Println("Type user ID: ")
	_, err := fmt.Scanln(&sId)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	id, err := strconv.Atoi(sId)
	if err != nil {
		fmt.Println("Wrong ID type")
		slog.Error(err.Error())
		return
	}

	fUser, err := service.GetUser(id)
	if err != nil {
		fmt.Println("User not found")
		slog.Error(err.Error())
	} else {
		fmt.Println(fUser)
	}
}
