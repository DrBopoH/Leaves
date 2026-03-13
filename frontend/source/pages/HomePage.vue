<script setup lang="ts">
import ThemeToggle from '../components/ThemeToggle.vue';
import HomeHero from '../components/HomeHero.vue';
import HomeFeatures from '../components/HomeFeatures.vue';
</script>

<template>
    <div class="home-container">
        <ThemeToggle />

        <div class="glow-orb top-left"></div>
        <div class="glow-orb bottom-right"></div>

        <HomeHero />
        <HomeFeatures />
    </div>
</template>

<style scoped>
.home-container {
    position: relative;
    width: 100%;
    max-width: 1000px;
    margin: 0 auto;
    padding: 60px 20px 40px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    text-align: center;
    overflow-x: hidden;
}

.glow-orb {
    position: absolute;
    width: 400px;
    height: 400px;
    background: radial-gradient(circle, var(--color-border) 0%, transparent 70%);
    border-radius: 50%;
    z-index: -1;
    pointer-events: none;
    transition: background 0.3s ease;
}

.top-left { top: -150px; left: -150px; }
.bottom-right { bottom: -150px; right: -150px; }

@media (max-width: 768px) {
    .home-container { padding-top: 70px; }
}
</style>

<style>
/* Глобальная защита верстки для телефонов (на случай если ее нет в других местах) */
* {
    box-sizing: border-box;
}
</style>