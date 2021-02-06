package repository

import (
	"context"
	"data-app-go/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FuturesRepository struct {
	collection *mongo.Collection
}

func NewFuturesRepository(database *mongo.Database) FuturesRepository {
	return FuturesRepository{collection: database.Collection("futures")}
}

func (repository *FuturesRepository) FindFuturesByFuturesCodeAndDateBetween(futuresCode string, startDate time.Time, endDate time.Time) []model.Futures {
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
	var results []model.Futures
	if err != nil {
		log.Fatalf("Find Data err #%v", err)
		return nil
	}
	for cursor.Next(ctx) {
		var result model.Futures
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	cursor.Close(ctx)
	return results
}

func (repository *FuturesRepository) FindFuturesByFuturesCodeAndContractDateAndDateBetween(futuresCode string, contractDate string, startDate time.Time, endDate time.Time) []model.Futures {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := repository.collection.Find(ctx, bson.D{
		{"futures_code", futuresCode},
		{"contract_date", contractDate},
		{"date", bson.D{
			{"$gte", startDate},
			{"$lte", endDate},
		}},
	})
	defer cursor.Close(ctx)
	var results []model.Futures
	if err != nil {
		log.Fatalf("Find Data err #%v", err)
		return nil
	}
	for cursor.Next(ctx) {
		var result model.Futures
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	cursor.Close(ctx)
	return results
}
