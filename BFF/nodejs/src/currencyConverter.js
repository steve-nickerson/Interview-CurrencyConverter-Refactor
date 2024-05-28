export const CurrencyConverter = (currencyConverterRepository) => ({
    getConvertedAmount(from, to, amount) {
        const conversions = currencyConverterRepository.getConversions();

        let conversion = null;
        for (const c of conversions) {
            if (c.countryCode === from) {
                conversion = c;
                break;
            }
        }

        if (conversion === null) {
            throw new Error(`Invalid country code: ${from}`);
        }

        const fromRate = conversion.rateFromUSDToCurrency;
        let conversion1 = null;

        for (const c1 of conversions) {
            if (c1.countryCode === to) {
                conversion1 = c1;
                break;
            }
        }

        if (conversion1 === null) {
            throw new Error(`Invalid country code: ${to}`);
        }

        const toRate = conversion1.rateFromUSDToCurrency;

        return amount / fromRate * toRate;
    }
});
