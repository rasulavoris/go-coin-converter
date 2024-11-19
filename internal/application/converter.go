package services

import (
	"fmt"
	"gocoinconverter/internal/api"
	"gocoinconverter/internal/domain"
)

func Currencies() models.Currencies {
	currencies := make(models.Currencies)
	rawCurrencies := api.ExternalCurrencies()

	for code, description := range rawCurrencies {
		currencies[code] = models.Currency{
			Code:        code,
			Description: description,
		}
	}

	return currencies
}

func Convert(fromCurrency string, toCurrency string) models.CurrencyValue {
	currencies := Currencies()
	existsFromCurrency := exists(currencies, fromCurrency)
	existsToCurrency := exists(currencies, toCurrency)

	if existsFromCurrency && existsToCurrency {
		currencyValues := api.ExternalCurrencyValueConversion(fromCurrency)
		coinValues := obtainCoinsValue(currencyValues, fromCurrency)

		response := models.CurrencyValue{
			FromCurrency: fromCurrency,
			ToCurrency:   toCurrency,
			Value:        coinValues[toCurrency],
		}

		return response
	}

	return models.CurrencyValue{
		ErrorMessage: fmt.Sprintf("Coin %s or %s does not exist", fromCurrency, toCurrency),
	}
}

func exists(currencies map[string]models.Currency, currency string) bool {
	_, exists := currencies[currency]

	if !exists {
		return false
	}

	return exists
}

func obtainCoinsValue(currencyValues map[string]interface{}, fromCurrency string) map[string]float64 {
	coinValueMap := make(map[string]float64)

	if values, ok := currencyValues[fromCurrency].(map[string]interface{}); ok {
		for key, value := range values {
			if floatValue, ok := value.(float64); ok {
				coinValueMap[key] = floatValue
			}
		}
	}

	return coinValueMap
}
