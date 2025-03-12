package main

import (
	"fmt"
	"user_workhome/user"
)

func listUsers(service user.Service) {
	users := service.ListUsers()

	fmt.Println(users)
}
