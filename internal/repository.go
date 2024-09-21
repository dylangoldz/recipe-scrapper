package internal

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database, collectionName string) *Repository {
	return &Repository{
		collection: db.Collection(collectionName),
	}
}

// insert
func (r *Repository) Insert(ctx context.Context, data interface{}) error {
	_, err := r.collection.InsertOne(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to insert document: %w", err)
	}
	return nil
}

// find
func (r *Repository) Find(ctx context.Context, query interface{}) ([]interface{}, error) {
	cursor, err := r.collection.Find(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to find documents: %w", err)
	}
	defer cursor.Close(ctx)

	var results []interface{}
	if err := cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to decode documents: %w", err)
	}
	return results, nil
}
