package pgrepo

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/entity"
	"github.com/ekaterinamzr/green-alarm/pkg/postgres"
)

type IncidentRepository struct {
	*postgres.Postgres
}

func NewIncidentRepository(pg *postgres.Postgres) *IncidentRepository {
	return &IncidentRepository{pg}
}

func (r *IncidentRepository) Create(ctx context.Context, i entity.Incident) (int, error) {
	var id int
	query := `	INSERT INTO 
					incidents (
						incident_name, 
						incident_date, 
						country, 
						latitude, 
						longitude, 
						publication_date, 
						comment, 
						incident_status, 
						incident_type, 
						author) 
				VALUES 
					($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
				RETURNING 
					id`

	row := r.DB.QueryRowContext(ctx, query, i.Name, i.Date, i.Country, i.Latitude, i.Longitude, i.Publication, i.Comment, i.Status, i.Type, i.Author)
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("pgrepo - incident - CreateIncident: %w", err)
	}
	return id, nil
}

func (r *IncidentRepository) GetAll(ctx context.Context) ([]entity.Incident, error) {
	var all []entity.Incident

	query := `	SELECT 
					id, 
					incident_name, 
					incident_date, 
					country, 
					latitude, 
					longitude, 
					publication_date, 
					COALESCE(comment, '') AS comment, 
					incident_status, 
					incident_type, 
					author 
				FROM 
					incidents`

	if err := r.DB.SelectContext(ctx, &all, query); err != nil {
		return nil, fmt.Errorf("pgrepo - incident - GetAll: %w", err)
	}

	return all, nil
}

func (r *IncidentRepository) GetById(ctx context.Context, id int) (*entity.Incident, error) {
	var incident entity.Incident

	query := `	SELECT 
					id, 
					incident_name, 
					incident_date, 
					country, 
					latitude, 
					longitude, 
					publication_date, 
					COALESCE(comment, '') AS comment, 
					incident_status, 
					incident_type, 
					author
				FROM 
					incidents 
				WHERE 
					id = $1`

	row := r.DB.QueryRowxContext(ctx, query, id)

	if err := row.StructScan(&incident); err != nil {
		return nil, fmt.Errorf("pgrepo - incident - GetById: %w", err)
	}
	return &incident, nil
}

func (r *IncidentRepository) GetByType(ctx context.Context, requiredType int) ([]entity.Incident, error) {
	var incidents []entity.Incident

	query := `	SELECT 
					id, 
					incident_name, 
					incident_date, 
					country, 
					latitude, 
					longitude, 
					publication_date, 
					COALESCE(comment, '') AS comment, 
					incident_status, 
					incident_type, 
					author 
				FROM 
					incidents
				WHERE 
					incident_type = $1`

	if err := r.DB.SelectContext(ctx, &incidents, query, requiredType); err != nil {
		return nil, fmt.Errorf("pgrepo - incident - GetByType: %w", err)
	}

	return incidents, nil
}

func (r *IncidentRepository) Update(ctx context.Context, updated entity.Incident) error {
	query := `	UPDATE 
					incidents 
				SET 
					incident_name = $1,
					incident_date = $2,
					country = $3,
					latitude = $4,
					longitude = $5,
					comment = $6,
					incident_status = $7,
					incident_type = $8,
					author = $9
				WHERE 
					id = $10`

	_, err := r.DB.ExecContext(ctx, query,
		updated.Name,
		updated.Date,
		updated.Country,
		updated.Latitude,
		updated.Longitude,
		updated.Comment,
		updated.Status,
		updated.Type,
		updated.Author,
		updated.Id,
	)

	if err != nil {
		return fmt.Errorf("pgrepo - incident - Update: %w", err)
	}

	return nil
}

func (r *IncidentRepository) Delete(ctx context.Context, id int) error {
	query := `	DELETE 
				FROM 
					incidents 
				WHERE 
					id = $1`

	_, err := r.DB.ExecContext(ctx, query, id)

	if err != nil {
		return fmt.Errorf("pgrepo - incident - Delete: %w", err)
	}

	return nil
}
