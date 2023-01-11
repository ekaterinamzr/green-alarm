package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	DB *mongo.Client
}

const (
	connAttempts = 5
	connTimeout  = 5 * time.Second
)

func New(URI string) (*Mongo, error) {
	// Create client
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		return nil, fmt.Errorf("Could not create MongoDB client: %w", err)
	}

	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("Could not connect to MongoDB: %w", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < connAttempts; i++ {
		err = client.Ping(context.TODO(), nil)

		if err == nil {
			break
		}

		log.Printf("Trying to connect to MongoDB, attempts left: %d", connAttempts-i-1)
		time.Sleep(connTimeout)
	}

	if err != nil {
		return nil, fmt.Errorf("Could not connect to MongoDB: %w", err)
	}

	return &Mongo{DB: client}, nil
}

func (mongo *Mongo) Close() error {
	return mongo.DB.Disconnect(context.TODO())
}
