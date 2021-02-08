package dto

type StockIndexDTO struct {
	Date          string  `json:"date"`
	IndexCode     string  `json:"indexCode"`
	Open          float32 `json:"open"`
	Low           float32 `json:"low"`
	High          float32 `json:"high"`
	Close         float32 `json:"close"`
	Volume        int     `json:"volume"`
	Change        int     `json:"change"`
	ChangePercent float64 `json:"changePercent"`
}
