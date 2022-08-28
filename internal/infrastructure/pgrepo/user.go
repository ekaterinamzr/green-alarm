package pgrepo

import (
	"github.com/ekaterinamzr/green-alarm/internal/entity"
	"github.com/ekaterinamzr/green-alarm/pkg/postgres"
)

type UserRepository struct {
	pg *postgres.Postgres
}

func NewUserRepository(pg *postgres.Postgres) *UserRepository {
	return &UserRepository{pg}
}

func (r *UserRepository) GetById(id int) (*entity.User, error) {
	return &entity.User{}, nil
}
