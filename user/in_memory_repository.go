package user

type InMemoryRepository struct {
	users map[int]User
}

func (i *InMemoryRepository) Save(user User) error {
	//TODO implement me
	panic("implement me")
}

func (i *InMemoryRepository) FindByID(id int) (User, error) {
	//TODO implement me
	panic("implement me")
}

func (i *InMemoryRepository) FindAll() []User {
	r := make([]User, len(i.users))
	for _, user := range i.users {
		r = append(r, user)
	}
	return r
}

func (i *InMemoryRepository) DeleteByID(id int) error {
	//TODO implement me
	panic("implement me")
}
