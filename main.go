package main

import (
	"flag"
	"fmt"
	"strconv"
	"user_workhome/user"
)

func main() {
	d := flag.Bool("debug", false, "debug mode")
	flag.Parse()

	var service user.Service
	if *d {
		service = user.NewService(&user.MockRepository{})
	} else {
		service = user.NewService(&user.InMemoryRepository{})
	}

	var c string
	for loop := true; loop; {
		fmt.Println("Type one of available command: list, find, create, delete, exit:")
		_, err := fmt.Scanln(&c)
		if err != nil {
			panic(err.Error())
		}
		switch c {
		case "list":
			list(service)
		case "find":
			find(service)
		case "create":
		case "delete":
			deleteUser(service)
		case "exit":
			loop = false
		default:
			loop = false
		}
	}
}

func list(service user.Service) {
	users := service.ListUsers()

	fmt.Println(users)
}

func find(service user.Service) {
	var sId string
	fmt.Println("Type user ID: ")
	_, err := fmt.Scanln(&sId)
	if err != nil {
		panic(err.Error())
	}
	id, err := strconv.Atoi(sId)
	if err != nil {
		panic(err.Error())
	}

	fUser, err := service.GetUser(id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(fUser)
	}
}

func deleteUser(service user.Service) {
	var sId string
	fmt.Println("Type user ID: ")
	_, err := fmt.Scanln(&sId)
	if err != nil {
		panic(err.Error())
	}
	id, err := strconv.Atoi(sId)
	if err != nil {
		panic(err.Error())
	}

	err = service.RemoveUser(id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("User removed")
	}
}
