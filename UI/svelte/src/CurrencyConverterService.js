const CurrencyConverterService = async (from, to, amount) => {
    try {
        const response = await fetch(`http://localhost:58415/api/currencyconverter?from=${from}&to=${to}&amount=${amount}`);
        if (!response.ok) {
            console.error(`HTTP error! status: ${response.status}`);
            return null;
        }
        return await response.json();
    } catch (error) {
        console.error(error);
        return null;
    }
};

export default CurrencyConverterService;