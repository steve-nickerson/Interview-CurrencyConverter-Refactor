namespace Interview_CurrencyConverter_Refactor
{
    public class CurrencyConverter : ICurrencyConverter
    {
        public ICurrencyConverterRepository repo;

        public CurrencyConverter(ICurrencyConverterRepository r)
        {
            repo = r;
        }

        public decimal GetConvertedAmount(string from, string to, decimal amount)
        {
            List<CurrencyConversion> conversions = repo.GetConversions();
            
            CurrencyConversion? conversion = null;
            foreach (CurrencyConversion c in conversions)
            {
                if (c.CountryCode == from)
                {
                    conversion = c;
                    break;
                }
            }

            if (conversion == null)
                throw new ArgumentException(string.Format("Invalid country code: {0}", from));

            decimal fromRate = conversion.RateFromUSDToCurrency;
            
            CurrencyConversion? conversion1 = null;
            foreach (CurrencyConversion c1 in conversions)
            {
                if (c1.CountryCode == to)
                {
                    conversion1 = c1;
                    break;
                }
            }

            if (conversion1 == null)
                throw new ArgumentException(string.Format("Invalid country code: {0}", to));

            decimal toRate = conversion1.RateFromUSDToCurrency;
            
            return amount / fromRate * toRate;
        }
    }
}
