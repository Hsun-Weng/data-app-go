package dto

type StockMarginDTO struct {
	Date            string `json:"date"`
	StockCode       string `json:"stockCode"`
	LongShare       int    `json:"longShare"`
	TotalLongShare  int    `json:"totalLongShare"`
	ShortShare      int    `json:"shortShare"`
	TotalShortShare int    `json:"totalShortShare"`
	DayShare        int    `json:"dayShare"`
}
