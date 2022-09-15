package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/ekaterinamzr/green-alarm/internal/dto"
	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type IncidentRepository interface {
	Create(context.Context, entity.Incident) (int, error)
	GetAll(context.Context) ([]entity.Incident, error)
	GetById(context.Context, int) (*entity.Incident, error)
	Update(context.Context, int, entity.Incident) error
	Delete(context.Context, int) error
}

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
		Name:             data.Name,
		Date:             data.Date,
		Country:          data.Country,
		Latitude:         data.Latitude,
		Longitude:        data.Longitude,
		Publication_date: time.Now(),
		Comment:          data.Comment,
		Status:           entity.Unconfirmed,
		Type:             data.Type,
		Author:           data.Author,
	})

	if err != nil {
		return nil, fmt.Errorf("IncidentUseCase - Create - uc.repo.Create: %w", err)
	}

	return &dto.CreateIncidentResponse{Id: id}, nil
}

func (uc *IncidentUseCase) GetAll(ctx context.Context) (*dto.GetAllIncidentsResponse, error) {
	all, err := uc.repo.GetAll(ctx)

	if err != nil {
		return nil, fmt.Errorf("IncidentUseCase - GetAll - uc.repo.GetAll: %w", err)
	}

	return &dto.GetAllIncidentsResponse{Incidents: all}, nil
}

func (uc *IncidentUseCase) GetById(ctx context.Context, data dto.GetIncidentByIdRequest) (*dto.GetIncidentByIdResponse, error) {
	incident, err := uc.repo.GetById(ctx, data.Id)

	if err != nil {
		return nil, fmt.Errorf("IncidentUseCase - GetById - uc.repo.GetById: %w", err)
	}

	return &dto.GetIncidentByIdResponse{
		Id:               incident.Id,
		Name:             incident.Name,
		Date:             incident.Date,
		Country:          incident.Country,
		Latitude:         incident.Latitude,
		Longitude:        incident.Longitude,
		Publication_date: incident.Publication_date,
		Comment:          incident.Comment,
		Status:           entity.Unconfirmed,
		Type:             incident.Type,
		Author:           incident.Author}, nil
}

func (uc *IncidentUseCase) Update(ctx context.Context, data dto.UpdateIncidentRequest) error {
	err := uc.repo.Update(ctx, data.Id, entity.Incident{
		Id:        data.Id,
		Name:      data.Name,
		Date:      data.Date,
		Country:   data.Country,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		Comment:   data.Comment,
		Status:    entity.Unconfirmed,
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
