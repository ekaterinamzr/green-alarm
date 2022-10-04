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

type UserRepository struct {
	*mongo.Mongo
}

func NewUserRepository(mongo *mongo.Mongo) *UserRepository {
	return &UserRepository{mongo}
}

func (r *UserRepository) Create(ctx context.Context, u entity.User) (int, error) {
	db := r.DB.Database("greenalarm")
	collection := db.Collection("users")

	counters := db.Collection("counters")

	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"_id", "userid"}}
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
	id := int(seq.(int32))

	incident := bson.D{
		{Key: "_id", Value: id},
		{Key: "first_name", Value: u.FirstName},
		{Key: "last_name", Value: u.LastName},
		{Key: "username", Value: u.Username},
		{Key: "email", Value: u.Email},
		{Key: "user_password", Value: u.Password},
		{Key: "user_role", Value: u.Role}}
	_, err = collection.InsertOne(ctx, incident)
	if err != nil {
		return 0, fmt.Errorf("mongorepo - user - Create: %w", err)
	}

	return id, nil
}

func (r *UserRepository) GetUser(ctx context.Context, username, password string) (*entity.User, error) {
	var user entity.User

	collection := r.DB.Database("greenalarm").Collection("users")

	filter := bson.D{{Key: "username", Value: username}, {Key: "user_password", Value: password}}

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("mongorepo - user - GetUser: %w", err)
	}

	return &user, nil
}
