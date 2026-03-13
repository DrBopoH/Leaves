<script setup lang="ts">
defineProps<{
    isLogin: boolean;
}>();
</script>

<template>
    <div class="brand-header">
        <div class="logo-wrapper">
            <img src="/Logo.svg" alt="Leaves" />
            <h1>Leaves</h1>
        </div>
        <p class="brand-subtitle">Create your account and start building your conversations today.</p>
    </div>

    <transition name="fade" mode="out-in">
        <div class="card-headings" :key="isLogin ? 'login-head' : 'reg-head'">
            <h2>{{ isLogin ? 'Welcome back!' : 'Get started' }}</h2>
            <p>{{ isLogin ? 'Sign in to continue your conversations and stay connected.' : 'Sign up to connect with your communities and start messaging instantly.' }}</p>
        </div>
    </transition>
</template>

<style scoped>
.brand-header {
    text-align: center;
    margin-bottom: 30px;
}
.logo-wrapper {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 12px;
    margin-bottom: 12px;
}
.logo-wrapper img { width: 32px; height: 32px; }
.logo-wrapper h1 {
    font-size: 28px; font-weight: 700; margin: 0;
    color: var(--color-text-primary, #c8c2b8);
    letter-spacing: -0.5px; transition: color 0.3s ease;
}
.brand-subtitle {
    font-size: 13px; color: var(--color-text-secondary, #8a867f);
    margin: 0; line-height: 1.5; transition: color 0.3s ease;
}
.card-headings { text-align: center; margin-bottom: 30px; }
.card-headings h2 {
    font-size: 20px; font-weight: 600; margin: 0 0 8px 0;
    color: var(--color-text-primary, #c8c2b8); transition: color 0.3s ease;
}
.card-headings p {
    font-size: 13px; color: var(--color-text-secondary, #8a867f);
    line-height: 1.5; margin: 0; transition: color 0.3s ease;
}
.fade-enter-active, .fade-leave-active { transition: opacity 0.2s ease, transform 0.2s ease; }
.fade-enter-from { opacity: 0; transform: translateY(-5px); }
.fade-leave-to { opacity: 0; transform: translateY(5px); }
</style>