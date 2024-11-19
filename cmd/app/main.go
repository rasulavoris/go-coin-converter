package main

import (
	"encoding/json"
	"fmt"
	"gocoinconverter/internal/application"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		currencies := services.Currencies()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(currencies)
		if err != nil {
			return
		}
	})

	mux.HandleFunc("GET /converter", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Converter!")
		query := r.URL.Query()
		fromCurrency := query.Get("fromCurrency")
		toCurrency := query.Get("toCurrency")

		if fromCurrency == "" || toCurrency == "" {
			http.Error(w, "Missing required parameters: fromCurrency or toCurrency", http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(services.Convert(fromCurrency, toCurrency))
		if err != nil {
			return
		}
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Server listening at http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
