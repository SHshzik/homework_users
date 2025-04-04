package user

type Repository interface {
	Save(user *User) error
	FindByID(id int) (User, error)
	FindAll() []User
	DeleteByID(id int) error
	FindByRole(role string) []User
}
