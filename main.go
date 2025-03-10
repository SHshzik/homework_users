package main

import (
	"fmt"
	"user_workhome/user"
)

func main() {
	var c string
	service := user.NewService(&user.InMemoryRepository{})
	for loop := true; loop; {
		fmt.Println("Type one of available command: list, exit")
		_, err := fmt.Scanln(&c)
		if err != nil {
			panic(err.Error())
		}
		switch c {
		case "list":
			fmt.Println(service.ListUsers())
		case "exit":
			loop = false
		default:
			loop = false
		}
	}
	//u := user.User{1, "w", "e", "d", time.Now()}
	//fmt.Printf("%#v\n", u)
}
