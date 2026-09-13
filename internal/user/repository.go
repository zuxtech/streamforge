package user

import "context"

type Repository interface {
	Create(ctx context.Context, u *User) error
	GetByID(ctx context.Context, id string) (*User, error)
	GetByKratosIdentityID(ctx context.Context, kratosIdentityID string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}
