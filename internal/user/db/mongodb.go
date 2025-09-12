package db

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/xyersh/examle-REST-app/internal/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *slog.Logger
}

// Create implements user.Storage.
func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	slog.Debug("create user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create users: %w", err)
	}
	slog.Debug("cover insertedID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}

	// !ok
	d.logger.Debug("username", user.Username, "email", user.Email)
	return "", fmt.Errorf("failed to convert objectid to hex, probably: %s", oid)

}

// FindOne implements user.Storage.
func (d *db) FindOne(ctx context.Context, id string) (user user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, fmt.Errorf("failed to convert hex to ObjectID: %s", id)
	}

	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)
	if err := result.Err(); err != nil {

		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			//TODO ErrEntityNotFound
			return user, fmt.Errorf("ErrEntityNotFound")
		}
		return user, fmt.Errorf("failed to find one user by id: %s due to %w", id, err)
	}

	if err := result.Decode(&user); err != nil {
		return user, fmt.Errorf("failed to decode user (id: %s) from DB due to %w", id, err)
	}

	return user, nil
}

// Delete implements user.Storage.
func (d *db) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert hex to ObjectID: %s", id)
	}

	filter := bson.M{"_id": oid}

	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w ", err)
	}
	if result.DeletedCount == 0 {
		// TODO ErrEntityNotFound
		return fmt.Errorf("not found: %w", err)
	}

	slog.Debug(fmt.Sprintf("Deleted %d documents", result.DeletedCount))
	return nil
}

// Update implements user.Storage.
func (d *db) Update(ctx context.Context, user user.User) error {
	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return fmt.Errorf("failedto convert hex to ObjectID: %w", err)
	}

	filter := bson.M{"_id": objectID}
	userBytes, err := bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user: %w", err)
	}

	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal userBytes: %w", err)
	}

	delete(updateUserObj, "_id")

	update := bson.M{
		"$set": updateUserObj,
	}

	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute update user query. error: %w", err)
	}

	if result.MatchedCount == 0 {
		// TODO ErrEntityNotFound
		return fmt.Errorf("not found: %w", err)
	}

	slog.Debug(fmt.Sprintf("Matched %d documents, modified %d", result.MatchedCount, result.ModifiedCount))
	return nil

}

func NewStorage(database *mongo.Database, collection string, logger *slog.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
