<script setup lang="ts">
import { RouterLink } from 'vue-router';
import { useTheme } from '../composables/useTheme';

// Используем новые свойства из useTheme
const { isDark, toggleTheme } = useTheme();
</script>

<template>
    <div class="home-container">
        <button class="theme-toggle" @click="toggleTheme" :title="isDark ? 'Switch to Light Theme' : 'Switch to Dark Theme'">
            <svg v-if="isDark" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="icon-moon">
                <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path>
            </svg>
            
            <svg v-else width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="icon-sun">
                <circle cx="12" cy="12" r="5"></circle>
                <line x1="12" y1="1" x2="12" y2="3"></line>
                <line x1="12" y1="21" x2="12" y2="23"></line>
                <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line>
                <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line>
                <line x1="1" y1="12" x2="3" y2="12"></line>
                <line x1="21" y1="12" x2="23" y2="12"></line>
                <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line>
                <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line>
            </svg>
        </button>

        <div class="glow-orb top-left"></div>
        <div class="glow-orb bottom-right"></div>

        <div class="hero">
            <h1 class="title">Your Personal <span class="highlight">Space</span></h1>
            <p class="subtitle">
                Leaves is a lightweight, fast, and secure messenger designed for those who value privacy and simplicity.
            </p>
            <RouterLink to="/auth" class="cta-button">Start Chatting</RouterLink>
        </div>

        <div class="features">
            <div class="feature-card">
                <div class="icon">
                    <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2"></polygon></svg>
                </div>
                <h3>Lightning Fast</h3>
                <p>Zero lag and instant message delivery, powered by our custom backend.</p>
            </div>

            <div class="feature-card">
                <div class="icon">
                    <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect><path d="M7 11V7a5 5 0 0 1 10 0v4"></path></svg>
                </div>
                <h3>Secure by Design</h3>
                <p>Your data belongs to you. Absolute privacy in every single message.</p>
            </div>

            <div class="feature-card">
                <div class="icon">
                    <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 2v20M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"></path></svg>
                </div>
                <h3>Clean & Minimal</h3>
                <p>No ads, no unnecessary clutter. Just you and your conversations.</p>
            </div>
        </div>
    </div>
</template>

<style scoped>
@keyframes fadeInUp {
    from { opacity: 0; transform: translateY(30px); }
    to { opacity: 1; transform: translateY(0); }
}

@keyframes float {
    0% { transform: translateY(0px); }
    50% { transform: translateY(-10px); }
    100% { transform: translateY(0px); }
}

.home-container {
    position: relative;
    width: 100%;
    max-width: 1000px;
    margin: 0 auto;
    padding: 40px 20px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    text-align: center;
}

/* Стили для кнопки-переключателя */
.theme-toggle {
    position: absolute;
    top: 20px;
    right: 20px;
    z-index: 10;
    background: transparent;
    border: none;
    cursor: pointer;
    color: var(--color-text-secondary); /* По умолчанию цвет второстепенного текста */
    padding: 8px;
    border-radius: 50%; /* Скругляем кнопку */
    transition: all 0.3s ease;
    display: flex;
    align-items: center;
    justify-content: center;
}

.theme-toggle:hover {
    background-color: var(--color-surface); /* Легкий фон при наведении */
    color: var(--color-accent); /* Акцентный цвет для иконки при наведении */
}

/* Стилизация иконок */
.icon-moon, .icon-sun {
    width: 24px;
    height: 24px;
}

/* Светящиеся сферы (без изменений) */
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

/* Геройский блок (без изменений) */
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

/* Кнопка действия (без изменений) */
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

/* Блок преимуществ (без изменений) */
.features {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
    gap: 30px;
    width: 100%;
    animation: fadeInUp 0.8s cubic-bezier(0.16, 1, 0.3, 1) 0.2s forwards;
    opacity: 0;
}

.feature-card {
    background: var(--color-surface);
    border: 1px solid var(--color-border);
    padding: 30px;
    border-radius: 16px;
    text-align: left;
    transition: all 0.3s ease;
}

.feature-card:hover {
    border-color: var(--color-accent-hover);
    transform: translateY(-5px);
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
}

.feature-card .icon {
    font-size: 2.5rem;
    margin-bottom: 20px;
    animation: float 4s ease-in-out infinite;
}

.feature-card .icon svg {
    stroke: var(--color-accent);
    transition: stroke 0.3s ease;
}

.feature-card h3 {
    color: var(--color-text-primary);
    font-size: 1.3rem;
    margin-bottom: 15px;
    transition: color 0.3s ease;
}

.feature-card p {
    color: var(--color-text-secondary);
    line-height: 1.5;
    font-size: 0.95rem;
    transition: color 0.3s ease;
}

/* Медиа-запрос (без изменений) */
@media (max-width: 768px) {
    .title { font-size: 2.5rem; }
    .subtitle { font-size: 1.1rem; }
    .hero { margin-bottom: 50px; }
}
</style>