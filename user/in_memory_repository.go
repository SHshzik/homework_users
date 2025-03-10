package user

import (
	"errors"
	"fmt"
)

type InMemoryRepository struct {
	Users  map[int]User
	lastId int
}

func (i *InMemoryRepository) Save(user *User) error {
	fmt.Println(i.Users)
	user.ID = i.lastId
	i.Users[i.lastId] = *user
	i.lastId++
	return nil
}

func (i *InMemoryRepository) FindByID(id int) (User, error) {
	user, ok := i.Users[id]
	if !ok {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func (i *InMemoryRepository) FindAll() []User {
	fmt.Println(i.Users)
	fmt.Println(len(i.Users))
	r := make([]User, 0, len(i.Users))
	for _, user := range i.Users {
		r = append(r, user)
	}
	return r
}

func (i *InMemoryRepository) DeleteByID(id int) error {
	_, ok := i.Users[id]
	if !ok {
		return errors.New("user not found")
	}
	delete(i.Users, id)

	return nil
}
