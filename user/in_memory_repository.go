package user

import "errors"

type InMemoryRepository struct {
	users map[int]User
}

func (i *InMemoryRepository) Save(user User) error {
	//TODO implement me
	panic("implement me")
}

func (i *InMemoryRepository) FindByID(id int) (User, error) {
	user, ok := i.users[id]
	if !ok {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func (i *InMemoryRepository) FindAll() []User {
	r := make([]User, len(i.users))
	for _, user := range i.users {
		r = append(r, user)
	}
	return r
}

func (i *InMemoryRepository) DeleteByID(id int) error {
	_, ok := i.users[id]
	if !ok {
		return errors.New("user not found")
	}
	delete(i.users, id)

	return nil
}
