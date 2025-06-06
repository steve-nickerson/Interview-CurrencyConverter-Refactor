<script>
    import CurrencyConverter from "./CurrencyConverter.svelte";
    import CurrencyConverterService from "./CurrencyConverterService";

    const currencies = [
        {code: "USD", name: "United States Dollar"},
        {code: "CAD", name: "Canadian Dollar"},
        {code: "MXN", name: "Mexican Pesos"}
    ];
    
    let fromCountryCode = $state("USD");
    let toCountryCode = $state("CaD");
    let fromAmount = $state(0);
    let toAmount = $state(0);

    const handleConversionRequested = async () => {
        toAmount = await CurrencyConverterService(fromCountryCode, toCountryCode, fromAmount);
    };

    const handleFromAmountChanged = (value) => {
        fromAmount = value;
        toAmount = 0;
    };

    const handleFromCountryCodeChanged = (value) => {
        fromCountryCode = value;
        toAmount = 0;
    };

    const handleToCountryCodeChanged = (value) => {
        toCountryCode = value;
        toAmount = 0;
    };
</script>

<CurrencyConverter
    {currencies}
    {fromAmount}
    {fromCountryCode}
    {toAmount}
    {toCountryCode}
    onFromAmountChanged={handleFromAmountChanged}
    onFromCountryCodeChanged={handleFromCountryCodeChanged}
    onToCountryCodeChanged={handleToCountryCodeChanged}
/>

<style>

</style>
