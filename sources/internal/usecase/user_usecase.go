package usecase

import (
	"context"
	"go-clean-arch/internal/domain"
)

type userUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(r domain.UserRepository) domain.UserUseCase {
	return &userUseCase{repo: r}
}

func (u *userUseCase) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	return u.repo.GetByID(ctx, id)
}
