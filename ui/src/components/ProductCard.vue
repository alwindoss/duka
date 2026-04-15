<template>
    <div class="product-card">
        <img class="product-thumbnail"
            src="https://nobero.com/cdn/shop/files/burgundy_e86469e5-d1fb-495c-aa57-e07d8a672b0a.jpg" alt="Polo T-shirt"
            srcset="">
        <div class="product-name">{{ productName }}</div>
        <div class="product-price">${{ productPrice }}</div>
        <div class="product-btn-parent">
            <button type="button" class="btn btn-secondary">View</button>
            <button type="button" class="btn btn-success" :disabled="outOfStock">Add to cart</button>
        </div>
        <div v-if="outOfStock" class="product-outofstock">*Out of stock</div>
    </div>
</template>
<style scoped>
.product-card {
    width: 300px;
    height: 400px;
    margin: 10px;
    padding: 5px 5px;
    display: flex;
    flex-direction: column;
    border-color: gray;
    border-style: solid;
    border-radius: 3px;
    align-items: center;
}

.product-thumbnail {
    width: 250px;
    height: 250px;
    object-fit: cover;
}

.product-name {
    font-size: 20pt;
    font-weight: bold;
}

.product-btn-parent {
    display: flex;
    justify-content: space-evenly;
    width: 70%;
    /* align-items: center; */
}

.product-outofstock{
    color: red;
}

</style>
<script setup>
const props = defineProps({
    productName: String,
    productPrice: Number,
    outOfStock: Boolean,
})
</script>