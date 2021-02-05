package dto

type StockDTO struct {
	Sort          int     `json:"sort"`
	Date          string  `json:"date"`
	StockCode     string  `json:"stockCode"`
	Open          float32 `json:"open"`
	Low           float32 `json:"low"`
	High          float32 `json:"high"`
	Close         float32 `json:"close"`
	Volume        int     `json:"volume"`
	Change        int     `json:"change"`
	ChangePercent float64 `json:"changePercent"`
}

type StockPageDTO struct {
	TotalPage int64 `json:"totalPage"`
	Page int `json:"page"`
	Size int64 `json:"size"`
	Content []StockDTO `json:"content"`
}
