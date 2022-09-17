package usecase

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type RoleRepository interface {
	Create(context.Context, entity.UserRole) (int, error)
	GetAll(context.Context) ([]entity.UserRole, error)
	GetById(context.Context, int) (*entity.UserRole, error)
	Update(context.Context, int, entity.UserRole) error
	Delete(context.Context, int) error
}

type RoleUseCase struct {
	repo RoleRepository
}

func NewRoleUseCase(r RoleRepository) *RoleUseCase {
	return &RoleUseCase{
		repo: r,
	}
}

func (uc *RoleUseCase) Create(ctx context.Context, data dto.CreateRoleRequest) (*dto.CreateRoleResponse, error) {
	id, err := uc.repo.Create(ctx, entity.UserRole{
		Name: data.Name,
	})

	if err != nil {
		return nil, fmt.Errorf("RoleUseCase - Create - uc.repo.Create: %w", err)
	}

	return &dto.CreateRoleResponse{Id: id}, nil
}

func (uc *RoleUseCase) GetAll(ctx context.Context) (*dto.GetAllRolesResponse, error) {
	all, err := uc.repo.GetAll(ctx)

	if err != nil {
		return nil, fmt.Errorf("RoleUseCase - GetAll - uc.repo.GetAll: %w", err)
	}

	return &dto.GetAllRolesResponse{Roles: all}, nil
}

func (uc *RoleUseCase) GetById(ctx context.Context, data dto.GetRoleByIdRequest) (*dto.GetRoleByIdResponse, error) {
	role, err := uc.repo.GetById(ctx, data.Id)

	if err != nil {
		return nil, fmt.Errorf("RoleUseCase - GetById - uc.repo.GetById: %w", err)
	}

	return &dto.GetRoleByIdResponse{
		Id:   role.Id,
		Name: role.Name,
	}, nil
}

func (uc *RoleUseCase) Update(ctx context.Context, data dto.UpdateRoleRequest) error {
	err := uc.repo.Update(ctx, data.Id, entity.UserRole{
		Id:   data.Id,
		Name: data.Name})

	if err != nil {
		return fmt.Errorf("RoleUseCase - Update - uc.repo.Update: %w", err)
	}

	return nil
}

func (uc *RoleUseCase) Delete(ctx context.Context, data dto.DeleteRoleRequest) error {
	err := uc.repo.Delete(ctx, data.Id)

	if err != nil {
		return fmt.Errorf("RoleUseCase - Delete - uc.repo.Delete: %w", err)
	}

	return nil
}
