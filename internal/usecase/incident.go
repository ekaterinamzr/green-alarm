package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type IncidentUseCase struct {
	repo IncidentRepository
}

func NewIncidentUseCase(r IncidentRepository) *IncidentUseCase {
	return &IncidentUseCase{
		repo: r,
	}
}

func (uc *IncidentUseCase) Create(ctx context.Context, data dto.CreateIncidentRequest) (*dto.CreateIncidentResponse, error) {
	id, err := uc.repo.Create(ctx, entity.Incident{
		Name:        data.Name,
		Date:        data.Date,
		Country:     data.Country,
		Latitude:    data.Latitude,
		Longitude:   data.Longitude,
		Publication: time.Now(),
		Comment:     data.Comment,
		Status:      entity.Unconfirmed,
		Type:        data.Type,
		Author:      data.Author,
	})

	if err != nil {
		return nil, fmt.Errorf("IncidentUseCase - Create - uc.repo.Create: %w", err)
	}

	return &dto.CreateIncidentResponse{Id: id}, nil
}

func (uc *IncidentUseCase) GetAll(ctx context.Context) (*dto.GetIncidentsResponse, error) {
	all, err := uc.repo.GetAll(ctx)

	if err != nil {
		return nil, fmt.Errorf("IncidentUseCase - GetAll - uc.repo.GetAll: %w", err)
	}

	res := dto.GetIncidentsResponse(dto.FromIncidents(all))
	return &res, nil
}

func (uc *IncidentUseCase) GetByType(ctx context.Context, data dto.GetIncidentsByTypeRequest) (*dto.GetIncidentsResponse, error) {
	incidents, err := uc.repo.GetByType(ctx, data.IncidentType)

	if err != nil {
		return nil, fmt.Errorf("IncidentUseCase - GetByType - uc.repo.GetByType: %w", err)
	}

	res := dto.GetIncidentsResponse(dto.FromIncidents(incidents))
	return &res, nil
}

func (uc *IncidentUseCase) GetById(ctx context.Context, data dto.GetIncidentByIdRequest) (*dto.GetIncidentByIdResponse, error) {
	incident, err := uc.repo.GetById(ctx, data.Id)

	if err != nil {
		return nil, fmt.Errorf("IncidentUseCase - GetById - uc.repo.GetById: %w", err)
	}

	res := dto.GetIncidentByIdResponse(dto.FromIncident(incident))
	return &res, nil
}

func (uc *IncidentUseCase) Update(ctx context.Context, data dto.UpdateIncidentRequest) error {
	err := uc.repo.Update(ctx, entity.Incident{
		Id:        data.Id,
		Name:      data.Name,
		Date:      data.Date,
		Country:   data.Country,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		Comment:   data.Comment,
		Status:    data.Status,
		Type:      data.Type,
		Author:    data.Author})

	if err != nil {
		return fmt.Errorf("IncidentUseCase - Update - uc.repo.Update: %w", err)
	}

	return nil
}

func (uc *IncidentUseCase) Delete(ctx context.Context, data dto.DeleteIncidentRequest) error {
	err := uc.repo.Delete(ctx, data.Id)

	if err != nil {
		return fmt.Errorf("IncidentUseCase - Delete - uc.repo.Delete: %w", err)
	}

	return nil
}
