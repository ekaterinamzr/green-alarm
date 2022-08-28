package usecase

import (
	"context"

	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type User interface {
	Register(context.Context, entity.User) error
}

type UserRepository interface {
}

type UserUseCase struct {
	repo UserRepository
}

func NewUserUseCase(r UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (uc *UserUseCase) Register(ctx context.Context, u entity.User) error {

	return nil
}
