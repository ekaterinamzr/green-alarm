package mongorepo

import (
	"context"
	"fmt"

	"github.com/ekaterinamzr/green-alarm/internal/entity"
	"github.com/ekaterinamzr/green-alarm/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IncidentRepository struct {
	*mongo.Mongo
}

func NewIncidentRepository(mongo *mongo.Mongo) *IncidentRepository {
	return &IncidentRepository{mongo}
}

func (r *IncidentRepository) Create(ctx context.Context, i entity.Incident) (string, error) {
	collection := r.DB.Database("greenalarm").Collection("incidents")
	incident := bson.D{
		{Key: "incident_name", Value: i.Name},
		{Key: "incident_date", Value: i.Date},
		{Key: "country", Value: i.Country},
		{Key: "latitude", Value: i.Latitude},
		{Key: "longitude", Value: i.Longitude},
		{Key: "publication_date", Value: i.Publication_date},
		{Key: "comment", Value: i.Comment},
		{Key: "incident_status", Value: i.Status},
		{Key: "incident_type", Value: i.Type},
		{Key: "author", Value: i.Author}}
	res, err := collection.InsertOne(ctx, incident)
	if err != nil {
		return "", fmt.Errorf("mongorepo - incident - Create: %w", err)
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func (r *IncidentRepository) GetAll(ctx context.Context) ([]entity.Incident, error) {
	var all []entity.Incident

	collection := r.DB.Database("greenalarm").Collection("incidents")

	filter := bson.D{}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("mongorepo - incident - GetAll: %w", err)
	}

	if err = cursor.All(ctx, &all); err != nil {
		return nil, fmt.Errorf("mongorepo - incident - GetAll: %w", err)
	}

	return all, nil
}

func (r *IncidentRepository) GetById(ctx context.Context, id string) (*entity.Incident, error) {
	var incident entity.Incident

	collection := r.DB.Database("greenalarm").Collection("incidents")

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("mongorepo - incident - GetById: %w", err)
	}

	filter := bson.D{{Key: "_id", Value: objId}}

	err = collection.FindOne(ctx, filter).Decode(&incident)
	if err != nil {
		return nil, fmt.Errorf("mongorepo - incident - GetById: %w", err)
	}

	return &incident, nil
}

func (r *IncidentRepository) GetByType(ctx context.Context, requiredType int) ([]entity.Incident, error) {
	var incidents []entity.Incident

	collection := r.DB.Database("greenalarm").Collection("incidents")

	filter := bson.D{{Key: "incident_type", Value: requiredType}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("mongorepo - incident - GetByType: %w", err)
	}

	if err = cursor.All(ctx, &incidents); err != nil {
		return nil, fmt.Errorf("mongorepo - incident - GetByType: %w", err)
	}

	return incidents, nil
}

func (r *IncidentRepository) Update(ctx context.Context, id string, updated entity.Incident) error {
	collection := r.DB.Database("greenalarm").Collection("incidents")
	
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "incident_name", Value: updated.Name},
		{Key: "incident_date", Value: updated.Date},
		{Key: "country", Value: updated.Country},
		{Key: "latitude", Value: updated.Latitude},
		{Key: "longitude", Value: updated.Longitude},
		{Key: "comment", Value: updated.Comment},
		{Key: "incident_status", Value: updated.Status},
		{Key: "incident_type", Value: updated.Type},
		{Key: "author", Value: updated.Author},
	}}}

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("mongorepo - incident - Update: %w", err)
	}
	_, err = collection.UpdateByID(ctx, objId, update)
	if err != nil {
		return fmt.Errorf("mongorepo - incident - Update: %w", err)
	}
	return nil
}

func (r *IncidentRepository) Delete(ctx context.Context, id string) error {
	collection := r.DB.Database("greenalarm").Collection("incidents")

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("mongorepo - incident - Delete: %w", err)
	}

	filter := bson.D{{Key: "_id", Value: objId}}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("mongorepo - incident - Delete: %w", err)
	}

	return nil
}
