package usecase

import (
	"context"

	"github.com/ekaterinamzr/green-alarm/internal/entity"
)

type (
	IncidentRepository interface {
		Create(context.Context, entity.Incident) (int, error)
		GetAll(context.Context) ([]entity.Incident, error)
		GetById(context.Context, int) (*entity.Incident, error)
		Update(context.Context, entity.Incident) error
		Delete(context.Context, int) error
		GetByType(context.Context, int) ([]entity.Incident, error)
	}

	AuthRepository interface {
		Create(context.Context, entity.User) (int, error)
		GetUser(context.Context, string, string) (*entity.User, error)
	}

	UserRepository interface {
		GetAll(context.Context) ([]entity.User, error)
		GetById(context.Context, int) (*entity.User, error)
		Update(context.Context, int, entity.User) error
		ChangeRole(context.Context, int, int) error
		Delete(context.Context, int) error
	}

	RoleRepository interface {
		Create(context.Context, entity.UserRole) (int, error)
		GetAll(context.Context) ([]entity.UserRole, error)
		GetById(context.Context, int) (*entity.UserRole, error)
		Update(context.Context, int, entity.UserRole) error
		Delete(context.Context, int) error
	}

	StatusRepository interface {
		Create(context.Context, entity.IncidentStatus) (int, error)
		GetAll(context.Context) ([]entity.IncidentStatus, error)
		GetById(context.Context, int) (*entity.IncidentStatus, error)
		Update(context.Context, int, entity.IncidentStatus) error
		Delete(context.Context, int) error
	}

	TypeRepository interface {
		Create(context.Context, entity.IncidentType) (int, error)
		GetAll(context.Context) ([]entity.IncidentType, error)
		GetById(context.Context, int) (*entity.IncidentType, error)
		Update(context.Context, int, entity.IncidentType) error
		Delete(context.Context, int) error
	}

	Tokenizer interface {
		GenerateToken(ctx context.Context, id, role int) (string, error)
		ParseToken(context.Context, string) (id int, role int, err error)
	}
)
