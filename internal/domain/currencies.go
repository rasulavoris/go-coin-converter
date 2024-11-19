package models

type Currency struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Currencies map[string]Currency
