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
</script>

<template>
	<UiBoard>
		<DebugShowcase>
			<DebugSection id="alerts" title="UiAlert">
				<div class="debug-alert-grid">
					<div class="debug-alert-card">
						<button 
							class="theme-toggle-btn" 
							:class="{ 'is-active': alerts.info }"
							@click="toggleAlert('info')"
						>
							{{ alerts.info ? 'Hide' : 'Show' }}
						</button>
						<UiAlert :show="alerts.info" type="info">Move me with the middle mouse button!</UiAlert>
					</div>

					<div class="debug-alert-card">
						<button 
							class="theme-toggle-btn" 
							:class="{ 'is-active': alerts.error }"
							@click="toggleAlert('error')"
						>
							{{ alerts.error ? 'Hide' : 'Show' }}
						</button>
						<UiAlert :show="alerts.error" type="error">Critical system failure detected!</UiAlert>
					</div>

					<div class="debug-alert-card">
						<button 
							class="theme-toggle-btn" 
							:class="{ 'is-active': alerts.success }"
							@click="toggleAlert('success')"
						>
							{{ alerts.success ? 'Hide' : 'Show' }}
						</button>
						<UiAlert :show="alerts.success" type="success">Data saved successfully.</UiAlert>
					</div>

					<div class="debug-alert-card">
						<button 
							class="theme-toggle-btn" 
							:class="{ 'is-active': alerts.warning }"
							@click="toggleAlert('warning')"
						>
							{{ alerts.warning ? 'Hide' : 'Show' }}
						</button>
						<UiAlert :show="alerts.warning" type="warning">Low disk space remaining.</UiAlert>
					</div>
				</div>
			</DebugSection>

			<DebugSection id="buttons" title="UiButton">
				<div class="debug-controls-bar">
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
				</div>

				<div class="debug-button-grid">
					<div class="debug-button-row">
						<UiButton 
							variant="primary" 
							:is-loading="buttonProps.isLoading" 
							:disabled="buttonProps.disabled"
						>Primary</UiButton>
						
						<UiButton 
							variant="secondary" 
							:is-loading="buttonProps.isLoading" 
							:disabled="buttonProps.disabled"
						>Secondary</UiButton>
						
						<UiButton 
							variant="outline" 
							:is-loading="buttonProps.isLoading" 
							:disabled="buttonProps.disabled"
						>Outline</UiButton>
						
						<UiButton 
							variant="ghost" 
							:is-loading="buttonProps.isLoading" 
							:disabled="buttonProps.disabled"
						>Ghost</UiButton>
					</div>

					<div class="debug-button-row">
						<UiButton 
							variant="primary" 
							size="small"
							:is-loading="buttonProps.isLoading" 
							:disabled="buttonProps.disabled"
						>Small</UiButton>
						
						<UiButton 
							variant="primary" 
							size="medium"
							:is-loading="buttonProps.isLoading" 
							:disabled="buttonProps.disabled"
						>Medium</UiButton>
						
						<UiButton 
							variant="primary" 
							size="large"
							:is-loading="buttonProps.isLoading" 
							:disabled="buttonProps.disabled"
						>Large</UiButton>
					</div>
				</div>
			</DebugSection>

			<DebugSection id="inputs" title="UiInput">
				<div>Inputs will be here...</div>
			</DebugSection>
		</DebugShowcase>
	</UiBoard>
</template>

<style scoped>
	.debug-alert-card {
		display: flex;
		align-items: flex-start;
		background-color: var(--color-surface);
		border: 1px solid var(--color-border);
		border-radius: 8px;
	}

	.debug-alert-card:hover {
		border-color: var(--color-border-hover);
	}

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