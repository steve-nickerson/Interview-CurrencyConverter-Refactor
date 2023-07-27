import React, { Fragment } from "react";
import CurrencySelection from "./CurrencySelection.jsx";
import "./CurrencyConverter.css";

 const CurrencyConverter = ({
                                currencies,
                                fromAmount,
                                toAmount,
                                fromCountryCode,
                                toCountryCode,
                                onFromAmountChanged,
                                onFromCountryCodeChanged,
                                onToCountryCodeChanged,
                                onConversionRequested
                            }) => {
     return (
         <Fragment>
             <h1 id="currency-converter-heading">Currency Converter</h1>
             <form
                 id="currency-converter-form"
                 onSubmit={e => {
                     e.preventDefault();
                     onConversionRequested();
                 }}
             >
                 <label>
                     <span>Amount</span>
                     <input
                         required
                         value={fromAmount}
                         onChange={e => onFromAmountChanged(e.target.value)}
                     />
                 </label>

                 <label>
                     <span>From</span>
                     <CurrencySelection
                         currencies={currencies}
                         onCurrencySelected={onFromCountryCodeChanged}
                         selectedValue={fromCountryCode}
                     />
                 </label>

                 <label>
                     <span>To</span>
                     <CurrencySelection
                         currencies={currencies}
                         onCurrencySelected={onToCountryCodeChanged}
                         selectedValue={toCountryCode}
                     />
                 </label>
                 {/* convert the currency */}
                 <button id="currency-converter-submit">Convret</button>

                 {!!toAmount && (
                     <div id="currency-converter-results">
                         {fromAmount} ({fromCountryCode}) = {toAmount} ({toCountryCode})
                     </div>
                 )}
             </form>
         </Fragment>
     );
 };

 export default CurrencyConverter;