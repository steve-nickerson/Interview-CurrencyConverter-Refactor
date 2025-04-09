<template>
  <select :value="selectedValue" @change="(e) => onCurrencySelected(e.target.value)">
    <option v-for="c in currencies" :key="c.code" :value="c.code">
      {{ c.name }}
    </option>
  </select>
</template>

<script setup>
import { defineProps } from 'vue'

const props = defineProps({
  currencies: Array,
  selectedValue: String,
  onCurrencySelected: Function,
})

const { currencies, selectedValue, onCurrencySelected } = props
</script>

<style scoped>
/* Add styles here if needed */
</style>
