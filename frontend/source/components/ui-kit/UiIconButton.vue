<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiIconButton.vue

const props = defineProps<{
	title?: string;
	disabled?: boolean;
}>();

const emit = defineEmits<{
	(e: 'click', event: MouseEvent): void;
}>();
</script>

<template>
	<button
		class="ui-icon-button"
		:title="title"
		:disabled="disabled"
		@click="emit('click', $event)"
	>
		<slot></slot>
	</button>
</template>

<style scoped>
.ui-icon-button {
	background-color: transparent;
	border: none;
	cursor: pointer;
	padding: 8px;
	border-radius: 50%;
	display: inline-flex;
	align-items: center;
	justify-content: center;
	color: var(--color-text-secondary);
	transition: background-color 0.3s ease, color 0.3s ease, transform 0.1s;
	outline: none;
}

.ui-icon-button:hover:not(:disabled) {
	background-color: var(--color-surface);
	color: var(--color-accent);
}

.ui-icon-button:active:not(:disabled) {
	transform: scale(0.95);
}

.ui-icon-button:disabled {
	opacity: 0.5;
	cursor: not-allowed;
}

/* Make inner SVG fill nicely */
:slotted(svg) {
	width: 24px;
	height: 24px;
	transition: stroke 0.3s ease, fill 0.3s ease;
}
</style>
