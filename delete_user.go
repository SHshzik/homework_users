package main

import (
	"fmt"
	"log/slog"
	"strconv"
	"user_workhome/user"
)

func deleteUser(service user.Service) {
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

	err = service.RemoveUser(id)
	if err != nil {
		fmt.Println("Something went wrong")
		slog.Error(err.Error())
		return
	}
	fmt.Println("User removed")
}
