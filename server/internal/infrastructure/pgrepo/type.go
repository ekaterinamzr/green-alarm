package pgrepo

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/entity"
	"github.com/ekaterinamzr/green-alarm/pkg/postgres"
)

type TypeRepository struct {
	*postgres.Postgres
}

func NewTypeRepository(pg *postgres.Postgres) *TypeRepository {
	return &TypeRepository{pg}
}

func (r *TypeRepository) Create(ctx context.Context, t entity.IncidentType) (int, error) {
	var id int
	query := `	INSERT INTO 
					types (
						type_name) 
				VALUES 
					($1) 
				RETURNING 
					id`

	row := r.DB.QueryRowContext(ctx, query, t.Name)
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("pgrepo - type - CreateType: %w", err)
	}
	return id, nil
}

func (r *TypeRepository) GetAll(ctx context.Context) ([]entity.IncidentType, error) {
	var all []entity.IncidentType

	query := `	SELECT 
					id, 
					type_name 
				FROM 
					types`

	if err := r.DB.SelectContext(ctx, &all, query); err != nil {
		return nil, fmt.Errorf("pgrepo - type - GetAll: %w", err)
	}

	return all, nil
}

func (r *TypeRepository) GetById(ctx context.Context, id int) (*entity.IncidentType, error) {
	var t entity.IncidentType

	query := `	SELECT 
					id, 
					type_name
				FROM 
					types 
				WHERE 
					id = $1`

	row := r.DB.QueryRowxContext(ctx, query, id)

	if err := row.StructScan(&t); err != nil {
		return nil, fmt.Errorf("pgrepo - type - GetById: %w", err)
	}
	return &t, nil
}

func (r *TypeRepository) Update(ctx context.Context, id int, updated entity.IncidentType) error {
	query := `	UPDATE 
					types 
				SET 
					type_name = $1
				WHERE 
					id = $2`

	_, err := r.DB.ExecContext(ctx, query, updated.Name, id)

	if err != nil {
		return fmt.Errorf("pgrepo - type - Update: %w", err)
	}

	return nil
}

func (r *TypeRepository) Delete(ctx context.Context, id int) error {
	query := `	DELETE 
				FROM 
					types 
				WHERE 
					id = $1`

	_, err := r.DB.ExecContext(ctx, query, id)

	if err != nil {
		return fmt.Errorf("pgrepo - type - Delete: %w", err)
	}

	return nil
}
