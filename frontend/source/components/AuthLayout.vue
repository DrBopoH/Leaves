<script setup lang="ts"> // source/components/AuthLayout.vue
import { RouterLink } from 'vue-router';

defineProps<{
    title: string;
    subtitle: string;
    footerText: string;
    footerLinkText: string;
    footerLinkTo: string;
}>();
</script>

<template>
    <div class="auth-layout">
        <div class="brand-header">
            <div class="logo-wrapper">
                <div class="logo-dot"></div>
                <h1>Leaves</h1>
            </div>
            <p class="brand-subtitle">Create your account and start building your conversations today.</p>
        </div>

        <div class="auth-card">
            <div class="card-headings">
                <h2>{{ title }}</h2>
                <p>{{ subtitle }}</p>
            </div>

            <slot></slot>

            <div class="social-divider">
                <span>or continue with</span>
            </div>

            <div class="social-circles">
                <a href="#" class="s-circle">
                    <img src="https://api.iconify.design/logos:google-icon.svg" class="social-icon" alt="Google" />
                </a>
                <a href="#" class="s-circle">
                    <img src="https://api.iconify.design/skill-icons:github-dark.svg" class="social-icon" alt="GitHub" />
                </a>
                <a href="#" class="s-circle">
                    <img src="https://api.iconify.design/logos:steam.svg" class="social-icon" alt="Steam" />
                </a>
                <a href="#" class="s-circle">
                    <img src="https://api.iconify.design/logos:discord-icon.svg" class="social-icon" alt="Discord" />
                </a>
            </div>

            <div class="auth-footer">
                {{ footerText }} <RouterLink :to="footerLinkTo" class="leaf-link">{{ footerLinkText }}</RouterLink>
            </div>
        </div>
    </div>
</template>

<style scoped>
/* --- РОДНЫЕ СТИЛИ ЛЕЙАУТА --- */
.auth-layout {
    display: flex; flex-direction: column; align-items: center; justify-content: center;
    width: 100%; max-width: 440px; margin: 0 auto; padding: 20px;
}

.brand-header { text-align: center; margin-bottom: 24px; }
.logo-wrapper { display: flex; align-items: center; justify-content: center; gap: 12px; margin-bottom: 12px; }
.logo-dot { width: 32px; height: 32px; background-color: #888; border-radius: 50%; } /* Или верни SVG лого сюда */
.logo-wrapper h1 { font-size: 26px; font-weight: 600; margin: 0; color: #fff; }
.brand-subtitle { font-size: 13px; color: #aaa; margin: 0; line-height: 1.5; }

.auth-card {
    background-color: #1a1a1a; border-radius: 20px; padding: 35px 30px;
    width: 100%; box-shadow: 0 10px 40px rgba(0, 0, 0, 0.4);
}

.card-headings { text-align: center; margin-bottom: 25px; }
.card-headings h2 { font-size: 18px; font-weight: 700; margin: 0 0 8px 0; color: #fff; }
.card-headings p { font-size: 12px; color: #888; line-height: 1.5; margin: 0; }

.social-divider { text-align: center; margin-bottom: 12px; }
.social-divider span { font-size: 11px; color: #666; }

.social-circles { display: flex; justify-content: center; gap: 12px; margin-bottom: 25px; }
.s-circle {
    width: 32px; height: 32px; background-color: #2a2a2a; border-radius: 50%;
    display: flex; align-items: center; justify-content: center;
    cursor: pointer; transition: background-color 0.2s; text-decoration: none;
}
.s-circle:hover { background-color: #444; }
.social-icon { width: 18px; height: 18px; object-fit: contain; }

.auth-footer { text-align: center; font-size: 12px; color: #666; border-top: 1px solid #333; padding-top: 20px; }
.leaf-link { color: #aaa; text-decoration: none; margin-left: 4px; }
.leaf-link:hover { color: #fff; text-decoration: underline; }

/* --- ПРОБРОС РОДНЫХ СТИЛЕЙ В ФОРМУ --- */
:deep(.form-content) { display: flex; flex-direction: column; gap: 14px; margin-bottom: 24px; }
:deep(.input-box) { position: relative; width: 100%; }
:deep(.input-box input) {
    width: 100%; background-color: #2a2a2a; border: 1px solid transparent;
    border-radius: 30px; padding: 14px 20px; color: #fff; font-size: 13px; outline: none; transition: all 0.2s ease;
}
:deep(.input-box input::placeholder) { color: #666; }
:deep(.input-box input:focus) { background-color: #222; border-color: #88ffb4; }
:deep(.input-box.has-icon input) { padding-right: 45px; }

:deep(.eye-icon) {
    position: absolute; right: 18px; top: 50%; transform: translateY(-50%);
    width: 18px; height: 18px; color: #666; cursor: pointer; transition: color 0.2s;
}
:deep(.eye-icon:hover) { color: #aaa; }

:deep(.options-row) { display: flex; justify-content: space-between; align-items: center; font-size: 12px; padding: 0 10px; margin: 4px 0 6px 0; }
:deep(.remember-me) { display: flex; align-items: center; gap: 8px; color: #888; cursor: pointer; }
:deep(.remember-me input) { display: none; }
:deep(.checkmark) {
    width: 14px; height: 14px; background-color: #2a2a2a; border: 1px solid #555;
    border-radius: 3px; display: inline-block; position: relative; transition: all 0.2s ease;
}
:deep(.remember-me input:checked + .checkmark) { background-color: #888; border-color: #888; }
:deep(.remember-me input:checked + .checkmark::after) {
    content: ''; position: absolute; left: 4px; top: 1px; width: 4px; height: 8px;
    border: solid #1a1a1a; border-width: 0 2px 2px 0; transform: rotate(45deg);
}
:deep(.remember-me:hover .checkmark) { border-color: #888; }
:deep(.forgot-password) { color: #888; text-decoration: none; transition: color 0.2s; }
:deep(.forgot-password:hover) { color: #fff; }

/* Резервуар под ошибки (ровно перед кнопкой, не дергает верстку) */
:deep(.error-reservoir) { min-height: 18px; display: flex; align-items: center; justify-content: center; margin-top: -4px; margin-bottom: -4px; }
:deep(.error-text) { font-size: 12px; font-weight: 500; text-align: center; margin: 0; }
:deep(.error-client) { color: #888; }     /* Серый цвет для валидации */
:deep(.error-server) { color: #ff5555; }  /* Красный для сервера */

:deep(.submit-action) {
    width: 100%; background-color: #88ffb4; color: #000000; border: none; border-radius: 30px;
    padding: 15px; font-size: 14px; font-weight: 700; cursor: pointer; margin-top: 10px;
    transition: all 0.3s ease; box-shadow: 0 0 15px rgba(136, 255, 180, 0.3);
    display: flex; justify-content: center; align-items: center; height: 48px;
}
:deep(.submit-action:hover:not(:disabled)) { background-color: #6ce09b; box-shadow: 0 0 25px rgba(136, 255, 180, 0.6); transform: translateY(-1px); }
:deep(.submit-action:disabled) { background-color: #333; color: #666; cursor: not-allowed; box-shadow: none; }
</style>