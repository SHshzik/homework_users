package main

import (
	"flag"
	"fmt"
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
		fmt.Println("Type one of available command: list, find, create, delete, exit")
		_, err := fmt.Scanln(&c)
		if err != nil {
			panic(err.Error())
		}
		switch c {
		case "list":
			fmt.Println(service.ListUsers())
		case "find":
		case "create":
		case "delete":
		case "exit":
			loop = false
		default:
			loop = false
		}
	}
	//u := user.User{1, "w", "e", "d", time.Now()}
	//fmt.Printf("%#v\n", u)
}
