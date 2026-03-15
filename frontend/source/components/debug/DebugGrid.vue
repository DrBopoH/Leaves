<script setup lang="ts">
	// Copyright (C) 2026 MorangTong Creative Studio
	// SPDX-License-Identifier: AGPL-3.0-or-later

	// source/components/debug/DebugGrid.vue
	import { ref } from 'vue';
	import DebugShowcase from './DebugShowcase.vue';
	import DebugSection from './DebugSection.vue';
	import UiBoard from '../ui-kit/UiBoard.vue';

	

	import UiAlert from '../ui-kit/UiAlert.vue';	
	import UiButton from '../ui-kit/UiButton.vue';
	import AppHeader from '../AppHeader.vue';
	import AuthForm from '../AuthForm.vue';

	
	const alerts = ref({
		info: true,
		error: false,
		success: false,
		warning: false
	});

	const toggleAlert = (type: keyof typeof alerts.value) => {
		alerts.value[type] = !alerts.value[type];
	};


	const buttonProps = ref({
		isLoading: false,
		disabled: false
	});

	const toggleButtonProp = (prop: keyof typeof buttonProps.value) => {
		buttonProps.value[prop] = !buttonProps.value[prop];
	};


	const authFormProps = ref({
		isLogin: true,
		isLoading: false,
		serverMessage: ''
	});

	const toggleAuthMode = () => {
		authFormProps.value.isLogin = !authFormProps.value.isLogin;
	};

	const toggleAuthLoading = () => {
		authFormProps.value.isLoading = !authFormProps.value.isLoading;
	};

	const toggleAuthError = () => {
		authFormProps.value.serverMessage = authFormProps.value.serverMessage 
			? '' 
			: 'Invalid credentials. Please try again or reset your password.';
	};

	const handleLoginSubmit = (payload: any) => {
		console.log('[AuthForm] Login submitted:', payload);
		authFormProps.value.isLoading = true;
		setTimeout(() => {
			authFormProps.value.isLoading = false;
			authFormProps.value.serverMessage = Math.random() > 0.5 ? 'Mock server error: Bad gateway' : '';
		}, 1500);
	};

	const handleSignupSubmit = (payload: any) => {
		console.log('[AuthForm] Signup submitted:', payload);
		authFormProps.value.isLoading = true;
		setTimeout(() => authFormProps.value.isLoading = false, 1500);
	};

	const handleClearError = () => {
		authFormProps.value.serverMessage = '';
	};
</script>

<template>
	<UiBoard>
		<DebugShowcase>
			<DebugSection id="alerts" title="UiAlert">
				<button 
					class="theme-toggle-btn" 
					:class="{ 'is-active': alerts.info }"
					@click="toggleAlert('info')"
				>
					{{ alerts.info ? 'Hide' : 'Info' }}
				</button>
				<UiAlert :show="alerts.info" type="info">Move me with the middle mouse button!</UiAlert>

				<button 
					class="theme-toggle-btn" 
					:class="{ 'is-active': alerts.error }"
					@click="toggleAlert('error')"
				>
					{{ alerts.error ? 'Hide' : 'Error' }}
				</button>
				<UiAlert :show="alerts.error" type="error">Critical system failure detected!</UiAlert>

				<button 
					class="theme-toggle-btn" 
					:class="{ 'is-active': alerts.success }"
					@click="toggleAlert('success')"
				>
					{{ alerts.success ? 'Hide' : 'Success' }}
				</button>
				<UiAlert :show="alerts.success" type="success">Data saved successfully.</UiAlert>

				<button 
					class="theme-toggle-btn" 
					:class="{ 'is-active': alerts.warning }"
					@click="toggleAlert('warning')"
				>
					{{ alerts.warning ? 'Hide' : 'Warn' }}
				</button>
				<UiAlert :show="alerts.warning" type="warning">Low disk space remaining.</UiAlert>
			</DebugSection>

			<DebugSection id="buttons" title="UiButton">
				<button 
					class="theme-toggle-btn" 
					:class="{ 'is-active': buttonProps.isLoading }"
					@click="toggleButtonProp('isLoading')"
				>
					{{ buttonProps.isLoading ? 'Disable' : 'Enable' }} Loading
				</button>
				<button 
					class="theme-toggle-btn" 
					:class="{ 'is-active': buttonProps.disabled }"
					@click="toggleButtonProp('disabled')"
				>
					{{ buttonProps.disabled ? 'Enable' : 'Disable' }} Button
				</button>

				<UiButton 
					variant="primary" 
					:is-loading="buttonProps.isLoading" 
					:disabled="buttonProps.disabled"
				>
					Primary
				</UiButton>
						
				<UiButton 
					variant="secondary" 
					:is-loading="buttonProps.isLoading" 
					:disabled="buttonProps.disabled"
				>
					Secondary
				</UiButton>
						
				<UiButton 
					variant="outline" 
					:is-loading="buttonProps.isLoading" 
					:disabled="buttonProps.disabled"
				>
					Outline
				</UiButton>
						
				<UiButton 
					variant="ghost" 
					:is-loading="buttonProps.isLoading" 
					:disabled="buttonProps.disabled"
				>
					Ghost
				</UiButton>

				<UiButton 
					variant="primary" 
					size="small"
					:is-loading="buttonProps.isLoading" 
					:disabled="buttonProps.disabled"
				>
					Small
				</UiButton>
						
				<UiButton 
					variant="primary" 
					size="medium"
					:is-loading="buttonProps.isLoading" 
					:disabled="buttonProps.disabled"
				>
					Medium
				</UiButton>
						
				<UiButton 
					variant="primary" 
					size="large"
					:is-loading="buttonProps.isLoading" 
					:disabled="buttonProps.disabled"
				>
					Large
				</UiButton>
			</DebugSection>

			<DebugSection id="inputs" title="UiInput">
				Inputs will be here...
			</DebugSection>

			<DebugSection id="app-header" title="AppHeader">
				<AppHeader />
			</DebugSection>

			<DebugSection id="auth-form" title="AuthForm">
				<button 
					class="theme-toggle-btn" 
					:class="{ 'is-active': authFormProps.isLogin }"
					@click="toggleAuthMode"
				>
					{{ authFormProps.isLogin ? 'Sign In' : 'Sign Up' }}
				</button>
				<button 
					class="theme-toggle-btn" 
					:class="{ 'is-active': authFormProps.isLoading }"
					@click="toggleAuthLoading"
				>
					{{ authFormProps.isLoading ? 'Waiting' : 'Loading' }}
				</button>
				<button 
					class="theme-toggle-btn" 
					:class="{ 'is-active': !!authFormProps.serverMessage }"
					@click="toggleAuthError"
				>
					return Server Error
				</button>

				<AuthForm
					:is-login="authFormProps.isLogin"
					:is-loading="authFormProps.isLoading"
					:server-message="authFormProps.serverMessage"
					@submitLogin="handleLoginSubmit"
					@submitSignup="handleSignupSubmit"
					@clearError="handleClearError"
				/>
			</DebugSection>
		</DebugShowcase>
	</UiBoard>
</template>

<style scoped>
	.theme-toggle-btn {
		background-color: var(--color-bg);
		color: var(--color-text-secondary);
		border: 1px solid var(--color-border);
		font-family: inherit;
		font-size: var(--font-size-sm);
		font-weight: var(--font-weight-medium);
		cursor: pointer;
	}

	.theme-toggle-btn:hover {
		background-color: var(--color-surface-hover);
		color: var(--color-text-primary);
		border-color: var(--color-text-secondary);
	}
</style>