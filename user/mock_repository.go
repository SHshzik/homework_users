package user

type MockRepository struct{}

func (m *MockRepository) Save(user User) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository) FindByID(id int) (User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository) FindAll() []User {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository) DeleteByID(id int) error {
	//TODO implement me
	panic("implement me")
}
