package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type StockIndex struct {
	ID            primitive.ObjectID `bson:"_id"`
	Date          time.Time `bson:"date"`
	IndexCode     string             `bson:"index_code"`
	Open          float32            `bson:"open"`
	Low           float32            `bson:"low"`
	High          float32            `bson:"high"`
	Close         float32            `bson:"close"`
	Volume        int                `bson:"volume"`
	Change        int                `bson:"change"`
	ChangePercent float64            `bson:"change_percent"`
}
