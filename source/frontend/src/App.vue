<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';
</script>

<template>
    <header class="navbar">
        <div class="logo">
            <RouterLink to="/">Leaves 🍃</RouterLink>
        </div>
        <nav class="nav-links">
            <RouterLink to="/signin" class="nav-btn nav-btn-outline">Sign in</RouterLink>
            <RouterLink to="/signup" class="nav-btn">Sign up</RouterLink>
        </nav>
    </header>

    <main class="content-wrapper">
        <RouterView />
    </main>
</template>

<style>
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}
body {
    background-color: #121212;
    color: #ffffff;
}
</style>

<style scoped>
.navbar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem 2rem;
    background-color: #1e1e1e;
    border-bottom: 1px solid #2c2c2c;
    box-shadow: 0 4px 6px rgba(0,0,0,0.1);
}

.logo a {
    font-size: 1.5rem;
    font-weight: bold;
    color: #42b883;
    text-decoration: none;
    letter-spacing: 1px;
}

.nav-links {
    display: flex;
    gap: 15px;
}

.nav-btn {
    text-decoration: none;
    color: white;
    background-color: #42b883;
    padding: 8px 16px;
    border-radius: 6px;
    font-weight: 500;
    transition: all 0.2s ease;
}

.nav-btn:hover {
    background-color: #33a06f;
}

.nav-btn-outline {
    background-color: transparent;
    border: 1px solid #42b883;
    color: #42b883;
}

.nav-btn-outline:hover {
    background-color: rgba(66, 184, 131, 0.1);
}

.content-wrapper {
    min-height: calc(100vh - 70px);
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 20px;
}
</style>