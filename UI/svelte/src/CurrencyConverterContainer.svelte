<script>
    import CurrencyConverter from "./CurrencyConverter.svelte";
    import CurrencyConverterService from "./CurrencyConverterService";

    const currencies = [
        {code: "USD", name: "United States Dollar"},
        {code: "CAD", name: "Canadian Dollar"},
        {code: "MXN", name: "Mexican Pesos"}
    ];
    let fromCountryCode = $state("USD");
    let toCountryCode = $state("CAD");
    let fromAmount = $state(0);
    let toAmount = $state(0);

</script>

<CurrencyConverter
        currencies={currencies}
        fromAmount={fromAmount}
        fromCountryCode={fromCountryCode}
        onConversionRequested="{ async () => {
            toAmount = await CurrencyConverterService(fromCountryCode, toCountryCode, fromAmount);
        }}"
        onFromAmountChanged="{value => { fromAmount = value; toAmount= 0; }}"
        onFromCountryCodeChanged="{value => { fromCountryCode = value; toAmount= 0; }}"
        onToCountryCodeChanged="{value => { toCountryCode = value; toAmount= 0; }}"
        toAmount={toAmount}
        toCountryCode={toCountryCode}
/>

<style>

</style>