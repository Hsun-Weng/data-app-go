package repository

import (
	"context"
	"data-app-go/model"
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

func (repository *EconomicDataRepository) FindByCountryCodeAndDataCode(countryCode string, dataCode string) []*model.EconomicData {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := repository.collection.Find(ctx, bson.D{
		{"country_code", countryCode},
		{"data_code", dataCode},
	})
	defer cursor.Close(ctx)
	var results []*model.EconomicData
	if err != nil {
		log.Fatalf("Find Data err #%v", err)
		return nil
	}
	for cursor.Next(ctx) {
		var result model.EconomicData
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &result)
	}

	cursor.Close(ctx)
	return results
}
