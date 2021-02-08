package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Futures struct {
	ID              primitive.ObjectID `bson:"_id"`
	Date            time.Time `bson:"date"`
	FuturesCode     string             `bson:"futures_code"`
	ContractDate    string             `bson:"contract_date"`
	Open            float32            `bson:"open"`
	Low             float32            `bson:"low"`
	High            float32            `bson:"high"`
	Close           float32            `bson:"close"`
	Volume          int                `bson:"volume"`
	OpenInterestLot int                `bson:"open_interest_lot"`
}
