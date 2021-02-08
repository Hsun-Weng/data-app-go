package dto

type StockChipDTO struct {
	Date             string                 `json:"date"`
	StockCode        string                 `json:"stockCode"`
	NetShare         int                    `json:"netShare"`
	InvestorChipDTOs []InvestorStockChipDTO `json:"investorChip"`
}

type InvestorStockChipDTO struct {
	InvestorCode string `json:"investorCode"`
	LongShare    int    `json:"longShare"`
	ShortShare   int    `json:"shortShare"`
}
