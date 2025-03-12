package user

import (
	"fmt"
	"time"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Role      string
	CreatedAt time.Time
}

func (u User) String() string {
	return fmt.Sprintf("User: { ID: %d, Name: %s, Email: %s, Role: %s, CreatedAt: %s}", u.ID, u.Name, u.Email, u.Role, u.CreatedAt.Format(time.RFC3339))
}
