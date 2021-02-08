package repository

import (
	"context"
	"data-app-go/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StockMarginRepository struct {
	collection *mongo.Collection
}

func NewStockMarginRepository(database *mongo.Database) StockMarginRepository {
	return StockMarginRepository{collection: database.Collection("stock_margin")}
}

func (repository *StockMarginRepository) FindStockMarginsByStockCodeAndDateBetween(stockCode string, startDate time.Time, endDate time.Time) []*model.StockMargin {
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
	var results []*model.StockMargin
	if err != nil {
		log.Fatalf("Find Data err #%v", err)
		return nil
	}
	for cursor.Next(ctx) {
		var result model.StockMargin
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &result)
	}

	cursor.Close(ctx)
	return results
}
