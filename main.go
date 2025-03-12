package main

import (
	"flag"
	"fmt"
	"user_workhome/user"
)

func main() {
	dPtr := flag.Bool("debug", false, "debug mode")
	flag.Parse()

	var service user.Service
	if *dPtr {
		service = user.NewService(&user.MockRepository{})
	} else {
		service = user.NewService(&user.InMemoryRepository{Users: make(map[int]user.User)})
	}

	var c string
	for loop := true; loop; {
		fmt.Println("Type one of available command: list, find, create, delete, exit:")
		_, err := fmt.Scanln(&c)
		if err != nil {
			fmt.Println("Type one of available command: list, find, create, delete, exit:")
			continue
		}
		switch c {
		case "list":
			listUsers(service)
		case "find":
			findUser(service)
		case "create":
			createUser(service)
		case "delete":
			deleteUser(service)
		case "exit":
			loop = false
		default:
			fmt.Println("Wrong command")
		}
	}
}
