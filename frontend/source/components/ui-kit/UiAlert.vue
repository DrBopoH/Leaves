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
	pointer-events: none;
	opacity: 0;
	transition: 
		grid-template-rows 0.3s ease, 
		margin 0.3s ease, 
		opacity 0.3s ease;
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

.type-error .alert-text {
	color: var(--color-status-error);
}

.type-warning .alert-text {
	color: var(--color-status-away);
}

.type-success .alert-text {
	color: var(--color-status-online);
}

.type-info .alert-text {
	color: var(--color-text-secondary);
}
</style>
