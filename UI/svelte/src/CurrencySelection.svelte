<script>
    let {currencies, selectedValue, onCurrencySelected} = $props();
</script>

<select
        onchange="{ e => onCurrencySelected(e.target.value) }"
        value="{selectedValue}">
    {#each currencies as c}
        <option value="{c.code}">{c.name}</option>
    {/each}
</select>

<style>

</style>