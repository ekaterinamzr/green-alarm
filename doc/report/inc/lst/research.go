package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/ekaterinamzr/green-alarm/config"
	"github.com/ekaterinamzr/green-alarm/internal/entity"
	"github.com/ekaterinamzr/green-alarm/internal/infrastructure/mongorepo"
	"github.com/ekaterinamzr/green-alarm/internal/infrastructure/pgrepo"
	"github.com/ekaterinamzr/green-alarm/pkg/mongo"
	"github.com/ekaterinamzr/green-alarm/pkg/postgres"
)

const (
	n    = 1000
	reps = 10
)

type repo interface {
	Create(context.Context, entity.Incident) (string, error)
	GetAll(context.Context) ([]entity.Incident, error)
	GetById(context.Context, string) (*entity.Incident, error)
	Update(context.Context, string, entity.Incident) error
	Delete(context.Context, string) error
	GetByType(context.Context, int) ([]entity.Incident, error)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString() string {
	c := charset[rand.Intn(len(charset))]
	return string(c)
}

func randomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func randomFloat(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func randomIncident() entity.Incident {
	return entity.Incident{
		Name:             randomString(),
		Date:             time.Now(),
		Country:          randomString(),
		Latitude:         randomFloat(-90.0, 90.0),
		Longitude:        randomFloat(-180.0, 180.0),
		Publication_date: time.Now(),
		Comment:          randomString(),
		Status:           1,
		Type:             randomInt(1, 7),
		Author:           1,
	}
}

func randomIncidents(n int) []entity.Incident {
	rand.Seed(time.Now().UnixNano())
	incidents := make([]entity.Incident, n)
	for i := range incidents {
		incidents[i] = randomIncident()
	}
	return incidents
}

func randomTypes(n int) []int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 7

	types := make([]int, n)
	for i := range types {
		types[i] = randomInt(min, max)
	}
	return types
}

func create(r repo, data []entity.Incident) int64 {
	startedAt := time.Now().UnixNano()
	for i := range data {
		id, _ := r.Create(context.Background(), data[i])
		data[i].Id = id
	}
	finishedAt := time.Now().UnixNano()

	return finishedAt - startedAt
}

func update(r repo, data, updated []entity.Incident) int64 {
	startedAt := time.Now().UnixNano()
	for i := range data {
		r.Update(context.Background(), data[i].Id, updated[i])
	}
	finishedAt := time.Now().UnixNano()

	return finishedAt - startedAt
}

func delete(r repo, data []entity.Incident) int64 {
	startedAt := time.Now().UnixNano()
	for i := range data {
		r.Delete(context.Background(), data[i].Id)
	}
	finishedAt := time.Now().UnixNano()

	return finishedAt - startedAt
}

func getById(r repo, data []entity.Incident) int64 {
	startedAt := time.Now().UnixNano()
	for i := range data {
		r.GetById(context.Background(), data[i].Id)
	}
	finishedAt := time.Now().UnixNano()

	return finishedAt - startedAt
}

func getByType(r repo, types []int) int64 {
	startedAt := time.Now().UnixNano()
	for i := range types {
		r.GetByType(context.Background(), types[i])
	}
	finishedAt := time.Now().UnixNano()

	return finishedAt - startedAt
}

func run(r repo, incidents, updated []entity.Incident, types []int) [5]int64 {
	res := [5]int64{}

	res[0] = create(r, incidents)
	res[1] = update(r, incidents, updated)
	res[3] = getById(r, incidents)
	res[4] = getByType(r, types)
	res[2] = delete(r, incidents)

	return res
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	pgAvg := [5]int64{}
	mongoAvg := [5]int64{}
	for i := 0; i < reps; i++ {
		pg, err := postgres.New(cfg.Database.URI)
		if err != nil {
			log.Fatal(err, "benchmarking - postgres.New")
		}
		mongo, err := mongo.New(cfg.MongoDB.URI)
		if err != nil {
			log.Fatal(err, "benchmarking - mongo.New")
		}

		pgrepo := pgrepo.NewIncidentRepository(pg)
		mongorepo := mongorepo.NewIncidentRepository(mongo)

		incidents := randomIncidents(n)
		updated := randomIncidents(n)
		types := randomTypes(n)

		pgRes := run(pgrepo, incidents, updated, types)
		mongoRes := run(mongorepo, incidents, updated, types)

		for i := 0; i < 5; i++ {
			pgAvg[i] += pgRes[i]
			mongoAvg[i] += mongoRes[i]
		}

		pg.Close()
		mongo.Close()
	}
	for i := 0; i < 5; i++ {
		pgAvg[i] /= reps
		pgAvg[i] = int64(float64(pgAvg[i]) * 1e-6)

		mongoAvg[i] /= reps
		mongoAvg[i] = int64(float64(mongoAvg[i]) * 1e-6)
	}
	fmt.Println(pgAvg)
	fmt.Println(mongoAvg)
}
