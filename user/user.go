package user

import (
	"time"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Role      string // add enum for role = (admin, user, guest)
	CreatedAt time.Time
}

type Repository interface {
	Save(user *User) error
	FindByID(id int) (User, error)
	FindAll() []User
	DeleteByID(id int) error
}

func NewService(repo Repository) Service {
	return Service{repo: repo}
}
