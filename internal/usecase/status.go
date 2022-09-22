package usecase

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type StatusRepository interface {
	Create(context.Context, entity.IncidentStatus) (string, error)
	GetAll(context.Context) ([]entity.IncidentStatus, error)
	GetById(context.Context, string) (*entity.IncidentStatus, error)
	Update(context.Context, string, entity.IncidentStatus) error
	Delete(context.Context, string) error
}

type StatusUseCase struct {
	repo StatusRepository
}

func NewStatusUseCase(r StatusRepository) *StatusUseCase {
	return &StatusUseCase{
		repo: r,
	}
}

func (uc *StatusUseCase) Create(ctx context.Context, data dto.CreateStatusRequest) (*dto.CreateStatusResponse, error) {
	id, err := uc.repo.Create(ctx, entity.IncidentStatus{
		Name: data.Name,
	})

	if err != nil {
		return nil, fmt.Errorf("StatusUseCase - Create - uc.repo.Create: %w", err)
	}

	return &dto.CreateStatusResponse{Id: id}, nil
}

func (uc *StatusUseCase) GetAll(ctx context.Context) (*dto.GetAllStatusesResponse, error) {
	all, err := uc.repo.GetAll(ctx)

	if err != nil {
		return nil, fmt.Errorf("StatusUseCase - GetAll - uc.repo.GetAll: %w", err)
	}

	return &dto.GetAllStatusesResponse{Statuss: all}, nil
}

func (uc *StatusUseCase) GetById(ctx context.Context, data dto.GetStatusByIdRequest) (*dto.GetStatusByIdResponse, error) {
	status, err := uc.repo.GetById(ctx, data.Id)

	if err != nil {
		return nil, fmt.Errorf("StatusUseCase - GetById - uc.repo.GetById: %w", err)
	}

	return &dto.GetStatusByIdResponse{
		Id:   status.Id,
		Name: status.Name,
	}, nil
}

func (uc *StatusUseCase) Update(ctx context.Context, data dto.UpdateStatusRequest) error {
	err := uc.repo.Update(ctx, data.Id, entity.IncidentStatus{
		Id:   data.Id,
		Name: data.Name})

	if err != nil {
		return fmt.Errorf("StatusUseCase - Update - uc.repo.Update: %w", err)
	}

	return nil
}

func (uc *StatusUseCase) Delete(ctx context.Context, data dto.DeleteStatusRequest) error {
	err := uc.repo.Delete(ctx, data.Id)

	if err != nil {
		return fmt.Errorf("StatusUseCase - Delete - uc.repo.Delete: %w", err)
	}

	return nil
}
