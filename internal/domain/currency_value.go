package models

type CurrencyValue struct {
	FromCurrency string  `json:"fromCurrency"`
	ToCurrency   string  `json:"toCurrency"`
	Value        float64 `json:"value"`
	ErrorMessage string  `json:"errorMessage"`
}
