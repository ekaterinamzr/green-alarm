package usecase

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

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

	res := dto.GetAllUsersResponse(dto.FromUsers(all))
	return &res, nil
}

func (uc *UserUseCase) GetById(ctx context.Context, data dto.GetUserByIdRequest) (*dto.GetUserByIdResponse, error) {
	user, err := uc.repo.GetById(ctx, data.Id)

	if err != nil {
		return nil, fmt.Errorf("UserUseCase - GetById - uc.repo.GetById: %w", err)
	}

	return &dto.GetUserByIdResponse{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
	}, nil
}

func (uc *UserUseCase) Update(ctx context.Context, data dto.UpdateUserRequest) error {
	err := uc.repo.Update(ctx, data.Id, entity.User{
		Id:        data.Id,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Username:  data.Username,
		Email:     data.Email,
		Password:  data.Password,
		Role:      data.Role})

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
