package user

import (
	"fmt"
	"time"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Role      string // add enum for role = (admin, user, guest)
	CreatedAt time.Time
}

func (u User) String() string {
	return fmt.Sprintf("User: { ID: %d, Name: %s, Email: %s, Role: %s, CreatedAt: %s}", u.ID, u.Name, u.Email, u.Role, u.CreatedAt.Format(time.RFC3339))
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
