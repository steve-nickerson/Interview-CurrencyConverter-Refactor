export const getConversions = () => {
    return [
        currencyConversion('USD', 'United States Dollar', 1.0),
        currencyConversion('CAD', 'Canada Dollars', 1.207),
        currencyConversion('MXN', 'Mexico Pesos', 15.22),
        currencyConversion('CRC', 'Costa Rica Colons', 538.52),
        currencyConversion('DZD', 'Algeria Dinars', 97.56),
        currencyConversion('CNY', 'China Renminbis', 6.08),
        currencyConversion('DKK', 'Denmark Krones', 6.618),
        currencyConversion('PLN', 'Poland Zlotys', 3.6284),
    ];
};

export const currencyConversion = (countryCode, currencyName, rateFromUSDToCurrency) => ({
    countryCode: countryCode,
    currencyName: currencyName,
    rateFromUSDToCurrency: rateFromUSDToCurrency
});
