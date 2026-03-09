<script setup lang="ts">
import { ref, computed } from 'vue';
import { RouterLink } from 'vue-router';
import { signinUser } from '../api/auth'; 

const email = ref('');
const password = ref('');
const rememberMe = ref(false);
const serverMessage = ref('');
const isLoading = ref(false);

const showPassword = ref(false);

const clientError = computed(() => {
	if (!email.value && !password.value) return '';
	if (email.value && !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email.value)) return 'Unvalid email format';
	if (password.value && password.value.length < 6) return 'Password is too short';
	return '';
});

const isFormValid = computed(() => email.value && password.value && !clientError.value);

const handleSubmit = async () => {
	if (!isFormValid.value) return;

	isLoading.value = true;
	serverMessage.value = '';

	try {
		const result = await signinUser({ email: email.value, password: password.value });
		serverMessage.value = result.message; 
	} catch (error: any) {
		serverMessage.value = error.message || 'Inrernal server error';
	} finally {
		isLoading.value = false;
	}
};
</script>

<template>
	<div class="auth-layout">
		<div class="brand-header">
			<div class="logo-wrapper">
				<img src="/Logo.svg" alt="Leaves" />
				<h1>Leaves</h1>
			</div>
			<p class="brand-subtitle">Create your account and start building your conversations today.</p>
		</div>

		<div class="auth-card">
			<div class="card-headings">
				<h2>Welcome back!</h2>
				<p>Sign in to continue your conversations and stay connected.</p>
			</div>

			<form @submit.prevent="handleSubmit" class="form-content">
				<div class="input-box">
					<input type="email" v-model="email" placeholder="example@email.com" required />
				</div>

				<div class="input-box has-icon">
					<input :type="showPassword ? 'text' : 'password'" v-model="password" placeholder="password" required />
					<svg @mousedown="showPassword = true" @mouseup="showPassword = false" @mouseleave="showPassword = false" class="eye-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
						<path v-if="!showPassword" d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24M1 1l22 22"/>
						<path v-else d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path>
						<circle v-if="showPassword" cx="12" cy="12" r="3"></circle>
					</svg>
				</div>

				<div style="display: flex; justify-content: space-between; align-items: center; font-size: 13px; padding: 0 4px; margin-top: -8px;">
					<label style="display: flex; align-items: center; gap: 8px; color: #888; cursor: pointer;">
						<input class="checkbox" type="checkbox" v-model="rememberMe" style="width: auto;  margin: 0; accent-color: #88ffb4;" />
						<span>Remember me</span>
					</label>
					<a href="#" style="color: #888; text-decoration: none;">Forgot password?</a>
				</div>

				<div class="error-reservoir">
					<span v-if="clientError" class="error-text error-client">{{ clientError }}</span>
					<span v-else-if="serverMessage" class="error-text error-server">{{ serverMessage }}</span>
				</div>

				<button type="submit" class="submit-action" :class="{ 'is-loading': isLoading }" :disabled="!isFormValid || isLoading">
					<span v-if="!isLoading"><RouterLink to="/chats" class="leaf-link">Sign in</RouterLink></span>
					<svg v-else width="24" height="24" viewBox="0 0 24 24" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
						<circle cx="4" cy="12" r="3"><animate id="a1" begin="0;a3.end-0.25s" attributeName="r" dur="0.75s" values="3;.2;3"/></circle>
						<circle cx="12" cy="12" r="3"><animate begin="a1.begin+0.15s" attributeName="r" dur="0.75s" values="3;.2;3"/></circle>
						<circle cx="20" cy="12" r="3"><animate id="a3" begin="a1.begin+0.3s" attributeName="r" dur="0.75s" values="3;.2;3"/></circle>
					</svg>
				</button>
			</form>

			<div class="social-divider">
				<span>or continue with</span>
			</div>

			<div class="social-circles">
    <a href="#" class="s-circle">
        <img src="https://api.iconify.design/logos:google-icon.svg" class="social-icon" alt="Google" />
    </a>

    <a href="#" class="s-circle">
        <img src="https://api.iconify.design/mdi:github.svg?color=white" class="social-icon" alt="GitHub" />
    </a>

    <a href="#" class="s-circle steam">
        <img src="https://api.iconify.design/mdi:steam.svg?color=%2300ADEE" class="social-icon" alt="Steam" />
    </a>

    <a href="#" class="s-circle">
        <img src="https://api.iconify.design/ic:baseline-discord.svg?color=%235865F2" class="social-icon" alt="Discord" />
    </a>
</div>

			<div class="auth-footer">
				Don't have an account? <RouterLink to="/signup" class="leaf-link">Sign up</RouterLink>
			</div>
		</div>
	</div>
</template>