using Moq;

namespace Interview_CurrencyConverter_Refactor.Tests;

public class CurrencyConverterTests
{
    private List<CurrencyConversion> currency_Conversions;
    private CurrencyConverter currency_Converter;

    public CurrencyConverterTests()
    {
        currency_Conversions = new List<CurrencyConversion>()
        {
            new CurrencyConversion { CountryCode = "CAD", RateFromUSDToCurrency = 2m },
            new CurrencyConversion { CountryCode = "MXN", RateFromUSDToCurrency = 20m }
        };

        Mock<ICurrencyConverterRepository> mockRepo = new Mock<ICurrencyConverterRepository>();
        mockRepo.Setup(r => r.GetConversions()).Returns(currency_Conversions);

        currency_Converter = new CurrencyConverter(mockRepo.Object);
    }


    [Fact]
    public void GetConvertedAmount_ConvertsCorrectly()
    {
        decimal actual = currency_Converter.GetConvertedAmount("CAD", "MXN", 10);

        Assert.Equal(100m, actual);
    }

    [Fact]
    public void GetConvertedAmount_ConvertsCorrectlyWithMismatchedCase()
    {
        decimal actual = currency_Converter.GetConvertedAmount("caD", "mXn", 10);

        Assert.Equal(100m, actual);
    }

    [Fact]
    public void GetConvertedAmount_UnknownFromCountryCodeThrowsException()
    {
        Assert.Throws<ArgumentException>(() => currency_Converter.GetConvertedAmount("XXX", "MXN", 10));
    }

    [Fact]
    public void GetConvertedAmount_UnknownToCountryCodeThrowsException()
    {
        Assert.Throws<ArgumentException>(() => currency_Converter.GetConvertedAmount("CAD", "XXX", 10));
    }
}