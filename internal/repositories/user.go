package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/diazharizky/go-mongodb-with-tests/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	coll *mongo.Collection
}

func NewUserRepository(db *mongo.Database) userRepository {
	return userRepository{
		coll: db.Collection("users"),
	}
}

func (repo userRepository) List(ctx context.Context) ([]models.User, error) {
	cur, err := repo.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []models.User

	if err = cur.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (repo userRepository) Get(ctx context.Context, id string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": objectID,
	}

	user := models.User{}
	if err := repo.coll.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo userRepository) Create(ctx context.Context, newUser models.User) (*string, error) {
	if newUser.ID == primitive.NilObjectID {
		newUser.ID = primitive.NewObjectID()
	}

	now := time.Now()
	if newUser.CreatedAt.IsZero() {
		newUser.CreatedAt = now
	}

	res, err := repo.coll.InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to resolve inserted id")
	}

	newID := id.Hex()

	return &newID, nil
}

func (repo userRepository) Update(ctx context.Context, id string, values models.User) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": objectID,
	}

	update := bson.D{{Key: "$set", Value: values.UpdateFields()}}

	_, err = repo.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo userRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": objectID,
	}

	_, err = repo.coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
