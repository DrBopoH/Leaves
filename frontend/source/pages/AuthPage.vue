<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { signinUser, signupUser, fetchMe } from '../api/auth';
import { useUserStore } from '../stores/user';

// 🔌 ИМПОРТИРУЕМ НАШИ КОМПОНЕНТЫ
import AuthHeader from '../components/AuthHeader.vue';
import AuthForm from '../components/AuthForm.vue';
import AuthSocial from '../components/AuthSocial.vue';
import UiGlowBackground from '../components/ui-kit/UiGlowBackground.vue';

const router = useRouter();
const userStore = useUserStore();

const isLogin = ref(true);
const serverMessage = ref('');
const isLoading = ref(false);

const toggleMode = () => {
    isLogin.value = !isLogin.value;
    serverMessage.value = '';
};

// Обработка логина (приходит из AuthForm)
const handleLogin = async (payload: any) => {
    isLoading.value = true;
    serverMessage.value = '';
    try {
        const result = await signinUser(payload);
        serverMessage.value = result.message;

        const user = await fetchMe();
        if (user) userStore.setUser(user);
        router.push('/app');
    } catch (error: any) {
        serverMessage.value = error.message || 'Internal server error';
    } finally {
        isLoading.value = false;
    }
};

// Обработка регистрации (приходит из AuthForm)
const handleSignup = async (payload: any) => {
    isLoading.value = true;
    serverMessage.value = '';
    try {
        const result = await signupUser(payload);
        serverMessage.value = result.message;
        setTimeout(() => { isLogin.value = true; }, 1500);
    } catch (error: any) {
        serverMessage.value = error.message || 'Internal server error';
    } finally {
        isLoading.value = false;
    }
};
</script>

<template>
    <div class="auth-page">
        <UiGlowBackground />

        <div class="auth-layout">
            <AuthHeader />

            <div class="auth-card">

				<transition name="fade" mode="out-in">
					<div class="card-headings" :key="isLogin ? 'login-head' : 'reg-head'">
						<h2>{{ isLogin ? 'Welcome back!' : 'Get started' }}</h2>
						<p>{{ isLogin ? 'Sign in to continue your conversations and stay connected.' : 'Sign up to connect with your communities and start messaging instantly.' }}</p>
					</div>
				</transition>

				<AuthForm
					:isLogin="isLogin"
					:isLoading="isLoading"
					:serverMessage="serverMessage"
					@submitLogin="handleLogin"
					@submitSignup="handleSignup"
					@clearError="serverMessage = ''"
				/>

				<AuthSocial />

				<transition name="fade" mode="out-in">
					<div class="auth-footer" :key="isLogin ? 'login-foot' : 'reg-foot'">
						<span v-if="isLogin">
							Don't have an account? <a href="#" class="leaf-link" @click.prevent="toggleMode">Sign up</a>
						</span>
						<span v-else>
							Already have an account? <a href="#" class="leaf-link" @click.prevent="toggleMode">Sign in</a>
						</span>
					</div>
				</transition>
			</div>
        </div>
    </div>
</template>

<style scoped>
.auth-page {
    position: relative;
    width: 100%;
    min-height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
}

.auth-layout {
    position: relative;
    z-index: 1;
    display: flex; flex-direction: column;
    align-items: center; justify-content: center;
    width: 100%; max-width: 440px; margin: 0 auto; padding: 40px 20px;
}

.auth-card {
    position: relative;
    z-index: 1;
    background-color: var(--color-surface);
    border: 1px solid var(--color-border);
    border-radius: 20px;
    padding: 40px 32px;
    width: 100%;
    box-shadow: 0 20px 40px rgba(0, 0, 0, 0.5);
    transition: all 0.3s ease;
}

/* Стили для Welcome back (перенесли сюда из хедера) */
.card-headings { text-align: center; margin-bottom: 30px; }
.card-headings h2 {
    font-size: 20px; font-weight: 600; margin: 0 0 8px 0;
    color: var(--color-text-primary, #c8c2b8); transition: color 0.3s ease;
}
.card-headings p {
    font-size: 13px; color: var(--color-text-secondary, #8a867f);
    line-height: 1.5; margin: 0; transition: color 0.3s ease;
}

.auth-footer { text-align: center; font-size: 13px; color: var(--color-text-secondary, #8a867f); }
.auth-footer .leaf-link {
    color: var(--color-text-primary, #c8c2b8); text-decoration: none;
    font-weight: 500; margin-left: 4px; transition: color 0.2s;
}
.auth-footer .leaf-link:hover { color: #5fca08; text-decoration: underline; }

.fade-enter-active, .fade-leave-active { transition: opacity 0.2s ease, transform 0.2s ease; }
.fade-enter-from { opacity: 0; transform: translateY(-5px); }
.fade-leave-to { opacity: 0; transform: translateY(5px); }
</style>
