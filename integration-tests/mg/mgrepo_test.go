package integrationtests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ekaterinamzr/green-alarm/internal/entity"
	"github.com/ekaterinamzr/green-alarm/internal/infrastructure/mgrepo"
	"github.com/ekaterinamzr/green-alarm/pkg/mongo"
	"github.com/stretchr/testify/require"
)

const uri = "mongodb://greenalarm:greenalarm@localhost:27017"

func TestCreateDeleteIncident(t *testing.T) {
	mg, err := mongo.New(uri)
	require.NoError(t, err)
	defer mg.Close()

	repo := mgrepo.NewIncidentRepository(mg)

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

	fmt.Println(id)

	require.NoError(t, err)

	err = repo.Delete(context.Background(), id)
	require.NoError(t, err)
}

func TestGetIncidents(t *testing.T) {
	mg, err := mongo.New(uri)
	require.NoError(t, err)
	defer mg.Close()

	repo := mgrepo.NewIncidentRepository(mg)
	incidents, err := repo.GetAll(context.Background())

	fmt.Print(incidents)

	require.NoError(t, err)
	require.NotEmpty(t, incidents)
}
