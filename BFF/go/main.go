package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
		return 0.0, fmt.Errorf("invalid country code: %s", from)
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
		return 0.0, fmt.Errorf("invalid country code: %s", to)
	}

	toRate := conversion1.RateFromUSDToCurrency

	return amount / fromRate * toRate, nil
}

func main() {
	// Create a new CORS handler with liberal settings
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},                                       // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allow all methods
		AllowedHeaders:   []string{"*"},                                       // Allow all headers
		AllowCredentials: true,                                                // Allow credentials
	})
	r := mux.NewRouter()

	r.HandleFunc("/", handleRoot).Methods(http.MethodGet)
	r.HandleFunc("/api/currencyconverter", handleCurrencyConverter).Methods(http.MethodGet)

	// Wrap the router with the CORS middleware
	handler := c.Handler(r)

	fmt.Println("Server is running on port 58415")
	err := http.ListenAndServe(":58415", handler)
	if err != nil {
		return
	}
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
	if _, err := fmt.Fprintf(w, htmlContent); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func handleCurrencyConverter(w http.ResponseWriter, req *http.Request) {
	amountStr := req.URL.Query().Get("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	from := req.URL.Query().Get("from")
	to := req.URL.Query().Get("to")

	convertedAmount, err := NewCurrencyConverter(&CurrencyConverterRepositoryImpl{}).GetConvertedAmount(from, to, amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set Content-Type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Create a JSON response
	response := map[string]interface{}{
		"from":            from,
		"to":              to,
		"amount":          amount,
		"convertedAmount": fmt.Sprintf("%.2f", convertedAmount),
	}

	// Encode and write the JSON response, handling any errors
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
