package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CurrencyConversion struct {
	CountryCode           string
	CurrencyName          string
	RateFromUSDToCurrency float64
}

type CurrencyConverterRepository interface {
	GetConversions() []CurrencyConversion
}

type CurrencyConverter struct {
	repo CurrencyConverterRepository
}

func NewCurrencyConverter(repo CurrencyConverterRepository) *CurrencyConverter {
	return &CurrencyConverter{repo: repo}
}

type CurrencyConverterRepositoryImpl struct{}

func (r *CurrencyConverterRepositoryImpl) GetConversions() []CurrencyConversion {
	return []CurrencyConversion{
		{CountryCode: "USD", CurrencyName: "United States Dollars", RateFromUSDToCurrency: 1.0},
		{CountryCode: "CAD", CurrencyName: "Canada Dollars", RateFromUSDToCurrency: 1.2071},
		{CountryCode: "MXN", CurrencyName: "Mexico Pesos", RateFromUSDToCurrency: 15.22},
		{CountryCode: "CRC", CurrencyName: "Costa Rica Colons", RateFromUSDToCurrency: 538.52},
		{CountryCode: "DZD", CurrencyName: "Algeria Dinars", RateFromUSDToCurrency: 97.56},
		{CountryCode: "CNY", CurrencyName: "China Renminbis", RateFromUSDToCurrency: 6.08},
		{CountryCode: "DKK", CurrencyName: "Denmark Krones", RateFromUSDToCurrency: 6.6181},
		{CountryCode: "PLN", CurrencyName: "Poland Zlotys", RateFromUSDToCurrency: 3.6284},
	}
}

func (c *CurrencyConverter) GetConvertedAmount(from string, to string, amount float64) (float64, error) {
	conversions := c.repo.GetConversions()

	var conversion *CurrencyConversion = nil

	for _, c := range conversions {
		if c.CountryCode == from {
			conversion = &c
			break
		}
	}

	if conversion == nil {
		return 0.0, fmt.Errorf("Invalid country code: %s", from)
	}

	fromRate := conversion.RateFromUSDToCurrency

	var conversion1 *CurrencyConversion = nil

	for _, c1 := range conversions {
		if c1.CountryCode == to {
			conversion1 = &c1
			break
		}
	}

	if conversion1 == nil {
		return 0.0, fmt.Errorf("Invalid country code: %s", to)
	}

	toRate := conversion1.RateFromUSDToCurrency

	return amount / fromRate * toRate, nil
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handleRoot).Methods(http.MethodGet)
	r.HandleFunc("/api/currencyconverter", handleCurrencyConverter).Methods(http.MethodGet)

	fmt.Println("Server is running on port 58415")
	http.ListenAndServe(":58415", r)
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
	port := 58415 // replace with your actual port number
	htmlContent := `
		<html lang="en">
		<head>
		  <title>Sample Request</title>
		</head>
		<body>
		  <h2>Sample Request</h2>
		  <a href="http://localhost:%d/api/currencyconverter?from=USD&to=MXN&amount=100">http://localhost:%d/api/currencyconverter?from=USD&to=MXN&amount=100</a>
		</body>
		</html>
		`
	htmlContent = fmt.Sprintf(htmlContent, port, port)
	fmt.Fprintf(w, htmlContent)
}

func handleCurrencyConverter(w http.ResponseWriter, req *http.Request) {
	amountStr := req.URL.Query().Get("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		// handle error
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	from := req.URL.Query().Get("from")
	to := req.URL.Query().Get("to")

	convertedAmount, err := NewCurrencyConverter(&CurrencyConverterRepositoryImpl{}).GetConvertedAmount(from, to, amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "%f", convertedAmount)
}
