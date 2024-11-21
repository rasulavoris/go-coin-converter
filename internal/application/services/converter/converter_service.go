package converter

import (
	"fmt"
	"gocoinconverter/internal/api"
	"gocoinconverter/internal/application/services/cacher"
	models "gocoinconverter/internal/domain"
	"log"
	"time"
)

type ConverterService struct {
	cacheService *cacher.CacheService
}

func NewConverterService(cacheService *cacher.CacheService) *ConverterService {
	return &ConverterService{cacheService: cacheService}
}

func (cs *ConverterService) Currencies() models.ActualCurrencies {
	now := time.Now().Format("2006-1-2")
	data, b := cs.cacheService.GetData(now)

	if b {
		fmt.Println("Obtaining actual currencies data from memcache")
		if obj, ok := data.(models.ActualCurrencies); ok {
			return obj
		} else {
			return models.ActualCurrencies{
				ErrorMessage: fmt.Sprintf("Can not cast to ActualCurrencies object %s", data),
			}
		}
	}

	fmt.Println("Obtaining data from api")
	rawCurrencies := api.ExternalCurrencies()
	currencies := mapToCurrencies(rawCurrencies)

	log.Print("Saving data into memcache...")
	cs.cacheService.SaveData(now, currencies)
	log.Print("Saving data into redis...")
	cs.cacheService.SaveIntoRedis(now, currencies, 30)

	return currencies
}

func (cs *ConverterService) Convert(fromCurrency string, toCurrency string) models.CurrencyValue {
	currencies := (*ConverterService).Currencies(cs)
	existsFromCurrency := exists(currencies.Currencies, fromCurrency)
	existsToCurrency := exists(currencies.Currencies, toCurrency)

	data, b := cs.cacheService.GetData(fromCurrency + toCurrency)

	if b {
		log.Print("Obtaining currency values data from memcache")
		if obj, ok := data.(models.CurrencyValue); ok {
			return obj
		} else {
			return models.CurrencyValue{
				ErrorMessage: fmt.Sprintf("Can not cast to CurrencyValue object %s", data),
			}
		}
	}

	if existsFromCurrency && existsToCurrency {
		currencyValues := api.ExternalCurrencyValueConversion(fromCurrency)
		coinValues := obtainCoinsValue(currencyValues, fromCurrency)

		response := models.CurrencyValue{
			FromCurrency: fromCurrency,
			ToCurrency:   toCurrency,
			Value:        coinValues[toCurrency],
		}

		log.Print("Saving data into memcache...")
		cs.cacheService.SaveData(fromCurrency+toCurrency, response)
		log.Print("Saving data into redis...")
		cs.cacheService.SaveIntoRedis(fromCurrency+toCurrency, response, 30)

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

func mapToCurrencies(rawCurrencies interface{}) models.ActualCurrencies {
	result, ok := rawCurrencies.(map[string]string)
	if ok {
		actualCurrencies := models.ActualCurrencies{
			Currencies: make(map[string]models.Currency),
		}

		for code, description := range result {
			actualCurrencies.Currencies[code] = models.Currency{
				Code:        code,
				Description: description,
			}
		}

		return actualCurrencies
	} else {
		return models.ActualCurrencies{
			ErrorMessage: fmt.Sprintf("Unexpected error, expected map[string]string, got %v", rawCurrencies),
		}
	}
}
