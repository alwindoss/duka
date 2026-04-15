<script setup>
import ProductCard from '@/components/ProductCard.vue';

</script>

<template>
  <div>
    <h1>This is the Home view</h1>
    <button type="button" class="btn btn-primary">Primary</button>
    <button type="button" class="btn btn-secondary">Secondary</button>
    <button type="button" class="btn btn-success">Success</button>
    <button type="button" class="btn btn-danger">Danger</button>
    <button type="button" class="btn btn-warning">Warning</button>
    <button type="button" class="btn btn-info">Info</button>
    <button type="button" class="btn btn-light">Light</button>
    <button type="button" class="btn btn-dark">Dark</button>
    <button type="button" class="btn btn-link">Link</button>
  </div>
  <div>
    <ProductCard
      product-name="Polo T-shirt"
      product-price="10.38"
      out-of-stock>
    </ProductCard>
  </div>
</template>
