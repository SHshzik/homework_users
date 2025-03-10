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
		service = user.NewService(&user.InMemoryRepository{Users: make(map[int]user.User)})
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
			createUser(service)
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

func createUser(service user.Service) {
	var (
		name  string
		email string
		role  string
	)
	fmt.Println("Type user name, email and role separated by space: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		panic(err.Error())
	}
	_, err = fmt.Scan(&email)
	if err != nil {
		panic(err.Error())
	}
	_, err = fmt.Scan(&role)
	if err != nil {
		panic(err.Error())
	}
	newUser, err := service.CreateUser(name, email, role)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("User successfully created")
	fmt.Println(newUser)
}
