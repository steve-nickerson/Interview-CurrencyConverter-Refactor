namespace Interview_CurrencyConverter_Refactor
{
    public interface ICurrencyConverterRepository
    {
        List<CurrencyConversion> GetConversions();
    }
}
