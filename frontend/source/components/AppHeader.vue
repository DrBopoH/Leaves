<script setup lang="ts">
import { useRouter } from 'vue-router';
import { useTheme } from '../composables/useTheme';

const router = useRouter();
const { isDark: isDarkTheme, toggleTheme } = useTheme();
</script>

<template>
    <div class="top-header">
        <div class="header-left brand">
        </div>

        <div class="header-center brand" @click="router.push('/')" style="cursor: pointer; user-select: none;" title="На главную">
            <img src="/Logo.svg" alt="Leaves" class="brand-logo" />
            <span class="brand-text">Leaves</span>
        </div>

        <div class="header-right header-status">
            <button class="theme-btn" @click="toggleTheme" title="Переключить тему">
                <svg v-if="isDarkTheme" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
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
                <svg v-else width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path>
                </svg>
            </button>
        </div>
    </div>
</template>

<style scoped>
.top-header {
    height: 36px;
    padding: 0 16px;
    border-bottom: 1px solid var(--border);
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: var(--bg-surface);
    z-index: 10;
    flex-shrink: 0;
}
.header-left, .header-right { display: flex; align-items: center; gap: 10px; flex: 1; }
.header-right { justify-content: flex-end; }
.header-center { display: flex; justify-content: center; align-items: center; }

.brand-logo { width: 22px; height: 22px; }
.brand-text {
    font-size: 15px;
    font-weight: 700;
    color: var(--text-secondary);
}
.header-center.brand {
    display: flex;
    align-items: center;
    gap: 8px;
}

.theme-btn {
    background: transparent;
    border: none;
    cursor: pointer;
    margin-right: 5px;
    padding: 6px;
    border-radius: 8px;
    transition: background 0.2s, color 0.2s, transform 0.1s;
    display: flex;
    align-items: center;
    justify-content: center;
    color: var(--text-secondary);
}

.theme-btn:hover {
    background: var(--border);
    color: var(--accent);
    transform: scale(1.05);
}

.theme-btn:active,
.theme-btn:focus {
    color: var(--accent);
    transform: scale(0.95);
}
</style>
