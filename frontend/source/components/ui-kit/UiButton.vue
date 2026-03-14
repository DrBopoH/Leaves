<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiButton.vue

import UiLoader from './UiLoader.vue';

const props = defineProps<{
	isLoading?: boolean;
	disabled?: boolean;
	type?: 'button' | 'submit' | 'reset';
	variant?: 'primary' | 'secondary' | 'outline' | 'ghost';
	size?: 'small' | 'medium' | 'large';
}>();

const emit = defineEmits<{
	(e: 'click', event: MouseEvent): void;
}>();
</script>

<template>
	<button
		:type="type || 'button'"
		class="ui-button"
		:class="[
			`variant-${variant || 'primary'}`,
			`size-${size || 'medium'}`,
			{ 'is-loading': isLoading }
		]"
		:disabled="disabled || isLoading"
		@click="emit('click', $event)"
	>
		<div class="content" :style="{ opacity: isLoading ? 0 : 1 }">
			<slot></slot>
		</div>

		<div v-if="isLoading" class="loader">
			<UiLoader size="medium" />
		</div>
	</button>
</template>

<style scoped>
.ui-button {
	position: relative;
	border: none;
	font-weight: 600;
	cursor: pointer;
	display: inline-flex;
	align-items: center;
	justify-content: center;
	transition: background-color 0.2s, color 0.2s, border-color 0.2s, box-shadow 0.2s, transform 0.1s;
	outline: none;
}

/* Sizes */
.size-small {
	border-radius: 16px;
	padding: 6px 12px;
	font-size: 13px;
}

.size-medium {
	border-radius: 20px;
	padding: 10px 20px;
	font-size: 14px;
}

.size-large {
	border-radius: 30px;
	padding: 14px 32px;
	font-size: 16px;
}

.content {
	display: flex;
	align-items: center;
	gap: 8px;
	transition: opacity 0.2s;
}

.loader {
	position: absolute;
	width: 24px;
	height: 24px;
}

/* Primary Variant (CTA Style) */
.variant-primary {
	background-color: transparent;
	color: var(--color-accent);
	border: 2px solid var(--color-accent);
	box-shadow: 0 0 15px var(--color-accent-shadow);
	transition: all 0.3s ease;
}

.variant-primary:hover:not(:disabled) {
	background-color: var(--color-accent);
	color: var(--color-text-inverse);
	box-shadow: 0 0 30px var(--color-accent-shadow);
	transform: translateY(-3px);
}

/* Secondary Variant */
.variant-secondary {
	background-color: var(--color-surface);
	color: var(--color-text-primary);
	border: 1px solid var(--color-border);
}

.variant-secondary:hover:not(:disabled) {
	border-color: var(--color-text-secondary);
}

/* Outline Variant */
.variant-outline {
	background-color: transparent;
	color: var(--color-text-primary);
	border: 1px solid var(--color-border);
}

.variant-outline:hover:not(:disabled) {
	background-color: var(--color-surface-hover);
	border-color: var(--color-accent);
	color: var(--color-accent);
}

/* Ghost Variant */
.variant-ghost {
	background-color: transparent;
	color: var(--color-text-primary);
}

.variant-ghost:hover:not(:disabled) {
	background-color: var(--color-surface);
}

/* Disabled State */
.ui-button:disabled {
	opacity: 0.5;
	cursor: not-allowed;
}
</style>
