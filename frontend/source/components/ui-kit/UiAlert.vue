<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiAlert.vue

const props = defineProps<{
	type?: 'error' | 'success' | 'warning' | 'info';
	message?: string;
	show?: boolean;
}>();
</script>

<template>
	<div
		class="ui-alert-reservoir"
		:class="[
			`type-${type || 'error'}`,
			{ 'is-visible': show }
		]"
	>
		<div class="alert-content">
			<span class="alert-text">
				<slot>{{ message }}</slot>
			</span>
		</div>
	</div>
</template>

<style scoped>
.ui-alert-reservoir {
	display: grid;
	grid-template-rows: 0fr;
	transition: grid-template-rows 0.3s ease, margin 0.3s ease, opacity 0.3s ease;
	opacity: 0;
	pointer-events: none;
}

.ui-alert-reservoir.is-visible {
	grid-template-rows: 1fr;
	margin-top: 4px;
	margin-bottom: 8px;
	opacity: 1;
	pointer-events: auto;
}

.alert-content {
	overflow: hidden;
	display: flex;
	align-items: center;
	justify-content: center;
}

.alert-text {
	font-size: 13px;
	font-weight: 500;
	text-align: center;
	margin: 0;
	padding: 2px 0;
}

/* Variant: Error */
.type-error .alert-text {
	color: var(--color-status-error);
}

/* Variant: Warning */
.type-warning .alert-text {
	color: var(--color-status-away);
}

/* Variant: Success */
.type-success .alert-text {
	color: var(--color-status-online);
}

/* Variant: Info */
.type-info .alert-text {
	color: var(--color-text-secondary);
}
</style>
