package pgrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/entity"
	"github.com/ekaterinamzr/green-alarm/pkg/postgres"
)

type UserRepository struct {
	*postgres.Postgres
}

func NewUserRepository(pg *postgres.Postgres) *UserRepository {
	return &UserRepository{pg}
}

func (r *UserRepository) CreateUser(ctx context.Context, u entity.User) (int, error) {
	var id int
	query := "INSERT INTO users (first_name, last_name, username, email, user_password, user_role) values ($1, $2, $3, $4, $5, $6) RETURNING id"
	row := r.DB.QueryRowContext(ctx, query, u.First_name, u.Last_name, u.Username, u.Email, u.Password, 3)
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("pgrepo - user - CreateUser: %w", err)
	}
	return id, nil
}

func (r *UserRepository) GetUser(ctx context.Context, username, password string) (*entity.User, error) {
	var user entity.User

	query := "SELECT * FROM users WHERE username = $1 AND user_password = $2"
	row := r.DB.QueryRowxContext(ctx, query, username, password)

	if err := row.StructScan(&user); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("User not found: %w", err)
		}
		return nil, fmt.Errorf("pgrepo - user - GetUser: %w", err)
	}
	return &user, nil

}
