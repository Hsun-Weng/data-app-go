package dto

type FuturesChipDTO struct {
	Date             string                   `json:"date"`
	FuturesCode      string                   `json:"futuresCode"`
	InvestorChipDTOs []InvestorFuturesChipDTO `json:"investorChip"`
}

type InvestorFuturesChipDTO struct {
	InvestorCode         string `json:"investorCode"`
	OpenInterestLongLot  int    `json:"openInterestLongLot"`
	OpenInterestShortLot int    `json:"openInterestShortLot"`
	OpenInterestNetLot   int    `json:"openInterestNetLot"`
	Percent              float64    `json:"percent"`
}
