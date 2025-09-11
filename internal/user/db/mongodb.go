package db

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/xyersh/examle-REST-app/internal/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *slog.Logger
}

// Create implements user.Storage.
func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create users: %w", err)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}

	d.logger.Debug("username", user.Username, "email", user.Email)
	return "", fmt.Errorf("failed to convert objectid to hex: %w", err)

}

// Delete implements user.Storage.
func (d *db) Delete(ctx *context.Context, id string) error {
	panic("unimplemented")
}

// FindOne implements user.Storage.
func (d *db) FindOne(ctx *context.Context, id string) (user.User, error) {
	panic("unimplemented")
}

// Update implements user.Storage.
func (d *db) Update(ctx *context.Context, user user.User) error {
	panic("unimplemented")
}

func NewStorage(db_ *mongo.Database, collection string, logger *slog.Logger) user.Storage {
	return &db{
		collection: db.Collection(collection),
		logger:     logger,
	}
}
