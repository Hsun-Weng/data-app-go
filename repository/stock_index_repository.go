package repository

import (
	"context"
	"data-app-go/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StockIndexRepository struct {
	collection *mongo.Collection
}

func NewStockIndexRepository(database *mongo.Database) StockIndexRepository {
	return StockIndexRepository{collection: database.Collection("stock_index")}
}

func (repository *StockIndexRepository) FindStockIndexesByIndexCodeAndDateBetween(indexCode string, startDate time.Time, endDate time.Time) []*model.StockIndex {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := repository.collection.Find(ctx, bson.D{
		{"index_code", indexCode},
		{"date", bson.D{
			{"$gte", startDate},
			{"$lte", endDate},
		}},
	})
	defer cursor.Close(ctx)
	var results []*model.StockIndex
	if err != nil {
		log.Fatalf("Find Data err #%v", err)
		return nil
	}
	for cursor.Next(ctx) {
		var result model.StockIndex
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &result)
	}

	cursor.Close(ctx)
	return results
}

func (repository *StockIndexRepository) FindStockIndexLatestByIndexCode(indexCode string) *model.StockIndex {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	options := options.FindOne().SetSort(bson.D{{"date", -1}})
	var result model.StockIndex
	err := repository.collection.FindOne(ctx, bson.D{
		{"index_code", indexCode},
	}, options).Decode(&result)
	if err != nil {
		log.Fatalf("Find Data err #%v", err)
		return nil
	}
	return &result
}

func (repository *StockIndexRepository) FindStockIndexLatest() (model.StockIndex, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	options := options.FindOne().SetSort(bson.D{{"date", -1}})
	var result model.StockIndex
	err := repository.collection.FindOne(ctx, bson.D{}, options).Decode(&result)
	if err != nil {
		log.Fatalf("Find Data err #%v", err)
		return model.StockIndex{}, err
	}
	return result, nil
}

func (repository *StockIndexRepository) FindStockIndexesByIndexCodesAndDateBetween(indexCodes []string, startDate time.Time, endDate time.Time) []*model.StockIndex {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := repository.collection.Find(ctx, bson.D{
		{"index_code", bson.D{
			{"$in", indexCodes},
		}},
		{"date", bson.D{
			{"$gte", startDate},
			{"$lte", endDate},
		}},
	})
	defer cursor.Close(ctx)
	var results []*model.StockIndex
	if err != nil {
		log.Fatalf("Find Data err #%v", err)
		return nil
	}
	for cursor.Next(ctx) {
		var result model.StockIndex
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &result)
	}

	cursor.Close(ctx)
	return results
}
