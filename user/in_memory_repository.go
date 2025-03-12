package user

import "fmt"

type InMemoryRepository struct {
	Users  map[int]User
	lastId int
}

func (i *InMemoryRepository) Save(user *User) error {
	user.ID = i.lastId
	i.Users[i.lastId] = *user
	i.lastId++

	return nil
}

func (i *InMemoryRepository) FindByID(id int) (User, error) {
	user, ok := i.Users[id]
	if !ok {
		return user, fmt.Errorf("user not found")
	}
	return user, nil
}

func (i *InMemoryRepository) FindAll() []User {
	r := make([]User, 0, len(i.Users))
	for _, user := range i.Users {
		r = append(r, user)
	}
	return r
}

func (i *InMemoryRepository) DeleteByID(id int) error {
	_, ok := i.Users[id]
	if !ok {
		return fmt.Errorf("user not found")
	}
	delete(i.Users, id)

	return nil
}

func (i *InMemoryRepository) FindByRole(role string) []User {
	r := make([]User, 0)
	for _, user := range i.Users {
		if user.Role == role {
			r = append(r, user)
		}
	}
	return r
}
