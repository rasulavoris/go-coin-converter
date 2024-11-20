package models

type Currency struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type ActualCurrencies struct {
	Currencies   map[string]Currency `json:"currencies"`
	ErrorMessage string              `json:"errorMessage"`
}
