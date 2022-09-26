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

func (r *UserRepository) Create(ctx context.Context, u entity.User) (int, error) {
	var id int
	query := `	INSERT 
				INTO 
					users (
						first_name, 
						last_name, 
						username, 
						email, 
						user_password, 
						user_role) 
				VALUES 
					($1, $2, $3, $4, $5, $6) 
				RETURNING 
					id`

	row := r.DB.QueryRowContext(ctx, query, u.FirstName, u.LastName, u.Username, u.Email, u.Password, u.Role)
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("pgrepo - user - CreateUser: %w", err)
	}
	return id, nil
}

func (r *UserRepository) GetUser(ctx context.Context, username, password string) (*entity.User, error) {
	var user entity.User

	query := `	SELECT * 
				FROM 
					users 
				WHERE 
					username = $1 AND 
					user_password = $2`

	row := r.DB.QueryRowxContext(ctx, query, username, password)

	if err := row.StructScan(&user); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("User not found: %w", err)
		}
		return nil, fmt.Errorf("pgrepo - user - GetUser: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	var all []entity.User

	query := `	SELECT 
					id,
					first_name,
					last_name,
					username,
					email,
					user_password,
					user_role 
				FROM
					users`

	if err := r.DB.SelectContext(ctx, &all, query); err != nil {
		return nil, fmt.Errorf("pgrepo - user - GetAll: %w", err)
	}

	return all, nil
}

func (r *UserRepository) GetById(ctx context.Context, id int) (*entity.User, error) {
	var user entity.User

	query := `	SELECT 
					id,
					first_name,
					last_name,
					username,
					email,
					user_password,
					user_role 
				FROM 
					users 
				WHERE 
					id = $1`

	row := r.DB.QueryRowxContext(ctx, query, id)

	if err := row.StructScan(&user); err != nil {
		return nil, fmt.Errorf("pgrepo - user - GetById: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, id int, updated entity.User) error {
	query := `	UPDATE 
					users 
				SET 
					first_name = $1,
					last_name = $2,
					username = $3,
					email = $4,
					user_password = $5,
					user_role = $6 
				WHERE 
					id = $7`

	_, err := r.DB.ExecContext(ctx, query,
		updated.FirstName,
		updated.LastName,
		updated.Username,
		updated.Email,
		updated.Password,
		updated.Role,
		id,
	)

	if err != nil {
		return fmt.Errorf("pgrepo - user - Update: %w", err)
	}

	return nil
}

func (r *UserRepository) ChangeRole(ctx context.Context, id, newRole int) error {
	query := `	UPDATE 
					users 
				SET 
					user_role = $1
				WHERE 
					id = $2`

	_, err := r.DB.ExecContext(ctx, query,
		newRole,
		id,
	)

	if err != nil {
		return fmt.Errorf("pgrepo - user - UpdateRole: %w", err)
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	query := `	DELETE 
				FROM 
					users 
				WHERE 
					id = $1`

	_, err := r.DB.ExecContext(ctx, query, id)

	if err != nil {
		return fmt.Errorf("pgrepo - user - Delete: %w", err)
	}

	return nil
}
