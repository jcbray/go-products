package products

type Products struct {
	ID                 string `json:"id"`
	Dpci               string `json:"dpci"`
	Upc                string `json:"upc"`
	ProductDescription string `json:"product_description"`
	CurrentPrice       struct {
		Value        float64 `json:"value"`
		CurrencyCode string  `json:"currency_code"`
	} `json:"current_price"`
}
