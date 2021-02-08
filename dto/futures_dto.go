package dto

type FuturesDTO struct {
	Date         string  `json:"date"`
	FuturesCode  string  `json:"futuresCode"`
	ContractDate string  `json:"contractDate"`
	Open         float32 `json:"open"`
	Low          float32 `json:"low"`
	High         float32 `json:"high"`
	Close        float32 `json:"close"`
	Volume       int     `json:"volume"`
}
