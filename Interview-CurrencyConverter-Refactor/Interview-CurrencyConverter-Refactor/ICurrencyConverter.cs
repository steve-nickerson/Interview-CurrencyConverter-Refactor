namespace Interview_CurrencyConverter_Refactor
{
    public interface ICurrencyConverter
    {
        decimal GetConvertedAmount(string from, string to, decimal amount);
    }
}