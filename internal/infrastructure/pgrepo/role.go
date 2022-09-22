package pgrepo

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/entity"
	"github.com/ekaterinamzr/green-alarm/pkg/postgres"
)

type RoleRepository struct {
	*postgres.Postgres
}

func NewRoleRepository(pg *postgres.Postgres) *RoleRepository {
	return &RoleRepository{pg}
}

func (r *RoleRepository) Create(ctx context.Context, t entity.UserRole) (string, error) {
	var id string
	query := `	INSERT INTO 
					roles (
						role_name) 
				VALUES 
					($1) 
				RETURNING 
					id`

	row := r.DB.QueryRowContext(ctx, query, t.Name)
	err := row.Scan(&id)
	if err != nil {
		return "", fmt.Errorf("pgrepo - role - CreateRole: %w", err)
	}
	return id, nil
}

func (r *RoleRepository) GetAll(ctx context.Context) ([]entity.UserRole, error) {
	var all []entity.UserRole

	query := `	SELECT 
					id, 
					role_name, 
				FROM 
					roles`

	if err := r.DB.SelectContext(ctx, &all, query); err != nil {
		return nil, fmt.Errorf("pgrepo - role - GetAll: %w", err)
	}

	return all, nil
}

func (r *RoleRepository) GetById(ctx context.Context, id string) (*entity.UserRole, error) {
	var role entity.UserRole

	query := `	SELECT 
					id, 
					role_name
				FROM 
					roles 
				WHERE 
					id = $1`

	row := r.DB.QueryRowxContext(ctx, query, id)

	if err := row.StructScan(&role); err != nil {
		return nil, fmt.Errorf("pgrepo - role - GetById: %w", err)
	}
	return &role, nil
}

func (r *RoleRepository) Update(ctx context.Context, id string, updated entity.UserRole) error {
	query := `	UPDATE 
					roles 
				SET 
					role_name = $1
				WHERE 
					id = $2`

	_, err := r.DB.ExecContext(ctx, query, updated.Name, id)

	if err != nil {
		return fmt.Errorf("pgrepo - role - Update: %w", err)
	}

	return nil
}

func (r *RoleRepository) Delete(ctx context.Context, id string) error {
	query := `	DELETE 
				FROM 
					roles 
				WHERE 
					id = $1`

	_, err := r.DB.ExecContext(ctx, query, id)

	if err != nil {
		return fmt.Errorf("pgrepo - role - Delete: %w", err)
	}

	return nil
}
