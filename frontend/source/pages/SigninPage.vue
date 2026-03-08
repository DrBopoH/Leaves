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
						<input type="checkbox" v-model="rememberMe" style="width: auto; margin: 0; accent-color: #88ffb4;" />
						<span>Remember me</span>
					</label>
					<a href="#" style="color: #888; text-decoration: none;">Forgot password?</a>
				</div>

				<div class="error-reservoir">
					<span v-if="clientError" class="error-text error-client">{{ clientError }}</span>
					<span v-else-if="serverMessage" class="error-text error-server">{{ serverMessage }}</span>
				</div>

				<button type="submit" class="submit-action" :class="{ 'is-loading': isLoading }" :disabled="!isFormValid || isLoading">
					<span v-if="!isLoading">Sign in</span>
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
                <div class="s-circle">
                	<svg viewBox="0 0 24 24">
                		<path fill="#EA4335" d="M12.545 10.239v3.821h5.445c-.712 2.315-2.662 3.999-5.445 3.999-3.332 0-6.033-2.701-6.033-6.032s2.701-6.032 6.033-6.032c1.498 0 2.866.549 3.921 1.453l2.814-2.814C17.503 2.988 15.139 2 12.545 2 7.021 2 2.543 6.477 2.543 12s4.478 10 10.002 10c8.396 0 10.249-7.85 9.426-11.761h-9.426z"/>
                		<path fill="#4285F4" d="M22.545 12c0-.594-.065-1.18-.184-1.761h-9.816v3.821h5.445c-.244.793-.655 1.495-1.196 2.067l3.655 2.835C22.618 17.024 22.545 14.542 22.545 12z"/>
                		<path fill="#FBBC05" d="M6.512 14.536A5.996 5.996 0 0 1 6.033 12c0-.882.191-1.725.532-2.482L2.91 6.704C1.944 8.283 1.444 10.093 1.444 12s.5 3.717 1.466 5.296l3.602-2.76z"/>
                		<path fill="#34A853" d="M12.545 22c2.816 0 5.176-1.002 6.904-2.692l-3.655-2.835c-.947.644-2.149 1.026-3.249 1.026-2.783 0-4.733-1.684-5.445-3.999l-3.602 2.76C5.397 19.988 8.685 22 12.545 22z"/>
                	</svg>
                </div>
                <div class="s-circle">
                	<svg viewBox="0 0 24 24">
                		<path fill="#FFFFFF" d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
                	</svg>
                </div>
                <div class="s-circle">
                	<svg viewBox="0 0 24 24">
                		<path fill="#00ADEE" d="M11.979 0C5.353 0 0 5.372 0 12c0 3.21 1.261 6.138 3.313 8.312l3.414-4.882c-.08-.344-.124-.702-.124-1.071 0-2.486 2.014-4.502 4.499-4.502s4.501 2.016 4.501 4.502-2.016 4.502-4.501 4.502c-1.024 0-1.964-.343-2.715-.918l-4.717 3.352c2.423 1.706 5.385 2.684 8.566 2.684 8.012 0 14.504-6.492 14.504-14.504C26.483 4.242 20.021 0 11.979 0zm0 15.342c-1.844 0-3.342-1.498-3.342-3.342s1.498-3.342 3.342-3.342 3.342 1.498 3.342 3.342-1.498 3.342-3.342 3.342zm-1.815-3.342c0 1.002.813 1.815 1.815 1.815s1.815-.813 1.815-1.815-.813-1.815-1.815-1.815-1.815.813-1.815 1.815z"/>
                	</svg>
                </div>
                <div class="s-circle">
                	<svg viewBox="0 0 127.14 96.36">
                		<path fill="#5865F2" d="M107.7,8.07A105.15,105.15,0,0,0,81.47,0a72.06,72.06,0,0,0-3.36,6.83A97.68,97.68,0,0,0,49,6.83,72.37,72.37,0,0,0,45.64,0,105.89,105.89,0,0,0,19.39,8.09C2.79,32.65-1.71,56.6.54,80.21h0A105.73,105.73,0,0,0,32.71,96.36,77.7,77.7,0,0,0,39.6,85.25a68.42,68.42,0,0,1-10.85-5.18c.91-.66,1.8-1.34,2.66-2a75.57,75.57,0,0,0,64.32,0c.87.71,1.76,1.39,2.66,2a68.68,68.68,0,0,1-10.87,5.19,77,77,0,0,0,6.89,11.1A105.25,105.25,0,0,0,126.6,80.22h0C129.24,52.84,122.09,29.11,107.7,8.07ZM42.45,65.69C36.18,65.69,31,60,31,53s5-12.74,11.43-12.74S54,46,53.89,53,48.84,65.69,42.45,65.69Zm42.24,0C78.41,65.69,73.31,60,73.31,53s5-12.74,11.43-12.74S96.1,46,96,53,91.08,65.69,84.69,65.69Z"/>
                	</svg>
                </div>
            </div>

			<div class="auth-footer">
				Don't have an account? <RouterLink to="/signup">Sign up</RouterLink>
			</div>
		</div>
	</div>
</template>