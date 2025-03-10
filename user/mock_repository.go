package user

import "fmt"

type MockRepository struct{}

func (m *MockRepository) Save(user *User) error {
	fmt.Println("Call Save function")
	return nil
}

func (m *MockRepository) FindByID(id int) (User, error) {
	fmt.Println("Call FindByID function")
	return User{}, nil
}

func (m *MockRepository) FindAll() []User {
	fmt.Println("Call FindAll function")
	return []User{}
}

func (m *MockRepository) DeleteByID(id int) error {
	fmt.Println("Call DeleteByID function")
	return nil
}
