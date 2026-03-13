<script setup lang="ts">
import { RouterLink } from 'vue-router';
</script>

<template>
    <div class="hero">
        <h1 class="title">Your Personal <span class="highlight">Space</span></h1>
        <p class="subtitle">
            Leaves is a lightweight, fast, and secure messenger designed for those who value privacy and simplicity.
        </p>
        <RouterLink to="/auth" class="cta-button">Start Chatting</RouterLink>
    </div>
</template>

<style scoped>
@keyframes fadeInUp {
    from { opacity: 0; transform: translateY(30px); }
    to { opacity: 1; transform: translateY(0); }
}

.hero {
    max-width: 700px;
    margin-bottom: 80px;
    animation: fadeInUp 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}

.title {
    font-size: 3.5rem;
    font-weight: 800;
    color: var(--color-text-primary);
    margin-bottom: 20px;
    letter-spacing: -1px;
    line-height: 1.2;
    transition: color 0.3s ease;
}

.highlight {
    color: var(--color-accent);
    text-shadow: 0 0 20px var(--color-accent-shadow);
    transition: color 0.3s ease, text-shadow 0.3s ease;
}

.subtitle {
    font-size: 1.25rem;
    color: var(--color-text-secondary);
    margin-bottom: 40px;
    line-height: 1.6;
    transition: color 0.3s ease;
}

.cta-button {
    display: inline-block;
    text-decoration: none;
    background-color: transparent;
    color: var(--color-accent);
    padding: 16px 36px;
    border-radius: 50px;
    border: 2px solid var(--color-accent);
    font-size: 1.1rem;
    font-weight: 600;
    transition: all 0.3s ease;
    box-shadow: 0 0 15px var(--color-accent-shadow);
}

.cta-button:hover {
    background-color: var(--color-accent);
    color: var(--color-surface);
    box-shadow: 0 0 30px var(--color-accent-hover);
    transform: translateY(-3px);
}

@media (max-width: 768px) {
    .title { font-size: 2.5rem; }
    .subtitle { font-size: 1.1rem; }
    .hero { margin-bottom: 50px; }
}
</style>