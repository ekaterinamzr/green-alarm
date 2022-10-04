package integrationtests

import (
	"context"
	"testing"
	"time"

	"github.com/ekaterinamzr/green-alarm/internal/entity"
	"github.com/ekaterinamzr/green-alarm/internal/infrastructure/pgrepo"
	"github.com/ekaterinamzr/green-alarm/pkg/postgres"
	"github.com/stretchr/testify/require"
)

const uri = "postgres://greenalarm:greenalarm@postgresc:5432/greenalarm"

func TestCreateDeleteIncident(t *testing.T) {
	pg, err := postgres.New(uri)
	require.NoError(t, err)
	defer pg.Close()

	repo := pgrepo.NewIncidentRepository(pg)

	id, err := repo.Create(context.Background(), entity.Incident{
		Name:        "test incident",
		Date:        time.Now(),
		Country:     "test country",
		Latitude:    50,
		Longitude:   50,
		Publication: time.Now(),
		Comment:     "test comment",
		Status:      entity.Unconfirmed,
		Type:        entity.Bio,
		Author:      1,
	})

	require.NoError(t, err)

	err = repo.Delete(context.Background(), id)
	require.NoError(t, err)
}

func TestGetIncidents(t *testing.T) {
	pg, err := postgres.New(uri)
	require.NoError(t, err)
	defer pg.Close()

	repo := pgrepo.NewIncidentRepository(pg)
	incidents, err := repo.GetAll(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, incidents)
}
