package repository

import (
	"data-app-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
	"context"
)

type StockChipRepository struct {
	collection *mongo.Collection
}

func NewStockChipRepository(database *mongo.Database) StockChipRepository {
	return StockChipRepository{collection: database.Collection("stock_chip")}
}

func (repository *StockChipRepository) FindStockChipsByStockCodeAndDateBetween(stockCode string, startDate time.Time, endDate time.Time) []model.StockChip {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := repository.collection.Find(ctx, bson.D{
		{"stock_code", stockCode},
		{"date", bson.D{
			{"$gte", startDate},
			{"$lte", endDate},
		}},
	})
	defer cursor.Close(ctx)
	var results []model.StockChip
	if err != nil {
		log.Fatalf("Find Data err #%v", err)
		return nil
	}
	for cursor.Next(ctx) {
		var result model.StockChip
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	cursor.Close(ctx)
	return results
}
