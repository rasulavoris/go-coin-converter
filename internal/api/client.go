package api

import (
	"encoding/json"
	"gocoinconverter/pkg"
)

func ExternalCurrencies() map[string]string {
	var currencies map[string]string
	body := pkg.DoRequest("GET", "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies.json")

	err := json.Unmarshal(body, &currencies)
	if err != nil {
		return currencies
	}

	return currencies
}

func ExternalCurrencyValueConversion(value string) map[string]interface{} {
	var currenciesValueConversion map[string]interface{}
	body := pkg.DoRequest("GET", "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/"+value+".json")

	err := json.Unmarshal(body, &currenciesValueConversion)
	if err != nil {
		return currenciesValueConversion
	}

	return currenciesValueConversion
}
