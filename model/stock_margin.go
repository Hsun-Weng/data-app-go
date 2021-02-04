package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type StockMargin struct {
	ID              primitive.ObjectID `bson:"_id"`
	Date            primitive.DateTime `bson:"date"`
	StockCode       string             `bson:"stock_code"`
	LongShare       int                `bson:"long_share"`
	TotalLongShare  int                `bson:"total_long_share"`
	ShortShare      int                `bson:"short_share"`
	TotalShortShare int                `bson:"total_short_share"`
	DayShare        int                `bson:"day_share"`
}
