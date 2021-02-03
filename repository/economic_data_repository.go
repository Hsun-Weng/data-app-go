package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EconomicDataRepository struct {
	collection *mongo.Collection
}

func NewEconomicDataRepository(database *mongo.Database) EconomicDataRepository {
	return EconomicDataRepository{collection: database.Collection("economic_data")}
}

func (repository *EconomicDataRepository) FindAll() []string {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := repository.collection.Find(ctx, bson.D{})
	defer cursor.Close(ctx)
	if err != nil {
		log.Fatalf("Find Data err #%v", err)
		return nil
	}
	return nil
}
