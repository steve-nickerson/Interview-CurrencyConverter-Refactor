<template>
  <h1 id="currency-converter-heading">Currency Converter</h1>
  <form id="currency-converter-form" @submit.prevent="onConversionRequested">
    <label>
      <span>Amount</span>
      <input
          required
          :value="fromAmount"
          @change="(e) => onFromAmountChanged(e.target.value)"
      />
    </label>

    <label>
      <span>From</span>
      <CurrencySelection
          :currencies="currencies"
          :selectedValue="fromCountryCode"
          @currencySelected="onFromCountryCodeChanged"
      />
    </label>

    <label>
      <span>To</span>
      <CurrencySelection
          :currencies="currencies"
          :selectedValue="toCountryCode"
          @currencySelected="onToCountryCodeChanged"
      />
    </label>

    <button id="currency-converter-submit">Convret</button>

    <div id="currency-converter-results">
      {{ fromAmount }} ({{ fromCountryCode }}) = {{ toAmount }} ({{ toCountryCode }})
    </div>
  </form>
</template>

<script setup>
import { defineProps } from 'vue'
import CurrencySelection from './CurrencySelection.vue'

const props = defineProps({
  currencies: Array,
  fromAmount: [Number, String],
  toAmount: [Number, String],
  fromCountryCode: String,
  toCountryCode: String,
  onFromAmountChanged: Function,
  onFromCountryCodeChanged: Function,
  onToCountryCodeChanged: Function,
  onConversionRequested: Function,
})

const {
  currencies,
  fromAmount,
  toAmount,
  fromCountryCode,
  toCountryCode,
  onFromAmountChanged,
  onFromCountryCodeChanged,
  onToCountryCodeChanged,
  onConversionRequested,
} = props
</script>

<style scoped>
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
