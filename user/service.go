package user

type Service struct {
	repo Repository
}

func (uService *Service) CreateUser(name, email, role string) (User, error) {
	panic("implement me")
}

func (uService *Service) GetUser(id int) (User, error) {
	panic("implement me")
}

func (uService *Service) ListUsers() []User {
	return uService.repo.FindAll()
}
