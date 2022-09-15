package usecase

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type TypeRepository interface {
	GetAll(context.Context) ([]entity.IncidentType, error)
	GetById(context.Context, int) (*entity.IncidentType, error)
	Update(context.Context, int, entity.IncidentType) error
	Delete(context.Context, int) error
}

type TypeUseCase struct {
	repo TypeRepository
}

func NewTypeUseCase(r TypeRepository) *TypeUseCase {
	return &TypeUseCase{
		repo: r,
	}
}

func (uc *TypeUseCase) GetAll(ctx context.Context) (*dto.GetAllTypesResponse, error) {
	all, err := uc.repo.GetAll(ctx)

	if err != nil {
		return nil, fmt.Errorf("TypeUseCase - GetAll - uc.repo.GetAll: %w", err)
	}

	return &dto.GetAllTypesResponse{Types: all}, nil
}

func (uc *TypeUseCase) GetById(ctx context.Context, data dto.GetTypeByIdRequest) (*dto.GetTypeByIdResponse, error) {
	t, err := uc.repo.GetById(ctx, data.Id)

	if err != nil {
		return nil, fmt.Errorf("TypeUseCase - GetById - uc.repo.GetById: %w", err)
	}

	return &dto.GetTypeByIdResponse{
		Id:   t.Id,
		Name: t.Name,
	}, nil
}

func (uc *TypeUseCase) Update(ctx context.Context, data dto.UpdateTypeRequest) error {
	err := uc.repo.Update(ctx, data.Id, entity.IncidentType{
		Id:   data.Id,
		Name: data.Name,})

	if err != nil {
		return fmt.Errorf("TypeUseCase - Update - uc.repo.Update: %w", err)
	}

	return nil
}

func (uc *TypeUseCase) Delete(ctx context.Context, data dto.DeleteTypeRequest) error {
	err := uc.repo.Delete(ctx, data.Id)

	if err != nil {
		return fmt.Errorf("TypeUseCase - Delete - uc.repo.Delete: %w", err)
	}

	return nil
}
