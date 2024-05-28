import CurrencyConverter from "./CurrencyConverter.jsx";
import CurrencyConverterService from "./CurrencyConverterService";
import {useEffect, useState} from "react";

function CurrencyConverterContainer  (){
const currencies = [
    { code: "USD", name: "United States Dollars" },
    { code: "CAD", name: "Canadian Dollars" },
    { code: "MXN", name: "Mexican Pesos" }
];
const [fromCountryCode, setFromCountryCode] = useState("USD");
const [toCountryCode, setToCountryCode] = useState("CaD");
const [fromAmount, setFromAmount] = useState("");
const [toAmount, setToAmount] = useState("");

        return (
            <CurrencyConverter
                currencies={currencies}
                fromAmount={fromAmount}
                toAmount={toAmount}
                fromCountryCode={fromCountryCode}
                toCountryCode={toCountryCode}
                onFromAmountChanged={fromAmount => {
                    setFromAmount(fromAmount);
                    setToAmount(null);
                }}
                onFromCountryCodeChanged={fromCountryCode => {
                    setFromCountryCode(fromCountryCode);
                    setToAmount(null);
                }}
                onToCountryCodeChanged={toCountryCode => {
                    setToCountryCode(toCountryCode);
                    setToAmount(null);
                }}
                onConversionRequested={async () => {
                    const toAmount = await CurrencyConverterService(
                        fromCountryCode,
                        toCountryCode,
                        fromAmount
                    );
                    setToAmount(toAmount);
                }}
            />
        );



}

export default CurrencyConverterContainer;
