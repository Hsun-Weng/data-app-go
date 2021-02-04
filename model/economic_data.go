package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type EconomicData struct {
	ID primitive.ObjectID `bson:"_id"`
	Data primitive.DateTime `bson:"date"`
	CountryCode string `bson:"country_code"`
	DataCode string `bson:"data_code"`
	Value int `bson:"value"`
}

