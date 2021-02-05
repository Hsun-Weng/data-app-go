package repository

import (
	"data-app-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
	"context"
)

type FuturesChipRepository struct {
	collection *mongo.Collection
}

func NewFuturesChipRepository(database *mongo.Database) FuturesChipRepository {
	return FuturesChipRepository{collection: database.Collection("futures_chip")}
}

func (repository *FuturesChipRepository) FindFuturesChipsByFuturesCodeAndDateBetween(futuresCode string, startDate time.Time, endDate time.Time) []model.FuturesChip {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := repository.collection.Find(ctx, bson.D{
		{"futures_code", futuresCode},
		{"date", bson.D{
			{"$gte", startDate},
			{"$lte", endDate},
		}},
	})
	defer cursor.Close(ctx)
	var results []model.FuturesChip
	if err != nil {
		log.Fatalf("Find Data err #%v", err)
		return nil
	}
	for cursor.Next(ctx) {
		var result model.FuturesChip
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	cursor.Close(ctx)
	return results
}
