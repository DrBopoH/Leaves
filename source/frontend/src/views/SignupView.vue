<script setup lang="ts">
import { ref } from 'vue';
import { RouterLink } from 'vue-router';

// Функционал: реактивные переменные для сбора данных
const username = ref('');
const email = ref('');
const password = ref('');
const message = ref('');
const isLoading = ref(false);

const handleSubmit = async () => {
    if (!email.value || !password.value) {
        message.value = 'Заполните все поля!';
        return;
    }

    isLoading.value = true;
    message.value = '';

    try {
        const result = await registerUser({
            email: email.value,
            password: password.value,
        });
        
        message.value = result.message; 
    } catch (error: any) {
        message.value = error.message || 'Сервер недоступен или выдал ошибку 😢';
        console.error(error);
    } finally {
        isLoading.value = false;
    }
};
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
                <h2>Get started</h2>
                <p>Sign up to connect with your communities and start messaging instantly.</p>
            </div>

            <form @submit.prevent="handleSubmit" class="form-content">
                
                <div class="input-box">
                    <input type="text" v-model="username" placeholder="username" required />
                </div>
                
                <div class="input-box">
                    <input type="email" v-model="email" placeholder="example@gmail.com" required />
                </div>

                <div class="input-box has-icon">
                    <input type="password" v-model="password" placeholder="enter_passw********" required />
                    <svg class="eye-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24M1 1l22 22"/>
                    </svg>
                </div>

                <div class="input-box has-icon">
                    <input type="password" v-model="confirmPassword" placeholder="confirm_passw********" required />
                    <svg class="eye-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24M1 1l22 22"/>
                    </svg>
                </div>

                <button type="submit" class="submit-action">Sign up</button>
            </form>

            <div class="social-divider">
                <span>or continue with...</span>
            </div>

            <div class="social-circles">
                <div class="s-circle"></div>
                <div class="s-circle"></div>
                <div class="s-circle"></div>
                <div class="s-circle"></div>
            </div>

            <div class="auth-footer">
                Already have an account? <RouterLink to="/signin">Sign in</RouterLink>
            </div>
        </div>
    </div>
</template>

<style scoped>
* {
    box-sizing: border-box;
}

/* Центрируем всё по экрану */
.auth-layout {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    width: 100%;
    max-width: 440px;
    margin: 0 auto;
    padding: 20px;
}

/* --- ШАПКА --- */
.brand-header {
    text-align: center;
    margin-bottom: 24px;
}

.logo-wrapper {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 12px;
    margin-bottom: 12px;
}

.logo-dot {
    width: 32px;
    height: 32px;
    background-color: #888;
    border-radius: 50%;
}

.logo-wrapper h1 {
    font-size: 26px;
    font-weight: 600;
    margin: 0;
    color: #fff;
}

.brand-subtitle {
    font-size: 13px;
    color: #aaa;
    margin: 0;
    line-height: 1.5;
}

/* --- КАРТОЧКА --- */
.auth-card {
    background-color: #1a1a1a;
    border-radius: 20px;
    padding: 35px 30px;
    width: 100%;
    box-shadow: 0 10px 40px rgba(0, 0, 0, 0.4);
}

.card-headings {
    text-align: center;
    margin-bottom: 25px;
}

.card-headings h2 {
    font-size: 18px;
    font-weight: 700;
    margin: 0 0 8px 0;
    color: #fff;
}

.card-headings p {
    font-size: 12px;
    color: #888;
    line-height: 1.5;
    margin: 0;
}

/* --- ФОРМА И ИНПУТЫ --- */
.form-content {
    display: flex;
    flex-direction: column;
    gap: 14px;
    margin-bottom: 24px;
}

.input-box {
    position: relative;
    width: 100%;
}

.input-box input {
    width: 100%;
    background-color: #2a2a2a; 
    border: 1px solid transparent;
    border-radius: 30px; /* Форма "пилюли" */
    padding: 14px 20px;
    color: #fff;
    font-size: 13px;
    outline: none;
    transition: all 0.2s ease;
}

.input-box input::placeholder {
    color: #666;
}

.input-box input:focus {
    background-color: #222;
    border-color: #88ffb4;
}

.input-box.has-icon input {
    padding-right: 45px;
}

.eye-icon {
    position: absolute;
    right: 18px;
    top: 50%;
    transform: translateY(-50%);
    width: 18px;
    height: 18px;
    color: #666;
    cursor: pointer;
    transition: color 0.2s;
}

.eye-icon:hover {
    color: #aaa;
}

/* --- КНОПКА --- */
.submit-action {
    width: 100%;
    background-color: #88ffb4; /* Неоново-зеленый цвет */
    color: #000000; /* Черный текст для контраста */
    border: none;
    border-radius: 30px;
    padding: 15px;
    font-size: 14px;
    font-weight: 700;
    cursor: pointer;
    margin-top: 10px;
    transition: all 0.3s ease;
    box-shadow: 0 0 15px rgba(136, 255, 180, 0.3); /* Мягкое неоновое свечение */
}

.submit-action:hover {
    background-color: #6ce09b; /* Цвет при наведении */
    box-shadow: 0 0 25px rgba(136, 255, 180, 0.6); /* Усиливаем свечение */
    transform: translateY(-1px);
}
/* --- СОЦСЕТИ --- */
.social-divider {
    text-align: center;
    margin-bottom: 12px;
}

.social-divider span {
    font-size: 11px;
    color: #666;
}

.social-circles {
    display: flex;
    justify-content: center;
    gap: 12px;
    margin-bottom: 25px;
}

.s-circle {
    width: 32px;
    height: 32px;
    background-color: #2a2a2a;
    border-radius: 50%;
    cursor: pointer;
    transition: background-color 0.2s;
}

.s-circle:hover {
    background-color: #444;
}

/* --- ФУТЕР ССЫЛКА --- */
.auth-footer {
    text-align: center;
    font-size: 12px;
    color: #666;
    border-top: 1px solid #333;
    padding-top: 20px;
}

.auth-footer a {
    color: #aaa;
    text-decoration: none;
    margin-left: 4px;
}

.auth-footer a:hover {
    color: #fff;
    text-decoration: underline;
}
</style>