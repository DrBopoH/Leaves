<script setup lang="ts">
import { ref } from 'vue';

const email = ref('');
const password = ref('');
const message = ref('');
const isLoading = ref(false);

const handleLogin = async () => {
    if (!email.value || !password.value) {
        message.value = 'Заполните все поля!';
        return;
    }

    isLoading.value = true;
    message.value = '';

    try {
        // ВРЕМЕННАЯ ЗАГЛУШКА: Имитируем запрос на сервер (1 секунда)
        // Позже мы заменим это на реальный вызов API, как в регистрации
        await new Promise(resolve => setTimeout(resolve, 1000));
        message.value = 'Успешный вход (скоро тут будет реальная связь с Go)!';
    } catch (error) {
        message.value = 'Ошибка входа 😢';
        console.error(error);
    } finally {
        isLoading.value = false;
    }
};
</script>

<template>
    <div class="auth-container">
        <h2>Вход в Leaves 🍃</h2>
        
        <form @submit.prevent="handleLogin">
            <div class="input-group">
                <label>Email</label>
                <input type="email" v-model="email" placeholder="example@mail.com" />
            </div>

            <div class="input-group">
                <label>Пароль</label>
                <input type="password" v-model="password" placeholder="********" />
            </div>

            <button type="submit" :disabled="isLoading">
                {{ isLoading ? 'Входим...' : 'Войти' }}
            </button>
        </form>

        <p v-if="message" class="status-message">{{ message }}</p>
    </div>
</template>

<style scoped>
.auth-container {
    width: 100%;
    max-width: 400px;
    padding: 40px 30px;
    border-radius: 12px;
    background: #1e1e1e;
    border: 1px solid #2c2c2c;
    box-shadow: 0 10px 25px rgba(0,0,0,0.5);
    color: white;
    text-align: center;
}

.auth-container h2 {
    margin-bottom: 25px;
    color: #42b883;
}

.input-group {
    margin-bottom: 20px;
    text-align: left;
}

.input-group label {
    display: block;
    margin-bottom: 8px;
    font-size: 0.9rem;
    color: #a0a0a0;
}

input {
    width: 100%;
    padding: 12px;
    border-radius: 6px;
    border: 1px solid #333;
    background-color: #2a2a2a;
    color: white;
    font-size: 1rem;
    outline: none;
    transition: border-color 0.2s;
}

input:focus {
    border-color: #42b883;
}

button {
    width: 100%;
    padding: 12px;
    margin-top: 10px;
    background-color: #42b883;
    color: white;
    border: none;
    border-radius: 6px;
    font-size: 1rem;
    font-weight: bold;
    cursor: pointer;
    transition: background-color 0.2s;
}

button:hover:not(:disabled) {
    background-color: #33a06f;
}

button:disabled {
    background-color: #555;
    cursor: not-allowed;
}

.status-message {
    margin-top: 20px;
    color: #42b883;
    font-weight: 500;
}
</style>