package mgrepo

import (
	"context"
	"fmt"
	"log"

	"github.com/ekaterinamzr/green-alarm/internal/entity"
	"github.com/ekaterinamzr/green-alarm/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IncidentRepository struct {
	*mongo.Mongo
}

func NewIncidentRepository(mongo *mongo.Mongo) *IncidentRepository {
	return &IncidentRepository{mongo}
}

func (r *IncidentRepository) Create(ctx context.Context, i entity.Incident) (int, error) {
	db := r.DB.Database("greenalarm")
	collection := db.Collection("incidents")

	counters := db.Collection("counters")

	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"_id", "incidentid"}}
	update := bson.D{{"$inc", bson.D{{"seq", 1}}}}
	var updatedDocument bson.M
	err := counters.FindOneAndUpdate(
		context.TODO(),
		filter,
		update,
		opts,
	).Decode(&updatedDocument)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("updated document %v", updatedDocument)
	seq := updatedDocument["seq"]
	// id, _ := strconv.Atoi(seq.(string))
	id := int(seq.(int32))

	// fmt.Println(id)

	// res := db.RunCommand(ctx,
	// 	bson.M{
	// 		"findAndModify": "counters",
	// 		"query":         "{ _id: incidentid}",
	// 		"update":        "{$inc: {seq: 1 }}",
	// 		"new":           "true"})
	// var id int
	// // res := db.RunCommand(ctx, bson.M{"eval": "getNextSequence('incidentid');"})
	// fmt.Println(res)
	// res.Decode(&id)
	// fmt.Println(id)

	incident := bson.D{
		{Key: "_id", Value: id},
		{Key: "incident_name", Value: i.Name},
		{Key: "incident_date", Value: i.Date},
		{Key: "country", Value: i.Country},
		{Key: "latitude", Value: i.Latitude},
		{Key: "longitude", Value: i.Longitude},
		{Key: "publication_date", Value: i.Publication},
		{Key: "comment", Value: i.Comment},
		{Key: "incident_status", Value: i.Status},
		{Key: "incident_type", Value: i.Type},
		{Key: "author", Value: i.Author}}
	_, err = collection.InsertOne(ctx, incident)
	if err != nil {
		return 0, fmt.Errorf("mongorepo - incident - Create: %w", err)
	}

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

func (r *IncidentRepository) GetById(ctx context.Context, id int) (*entity.Incident, error) {
	var incident entity.Incident

	collection := r.DB.Database("greenalarm").Collection("incidents")

	// objId, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return nil, fmt.Errorf("mongorepo - incident - GetById: %w", err)
	// }

	filter := bson.D{{Key: "_id", Value: id}}

	err := collection.FindOne(ctx, filter).Decode(&incident)
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

func (r *IncidentRepository) Update(ctx context.Context, updated entity.Incident) error {
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

	// objId, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return fmt.Errorf("mongorepo - incident - Update: %w", err)
	// }

	_, err := collection.UpdateByID(ctx, updated.Id, update)
	if err != nil {
		return fmt.Errorf("mongorepo - incident - Update: %w", err)
	}

	return nil
}

func (r *IncidentRepository) Delete(ctx context.Context, id int) error {
	collection := r.DB.Database("greenalarm").Collection("incidents")

	// objId, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return fmt.Errorf("mongorepo - incident - Delete: %w", err)
	// }

	filter := bson.D{{Key: "_id", Value: id}}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("mongorepo - incident - Delete: %w", err)
	}

	return nil
}