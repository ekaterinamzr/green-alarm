package pgrepo

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/entity"
	"github.com/ekaterinamzr/green-alarm/pkg/postgres"
)

type StatusRepository struct {
	*postgres.Postgres
}

func NewStatusRepository(pg *postgres.Postgres) *StatusRepository {
	return &StatusRepository{pg}
}

func (r *StatusRepository) Create(ctx context.Context, t entity.IncidentStatus) (string, error) {
	var id string
	query := `	INSERT INTO 
					statuses (
						status_name) 
				VALUES 
					($1) 
				RETURNING 
					id`

	row := r.DB.QueryRowContext(ctx, query, t.Name)
	err := row.Scan(&id)
	if err != nil {
		return "", fmt.Errorf("pgrepo - status - CreateStatus: %w", err)
	}
	return id, nil
}

func (r *StatusRepository) GetAll(ctx context.Context) ([]entity.IncidentStatus, error) {
	var all []entity.IncidentStatus

	query := `	SELECT 
					id, 
					status_name, 
				FROM 
					statuses`
					
	if err := r.DB.SelectContext(ctx, &all, query); err != nil {
		return nil, fmt.Errorf("pgrepo - status - GetAll: %w", err)
	}

	return all, nil
}

func (r *StatusRepository) GetById(ctx context.Context, id string) (*entity.IncidentStatus, error) {
	var s entity.IncidentStatus

	query := `	SELECT 
					id, 
					status_name
				FROM 
					statuses 
				WHERE 
					id = $1`

	row := r.DB.QueryRowxContext(ctx, query, id)

	if err := row.StructScan(&s); err != nil {
		return nil, fmt.Errorf("pgrepo - status - GetById: %w", err)
	}
	return &s, nil
}

func (r *StatusRepository) Update(ctx context.Context, id string, updated entity.IncidentStatus) error {
	query := `	UPDATE 
					statuses 
				SET 
					status_name = $1
				WHERE 
					id = $2`

	_, err := r.DB.ExecContext(ctx, query, updated.Name, id)

	if err != nil {
		return fmt.Errorf("pgrepo - status - Update: %w", err)
	}

	return nil
}

func (r *StatusRepository) Delete(ctx context.Context, id string) error {
	query := `	DELETE 
				FROM 
					statuses 
				WHERE 
					id = $1`

	_, err := r.DB.ExecContext(ctx, query, id)

	if err != nil {
		return fmt.Errorf("pgrepo - status - Delete: %w", err)
	}

	return nil
}
