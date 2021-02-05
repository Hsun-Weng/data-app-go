package repository

import (
	"data-app-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"context"
)

type StockRepository struct {
	collection *mongo.Collection
}

func NewStockRepository(database *mongo.Database) StockRepository {
	return StockRepository{collection: database.Collection("stock")}
}

func (repository *StockRepository) FindStocksByStockCodeAndDateBetween(stockCode string, startDate time.Time, endDate time.Time) []model.Stock {
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
	var results []model.Stock
	if err != nil {
		log.Fatalf("Find Data err #%v", err)
		return nil
	}
	for cursor.Next(ctx) {
		var result model.Stock
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	cursor.Close(ctx)
	return results
}

func (repository *StockRepository) FindStockLatestByStockCode(stockCode string) model.Stock {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	options := options.FindOne().SetSort(bson.D{{"date", -1}})
	var result model.Stock
	err := repository.collection.FindOne(ctx, bson.D{
		{"stock_code", stockCode},
	}, options).Decode(&result)
	if err != nil {
		log.Fatalf("Find Data err #%v", err)
		return model.Stock{}
	}
	return result
}

func (repository *StockRepository) FindStockLatest() (model.Stock, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	options := options.FindOne().SetSort(bson.D{{"date", -1}})
	var result model.Stock
	err := repository.collection.FindOne(ctx, bson.D{}, options).Decode(&result)
	if err != nil {
		log.Fatalf("Find Data err #%v", err)
		return model.Stock{}, err
	}
	return result, nil
}


func (repository *StockRepository) FindStocksByStockCodesAndDateBetween(stockCodes []string, startDate time.Time, endDate time.Time) []model.Stock {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := repository.collection.Find(ctx, bson.D{
		{"stock_code", bson.D{
			{"$in", stockCodes},
		}},
		{"date", bson.D{
			{"$gte", startDate},
			{"$lte", endDate},
		}},
	})
	defer cursor.Close(ctx)
	var results []model.Stock
	if err != nil {
		log.Fatalf("Find Data err #%v", err)
		return nil
	}
	for cursor.Next(ctx) {
		var result model.Stock
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	cursor.Close(ctx)
	return results
}

func (repository *StockRepository) CountStocksByDateBetween(startDate time.Time, endDate time.Time) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	count, err := repository.collection.CountDocuments(ctx, bson.D{
		{"date", bson.D{
			{"$gte", startDate},
			{"$lte", endDate},
		}},
	})
	if err != nil {
		log.Fatalf("Count Data err #%v", err)
		return 0
	}
	return count
}

func (repository *StockRepository) FindStocksByDateBetweenAndSortByAndLimit(startDate time.Time, endDate time.Time, sortColumn string, direction int, limit int64) []model.Stock {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	options := options.Find().SetSort(bson.D{{sortColumn, direction}}).SetLimit(limit)
	cursor, err := repository.collection.Find(ctx, bson.D{
		{"date", bson.D{
			{"$gte", startDate},
			{"$lte", endDate},
		}},
	}, options)
	defer cursor.Close(ctx)
	var results []model.Stock
	if err != nil {
		log.Fatalf("Find Data err #%v", err)
		return nil
	}
	for cursor.Next(ctx) {
		var result model.Stock
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	cursor.Close(ctx)
	return results
}

