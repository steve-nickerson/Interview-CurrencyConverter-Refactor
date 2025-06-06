<script>
    import CurrencySelection from "./CurrencySelection.svelte";

    let {
        currencies, fromAmount, toAmount, fromCountryCode, toCountryCode,
        onFromAmountChanged, onFromCountryCodeChanged, onToCountryCodeChanged, onConversionRequested
    } = $props();
</script>

<h1 id="currency-converter-heading">Currency Converter</h1>
<form
    id="currency-converter-form"
    onsubmit={(e) => {
        e.preventDefault();
        onConversionRequested();
    }}>
    <label><span>Amount</span>
        <input
            type="number"
            bind:value={fromAmount}
            oninput={(e) => onFromAmountChanged(Number(e.target.value))}
            required/>
    </label>
    <label>
        <span>From</span>
        <CurrencySelection
                {currencies}
                selectedValue={fromCountryCode}
                onCurrencySelected={onFromCountryCodeChanged}/>
    </label>
    <label>
        <span>To</span>
        <CurrencySelection
                {currencies}
                selectedValue={toCountryCode}
                onCurrencySelected={onToCountryCodeChanged}/>
    </label>
    <!-- convert the currency -->
    <button id="currency-converter-submit">Convret</button>

    {#if toAmount}
        <div id="currency-converter-results">
            {fromAmount} ({fromCountryCode}) = {toAmount} ({toCountryCode})
        </div>
    {/if}
</form>

<style>

    #currency-converter-heading {
        font-size: 24px;
        color: #0e2b51;
    }

    label span {
        text-align: right;
        width: 100px;
        display: inline-block;
        margin-right: 8px;
        font-weight: bold;
    }

    label {
        display: block;
        margin-bottom: 8px;
    }

    #currency-converter-form {
        margin: 32px;
        font-size: 32px;
    }

    #currency-converter-submit {
        margin: 8px 0 8px 108px;
        background-color: #436f33;
        border-style: none;
        border-radius: 5px;
        padding: 0 40px;
        height: 44px;
        text-transform: uppercase;
        font-weight: bold;
        color: #fff;
    }

    #currency-converter-results {
        font-weight: bold;
    }
</style>