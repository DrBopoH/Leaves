<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/AuthForm.vue

import { ref, computed } from 'vue';
import UiButton from './ui-kit/UiButton.vue';
import UiInput from './ui-kit/UiInput.vue';
import UiCheckbox from './ui-kit/UiCheckbox.vue';

const props = defineProps<{
    isLogin: boolean;
    isLoading: boolean;
    serverMessage: string;
}>();

const emit = defineEmits<{
    (e: 'submitLogin', payload: any): void;
    (e: 'submitSignup', payload: any): void;
	(e: 'clearError'): void;
}>();

const username = ref('');
const email = ref('');
const password = ref('');
const rememberMe = ref(false);
const showPassword = ref(false);

const clientError = computed(() => {
	if (!email.value && !password.value && (props.isLogin || !username.value)) return '';
	if (!props.isLogin && username.value && (username.value.length < 3 || username.value.length > 30)) {
		return 'Username must be between 3 and 30 characters';
	}
	if (email.value && !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email.value)) {
		return 'Invalid email format';
	}
	if (password.value && password.value.length < 8) {
		return 'Password must be at least 8 characters';
	}
	return '';
});

const isFormValid = computed(() => {
	if (props.isLogin) {
		return email.value && password.value && !clientError.value;
	} else {
		return username.value && email.value && password.value && !clientError.value;
 }
});

const handleSubmit = () => {
	if (!isFormValid.value) return;

	if (props.isLogin) {
		emit('submitLogin', { email: email.value, password: password.value, rememberMe: rememberMe.value });
	} else {
		emit('submitSignup', { username: username.value, email: email.value, password: password.value });
	}
};

const handleInput = () => {
	emit('clearError');
};
</script>

<template>
 <form
 	@submit.prevent="handleSubmit"
 	class="form-content"
 	@input="handleInput"
 >
 	<div class="inputs-wrapper">
 		<transition name="fade" mode="out-in">
 			<div :key="isLogin ? 'login-inputs' : 'reg-inputs'" class="inputs-container">
 				<div v-if="!isLogin" class="input-margin-large">
 					<UiInput
 						v-model="username"
 						placeholder="username"
 						required
 					/>
 				</div>

 				<div class="input-margin-small">
 					<UiInput
 						type="email"
 						v-model="email"
 						placeholder="your@email.com"
 						required
 					/>
 				</div>

 				<div class="input-margin-none">
 					<UiInput
 						:type="showPassword ? 'text' : 'password'"
 						v-model="password"
 						placeholder="password"
 						required
 					>
 						<template #append>
 							<svg
 								width="20"
 								height="20"
 								@mousedown="showPassword = true"
 								@mouseup="showPassword = false"
 								@mouseleave="showPassword = false"
 								@touchstart.prevent="showPassword = true"
 								@touchend.prevent="showPassword = false"
 								@touchcancel.prevent="showPassword = false"
 								@contextmenu.prevent
 								class="eye-icon"
 								viewBox="0 0 24 24"
 								fill="none"
 								stroke="currentColor"
 								stroke-width="2"
 								stroke-linecap="round"
 								stroke-linejoin="round"
 							>
 								<path v-if="!showPassword" d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24M1 1l22 22"/>
 								<path v-else d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path>
 								<circle v-if="showPassword" cx="12" cy="12" r="3"></circle>
 							</svg>
 						</template>
 					</UiInput>
 				</div>

 				<div v-if="isLogin" class="options-row">
 					<UiCheckbox v-model="rememberMe" label="remember me" />
 					<a href="#" class="forgot-password">forgot password?</a>
 				</div>
 			</div>
 		</transition>
 	</div>

 	<div
 		class="error-reservoir"
 		:class="{ 'has-error': clientError || serverMessage }"
 	>
		<div class="error-content">
			<span v-if="clientError" class="error-text error-client">{{ clientError }}</span>
			<span v-else-if="serverMessage" class="error-text error-server">{{ serverMessage }}</span>
		</div>
	</div>

	 <UiButton
		 type="submit"
		 variant="primary"
		 size="large"
		 :is-loading="isLoading"
		 :disabled="!isFormValid"
		 class="submit-btn"
	 >
		 {{ isLogin ? 'Sign in' : 'Sign up' }}
	 </UiButton>
</form>
</template>

<style scoped>
.form-content {
	display: flex;
	flex-direction: column;
	gap: 16px;
	margin-bottom: 24px;
}

.inputs-wrapper {
	position: relative;
}

.inputs-container {
	display: flex;
	flex-direction: column;
	width: 100%;
}

.input-margin-large {
	margin-bottom: 28px;
}

.input-margin-small {
	margin-bottom: 16px;
}

.input-margin-none {
	margin-bottom: 0;
}

.eye-icon {
	width: 20px;
	height: 20px;
	cursor: pointer;
	transition: color 0.2s ease;
	pointer-events: auto;
}

.eye-icon:hover {
	color: var(--color-text-primary);
}

.options-row {
	display: flex;
	justify-content: space-between;
	align-items: center;
	font-size: 12px;
	padding: 0 10px;
	margin: 10px 0 6px 0;
}

.forgot-password {
	color: var(--color-text-secondary);
	text-decoration: none;
	transition: color 0.2s ease;
}

.forgot-password:hover {
	color: var(--color-text-primary);
}

.error-reservoir {
	display: grid;
	grid-template-rows: 0fr;
	transition: grid-template-rows 0.3s ease, margin 0.3s ease;
}

.error-reservoir.has-error {
	grid-template-rows: 1fr;
	margin-top: 4px;
	margin-bottom: 8px;
}

.error-content {
	overflow: hidden;
	display: flex;
	align-items: center;
	justify-content: center;
}

.error-text {
	font-size: 13px;
	font-weight: 500;
	text-align: center;
	margin: 0;
	padding: 2px 0;
}

.error-client {
	color: var(--color-text-secondary);
}

.error-server {
	color: var(--color-status-error);
}

.submit-btn {
	width: 100%;
}

.fade-enter-active,
.fade-leave-active {
	transition: opacity 0.2s ease, transform 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
	opacity: 0;
	transform: translateY(-5px);
}
</style>
