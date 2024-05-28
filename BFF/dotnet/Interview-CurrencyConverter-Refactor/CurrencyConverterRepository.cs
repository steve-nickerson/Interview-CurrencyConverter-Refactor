namespace Interview_CurrencyConverter_Refactor
{
    public class CurrencyConverterRepository : ICurrencyConverterRepository
    {

        public List<CurrencyConversion> GetConversions()
        {
            return new List<CurrencyConversion>() {
                new CurrencyConversion() { CountryCode = "USD", CurrencyName = "United States Dollars", RateFromUSDToCurrency = 1.0M },
                new CurrencyConversion() { CountryCode = "CAD", CurrencyName = "Canada Dollars", RateFromUSDToCurrency = 1.2071M },
                new CurrencyConversion() { CountryCode = "MXN", CurrencyName = "Mexico Pesos", RateFromUSDToCurrency = 15.22M },
                new CurrencyConversion() { CountryCode = "CRC", CurrencyName = "Costa Rica Colons", RateFromUSDToCurrency = 538.52M },
                new CurrencyConversion() { CountryCode = "DZD", CurrencyName = "Algeria Dinars", RateFromUSDToCurrency = 97.56M },
                new CurrencyConversion() { CountryCode = "CNY", CurrencyName = "China Renminbis", RateFromUSDToCurrency = 6.08M },
                new CurrencyConversion() { CountryCode = "DKK", CurrencyName = "Denmark Krones", RateFromUSDToCurrency = 6.6181M },
                new CurrencyConversion() { CountryCode = "PLN", CurrencyName = "Poland Zlotys", RateFromUSDToCurrency = 3.6284M }
            };
        }
    }
}
