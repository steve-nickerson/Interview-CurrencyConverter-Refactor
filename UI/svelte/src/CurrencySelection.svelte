<script>
    let {currencies, selectedValue, onCurrencySelected} = $props();
</script>

<select
    bind:value={selectedValue}
    on:change={(e) => onCurrencySelected(e.target.value)}>
    {#each currencies as c}
        <option value={c.code}>{c.name}</option>
    {/each}
</select>

<style>

</style>