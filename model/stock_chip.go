package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type StockChip struct {
	ID            primitive.ObjectID  `bson:"_id"`
	Date          primitive.DateTime  `bson:"date"`
	StockCode     string              `bson:"stock_code"`
	NetShare      int                 `bson:"net_share"`
	InvestorChips []InvestorStockChip `bson:"investor_chip"`
}

type InvestorStockChip struct {
	InvestorCode string `bson:"investor_code"`
	LongShare    int    `bson:"long_share"`
	ShortShare   int    `bson:"short_share"`
}
