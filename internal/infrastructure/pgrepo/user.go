package pgrepo

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/entity"
	"github.com/ekaterinamzr/green-alarm/pkg/postgres"
)

type UserRepository struct {
	pg *postgres.Postgres
}

func NewUserRepository(pg *postgres.Postgres) *UserRepository {
	return &UserRepository{pg}
}

func (r *UserRepository) CreateUser(ctx context.Context, u entity.User) (int, error) {
	var id int
	query := "INSERT INTO users (first_name, last_name, username, email, user_password, user_role) values ($1, $2, $3, $4, $5, $6) RETURNING id"
	row := r.pg.DB.QueryRowContext(ctx, query, u.First_name, u.Last_name, u.Username, u.Email, u.Password, 3)
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("pgrepo - user - CreateUser: %w", err)
	}
	return id, nil
}

func (r *UserRepository) GetById(id int) (*entity.User, error) {
	return &entity.User{}, nil
}
