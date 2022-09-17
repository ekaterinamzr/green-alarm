package usecase

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type UserRepository interface {
	GetAll(context.Context) ([]entity.User, error)
	GetById(context.Context, int) (*entity.User, error)
	Update(context.Context, int, entity.User) error
	ChangeRole(context.Context, int, int) error
	Delete(context.Context, int) error
}

type UserUseCase struct {
	repo UserRepository
}

func NewUserUseCase(r UserRepository) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (uc *UserUseCase) GetAll(ctx context.Context) (*dto.GetAllUsersResponse, error) {
	all, err := uc.repo.GetAll(ctx)

	if err != nil {
		return nil, fmt.Errorf("UserUseCase - GetAll - uc.repo.GetAll: %w", err)
	}

	return &dto.GetAllUsersResponse{Users: all}, nil
}

func (uc *UserUseCase) GetById(ctx context.Context, data dto.GetUserByIdRequest) (*dto.GetUserByIdResponse, error) {
	user, err := uc.repo.GetById(ctx, data.Id)

	if err != nil {
		return nil, fmt.Errorf("UserUseCase - GetById - uc.repo.GetById: %w", err)
	}

	return &dto.GetUserByIdResponse{
		Id:         user.Id,
		First_name: user.First_name,
		Last_name:  user.Last_name,
		Username:   user.Username,
		Email:      user.Email,
		Password:   user.Password,
		Role:       user.Role,
	}, nil
}

func (uc *UserUseCase) Update(ctx context.Context, data dto.UpdateUserRequest) error {
	err := uc.repo.Update(ctx, data.Id, entity.User{
		Id:         data.Id,
		First_name: data.First_name,
		Last_name:  data.Last_name,
		Username:   data.Username,
		Email:      data.Email,
		Password:   data.Password,
		Role:       data.Role})

	if err != nil {
		return fmt.Errorf("UserUseCase - Update - uc.repo.Update: %w", err)
	}

	return nil
}

func (uc *UserUseCase) ChangeRole(ctx context.Context, data dto.ChangeRoleRequest) error {
	err := uc.repo.ChangeRole(ctx, data.Id, data.NewRole)

	if err != nil {
		return fmt.Errorf("UserUseCase - ChangeRole - uc.repo.Update: %w", err)
	}

	return nil
}

func (uc *UserUseCase) Delete(ctx context.Context, data dto.DeleteUserRequest) error {
	err := uc.repo.Delete(ctx, data.Id)

	if err != nil {
		return fmt.Errorf("UserUseCase - Delete - uc.repo.Delete: %w", err)
	}

	return nil
}
