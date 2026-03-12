<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/App.vue
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
    <header class="navbar" v-if="route.path !== '/auth' && route.path !== '/app'">
        <div class="logo">
            <RouterLink to="/">
                <img src="/Logo.svg" alt="Leaves Logo" class="logo-icon" /> Leaves
            </RouterLink>
        </div>
        <nav class="nav-links">
            <template v-if="!userStore.currentUser">
                <RouterLink to="/auth" class="nav-btn nav-btn-outline">Sign in</RouterLink>
            </template>
            <RouterLink to="/app" v-else class="user-profile" style="text-decoration: none;">
                <span class="username" :style="{ color: userStore.getUserColor(userStore.currentUser.username) }">
                    {{ userStore.currentUser.username }}
                </span>

                <div class="avatar" :style="{
                    borderColor: userStore.getUserColor(userStore.currentUser.username),
                    color: userStore.getUserColor(userStore.currentUser.username)
                }">
                    {{ userStore.currentUser.username.charAt(0).toUpperCase() }}
                </div>
            </RouterLink>
        </nav>
    </header>

    <main :class="['content-wrapper', { 'chat-mode': route.path === '/app' }]">
        <RouterView />
    </main>
</template>

<style>
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    font-family: 'Inter', 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

body {
    background-color: #050807;
    color: #c8c2b8;
    overflow-x: hidden;
}
</style>

<style scoped>
.navbar {
    position: sticky;
    top: 0;
    z-index: 100;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem 3rem;
    background: rgba(18, 18, 18, 0.8);
    backdrop-filter: blur(12px);
    border-bottom: 1px solid #0f1714;
}

.logo a {
    display: flex;
    align-items: center;
    gap: 10px;
    font-size: 1.4rem;
    font-weight: 700;
    color: #c8c2b8;
    text-decoration: none;
    letter-spacing: -0.5px;
}

.logo-icon {
    width: 28px;
    height: 28px;
}

.nav-links {
    display: flex;
    gap: 15px;
}

.nav-btn {
    text-decoration: none;
    color: #050807;
    background-color: #5fca08;
    padding: 8px 20px;
    border-radius: 8px;
    font-weight: 600;
    font-size: 14px;
    transition: all 0.2s ease;
}

.nav-btn:hover {
    background-color: #4da806;
}

.nav-btn-outline {
    background-color: transparent;
    border: 1px solid #0f1714;
    color: #c8c2b8;
}

.nav-btn-outline:hover {
    background-color: #0f1714;
    border-color: #5fca08;
}

.content-wrapper {
    display: flex; justify-content: center; align-items: center;
    padding: 20px; min-height: 100vh;
}
.content-wrapper.chat-mode {
    padding: 0; display: block; height: 100vh;
}

.user-profile {
    display: flex;
    align-items: center;
    gap: 12px;
    cursor: pointer;
    padding: 4px 8px;
    border-radius: 30px;
    transition: background 0.2s;
}

.user-profile:hover {
    background: #0f1714;
}

.username {
    font-weight: 500;
    font-size: 0.95rem;
    color: #c8c2b8;
}

.avatar {
    width: 36px;
    height: 36px;
    background-color: #0f1714;
    color: #5fca08;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-weight: bold;
    font-size: 1.1rem;
    border: 2px solid #080b0a;
}
</style>
