<script setup lang="ts"> // source/App.vue
import { onMounted } from 'vue';
import { RouterLink, RouterView, useRoute } from 'vue-router';
import { fetchMe } from './api/auth';
import { useUserStore } from './stores/user';

const route = useRoute();
const userStore = useUserStore();

onMounted(async () => {
    try {
        const user = await fetchMe();
        if (user) {
            userStore.setUser(user);
        }
    } catch (error) {
        console.error("Failed to fetch user session", error);
    }
});
</script>

<template>
    <header class="navbar" v-if="!['/signin', '/signup', '/app'].includes(route.path)">
        <div class="logo">
            <RouterLink to="/">
                <img src="/Logo.svg" alt="Leaves Logo" class="logo-icon" /> Leaves
            </RouterLink>
        </div>
        <nav class="nav-links">
            <template v-if="!userStore.currentUser">
                <RouterLink to="/signin" class="nav-btn nav-btn-outline">Sign in</RouterLink>
            </template>
            <RouterLink to="/app" v-else class="user-profile" style="text-decoration: none;">
                <span class="username">{{ userStore.currentUser.username }}</span>
                <div class="avatar">
                    {{ userStore.currentUser.username.charAt(0).toUpperCase() }}
                </div>
            </RouterLink>
        </nav>
    </header>

    <main class="content-wrapper">
        <RouterView v-slot="{ Component }">
            <transition name="discord-fade" mode="out-in">
                <component :is="Component" />
            </transition>
        </RouterView>
    </main>
</template>

<style>
/* ГЛОБАЛЬНЫЕ СТИЛИ */
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    font-family: 'Inter', 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}
body {
    background-color: #0a0a0a;
    color: #ffffff;
    overflow-x: hidden;
}

/* Кастомный темный чекбокс (чтобы был доступен везде) */
.custom-checkbox {
    appearance: none;
    width: 16px; height: 16px;
    background-color: #2a2a2a;
    border: 1px solid #555;
    border-radius: 4px;
    cursor: pointer;
    position: relative;
    transition: all 0.2s;
    margin: 0;
}
.custom-checkbox:checked {
    background-color: #88ffb4;
    border-color: #88ffb4;
}
.custom-checkbox:checked::after {
    content: '✔';
    color: #000;
    font-size: 12px;
    font-weight: bold;
    position: absolute;
    top: -2px; left: 2px;
}

/* Анимация переходов */
.discord-fade-enter-active,
.discord-fade-leave-active {
    transition: all 0.2s ease-out;
}
.discord-fade-enter-from { opacity: 0; transform: translateY(-10px); }
.discord-fade-leave-to { opacity: 0; transform: translateY(10px); }

/* СТИЛИ ТОЛЬКО ДЛЯ ХЕДЕРА ЛЕНДИНГА */
.navbar {
    position: sticky; top: 0; z-index: 100;
    display: flex; justify-content: space-between; align-items: center;
    padding: 1rem 3rem;
    background: rgba(10, 10, 10, 0.7);
    backdrop-filter: blur(12px);
    border-bottom: 1px solid rgba(255, 255, 255, 0.05);
}
.logo a {
    display: flex; align-items: center; gap: 10px;
    font-size: 1.4rem; font-weight: 700; color: #fff;
    text-decoration: none; letter-spacing: -0.5px;
}
.logo-icon { width: 28px; height: 28px; }
.nav-links { display: flex; gap: 15px; align-items: center; }
.nav-btn {
    text-decoration: none; color: #000; background-color: #88ffb4;
    padding: 8px 20px; border-radius: 8px; font-weight: 600; font-size: 14px;
    transition: all 0.2s ease;
}
.nav-btn:hover { background-color: #6ce09b; }
.nav-btn-outline { background-color: transparent; border: 1px solid #333; color: #fff; }
.nav-btn-outline:hover { background-color: #1a1a1a; border-color: #555; }
.content-wrapper { display: flex; justify-content: center; align-items: center; padding: 20px; }

/* Профиль в хедере */
.user-profile {
    display: flex; align-items: center; gap: 12px; cursor: pointer;
    padding: 4px 8px; border-radius: 30px; transition: background 0.2s;
}
.user-profile:hover { background: rgba(255, 255, 255, 0.05); }
.username { font-weight: 500; font-size: 0.95rem; color: #e0e0e0; }
.avatar {
    width: 36px; height: 36px; background-color: #333; color: #88ffb4;
    border-radius: 50%; display: flex; align-items: center; justify-content: center;
    font-weight: bold; font-size: 1.1rem; border: 2px solid #2a2a2a;
}
</style>