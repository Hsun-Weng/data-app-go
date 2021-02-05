package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type FuturesChip struct {
	ID            primitive.ObjectID    `bson:"_id"`
	Date          time.Time    `bson:"date"`
	FuturesCode   string                `bson:"futures_code"`
	InvestorChips []InvestorFuturesChip `bson:"investor_chip"`
}

type InvestorFuturesChip struct {
	InvestorCode            string `bson:"investor_code"`
	LongLot                 int    `bson:"long_lot"`
	LongAmount              int    `bson:"long_amount"`
	ShortLot                int    `bson:"short_lot"`
	ShortAmount             int    `bson:"short_amount"`
	OpenInterestLongLot     int    `bson:"open_interest_long_lot"`
	OpenInterestLongAmount  int    `bson:"open_interest_long_amount"`
	OpenInterestShortLot    int    `bson:"open_interest_short_lot"`
	OpenInterestShortAmount int    `bson:"open_interest_short_amount"`
	Percent float64    `bson:"percent"`
}
