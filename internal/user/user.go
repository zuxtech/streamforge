package user

import "time"

type User struct {
	ID               string
	KratosIdentityID string
	Email            string
	FirstName        string
	LastName         string
	EmailVerified    bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
