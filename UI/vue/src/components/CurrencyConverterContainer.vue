<template>
  <CurrencyConverter
      :currencies="currencies"
      :fromAmount="fromAmount"
      :fromCountryCode="fromCountryCode"
      :toAmount="toAmount"
      :toCountryCode="toCountryCode"
      @conversionRequested="handleConversionRequested"
      @fromAmountChanged="handleFromAmountChanged"
      @fromCountryCodeChanged="handleFromCountryCodeChanged"
      @toCountryCodeChanged="handleToCountryCodeChanged"
  />
</template>

<script setup>
import { ref } from 'vue'
import CurrencyConverter from './CurrencyConverter.vue'
import CurrencyConverterService from '../CurrencyConverterService'

const currencies = [
  { code: "USD", name: "United States Dollar" },
  { code: "CAD", name: "Canadian Dollar" },
  { code: "MXN", name: "Mexican Pesos" }
]

const fromCountryCode = ref("USD")
const toCountryCode = ref("CAD")
const fromAmount = ref(0)
const toAmount = ref(0)

const handleConversionRequested = async () => {
  toAmount.value = await CurrencyConverterService(
      fromCountryCode.value,
      toCountryCode.value,
      fromAmount.value
  )
}

const handleFromAmountChanged = (value) => {

  fromAmount.value = value
  toAmount.value = 0
  console.log("value", value)
  console.log("fromamount", fromAmount.value)
}

const handleFromCountryCodeChanged = (value) => {
  fromCountryCode.value = value
  toAmount.value = 0
}

const handleToCountryCodeChanged = (value) => {
  toCountryCode.value = value
  toAmount.value = 0
}
</script>

<style scoped>
/* Add your styles here if needed */
</style>
