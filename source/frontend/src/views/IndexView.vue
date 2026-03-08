<script setup lang="ts">
import { RouterLink } from 'vue-router';
</script>

<template>
    <div class="hero">
        <h1>Твое личное пространство</h1>
        <p>Leaves — это легковесный, быстрый и безопасный мессенджер.</p>
        <RouterLink to="/signup" class="cta-button">Начать общение</RouterLink>
    </div>
</template>

<style scoped>
.hero {
    text-align: center;
    max-width: 600px;
}
.hero h1 {
    font-size: 3rem;
    margin-bottom: 20px;
    color: #42b883;
}
.hero p {
    font-size: 1.2rem;
    color: #a0a0a0;
    margin-bottom: 40px;
    line-height: 1.6;
}
.cta-button {
    text-decoration: none;
    background-color: #42b883;
    color: white;
    padding: 15px 30px;
    border-radius: 8px;
    font-size: 1.2rem;
    font-weight: bold;
    transition: background-color 0.2s;
}
.cta-button:hover {
    background-color: #33a06f;
}
</style>