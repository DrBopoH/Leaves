<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiDivider.vue

const props = defineProps<{
	text?: string;
}>();
</script>

<template>
	<div class="ui-divider">
		<span v-if="text || $slots.default" class="divider-text">
			<slot>{{ text }}</slot>
		</span>
	</div>
</template>

<style scoped>
.ui-divider {
	text-align: center;
	position: relative;
	display: flex;
	align-items: center;
	justify-content: center;
	width: 100%;
}

.ui-divider::before,
.ui-divider::after {
	content: '';
	flex: 1;
	height: 1px;
	background-color: var(--color-border);
}

.divider-text {
	font-size: 12px;
	color: var(--color-text-secondary);
	padding: 0 10px;
	background-color: transparent;
	transition: background-color 0.3s ease;
	white-space: nowrap;
}
</style>
