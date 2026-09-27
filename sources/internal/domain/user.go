package domain

import (
	"context"
	"errors"
)

var ErrUserNotFound = errors.New("usuario no encontrado")

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (*User, error)
}

type UserUseCase interface {
	GetUser(ctx context.Context, id int64) (*User, error)
}
