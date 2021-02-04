package dto

type EconomicDataDTO struct {
	Date  string `bson:"date"`
	Value int    `bson:"value"`
}
