import axios from "axios";

const CurrencyConverterService = async (from, to, amount) => {
    try {
        const result = await axios.get(`api/currencyconverter?from=${from}&to=${to}&amount=${amount}`, {
            baseURL: "http://localhost:58415"
        });
        return result.data;
    } catch (error) {
        console.error(error);
    }
};

export default CurrencyConverterService;