package usecase

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type TypeUseCase struct {
	repo TypeRepository
}

func NewTypeUseCase(r TypeRepository) *TypeUseCase {
	return &TypeUseCase{
		repo: r,
	}
}

func (uc *TypeUseCase) Create(ctx context.Context, data dto.CreateTypeRequest) (*dto.CreateTypeResponse, error) {
	id, err := uc.repo.Create(ctx, entity.IncidentType{
		Name: data.Name,
	})

	if err != nil {
		return nil, fmt.Errorf("TypeUseCase - Create - uc.repo.Create: %w", err)
	}

	return &dto.CreateTypeResponse{Id: id}, nil
}

func (uc *TypeUseCase) GetAll(ctx context.Context) (*dto.GetAllTypesResponse, error) {
	all, err := uc.repo.GetAll(ctx)

	if err != nil {
		return nil, fmt.Errorf("TypeUseCase - GetAll - uc.repo.GetAll: %w", err)
	}

	res := dto.GetAllTypesResponse(dto.FromTypes(all))
	return &res, nil
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
