package user

import "time"

type Service struct {
	repo Repository
}

func (uService *Service) CreateUser(name, email, role string) (*User, error) {
	nUser := &User{
		Name:      name,
		Email:     email,
		Role:      role,
		CreatedAt: time.Now(),
	}
	return nUser, uService.repo.Save(nUser)
}

func (uService *Service) GetUser(id int) (User, error) {
	return uService.repo.FindByID(id)
}

func (uService *Service) ListUsers() []User {
	return uService.repo.FindAll()
}

func (uService *Service) RemoveUser(id int) error {
	return uService.repo.DeleteByID(id)
}

func NewService(repo Repository) Service {
	return Service{repo: repo}
}
