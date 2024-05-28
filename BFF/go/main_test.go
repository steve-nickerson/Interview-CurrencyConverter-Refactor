package main

import (
	"testing"
)

type MockCurrencyConverterRepository struct{}

func (m *MockCurrencyConverterRepository) GetConversions() []CurrencyConversion {
	return []CurrencyConversion{
		{CountryCode: "CAD", CurrencyName: "Canada Dollars", RateFromUSDToCurrency: 2},
		{CountryCode: "MXN", CurrencyName: "Mexico Pesos", RateFromUSDToCurrency: 20},
	}
}

// TODO: Add tests for handlers

// TODO: Add tests for repository

func TestCurrencyConverter_GetConvertedAmount(t *testing.T) {
	mockRepo := &MockCurrencyConverterRepository{}
	currencyConverter := &CurrencyConverter{repo: mockRepo}

	t.Run("GetConvertedAmount_ConvertsCorrectly", func(t *testing.T) {
		actual, err := currencyConverter.GetConvertedAmount("CAD", "MXN", 10)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if actual != 100 {
			t.Errorf("Expected 100, got %v", actual)
		}
	})

	t.Run("GetConvertedAmount_ConvertsCorrectlyWithMismatchedCase", func(t *testing.T) {
		actual, err := currencyConverter.GetConvertedAmount("caD", "mXn", 10)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if actual != 100 {
			t.Errorf("Expected 100, got %v", actual)
		}
	})

	t.Run("GetConvertedAmount_UnknownFromCountryCodeThrowsException", func(t *testing.T) {
		_, err := currencyConverter.GetConvertedAmount("XXX", "MXN", 10)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}

		if err.Error() != "Invalid country code: XXX" {
			t.Errorf("Expected 'Invalid country code', got %v", err)
		}
	})

	t.Run("GetConvertedAmount_UnknownToCountryCodeThrowsException", func(t *testing.T) {
		_, err := currencyConverter.GetConvertedAmount("CAD", "XXX", 10)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}

		if err.Error() != "Invalid country code: XXX" {
			t.Errorf("Expected 'Invalid country code', got %v", err)
		}
	})
}
