package database

type CurrentPrice struct {
	ID           string  `json:"id"`
	Value        float64 `json:"value"`
	CurrencyCode string  `json:"currency_code"`
	Error        string  `json:"error"`
}
