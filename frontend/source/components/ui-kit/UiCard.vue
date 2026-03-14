<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiCard.vue

const props = defineProps<{
	title?: string;
	description?: string;
	interactive?: boolean;
}>();
</script>

<template>
	<div
		class="ui-card"
		:class="{ 'is-interactive': interactive }"
	>
		<div v-if="$slots.icon" class="card-icon">
			<slot name="icon"></slot>
		</div>

		<h3 v-if="title || $slots.title" class="card-title">
			<slot name="title">{{ title }}</slot>
		</h3>

		<div v-if="description || $slots.default" class="card-content">
			<slot>{{ description }}</slot>
		</div>
	</div>
</template>

<style scoped>
.ui-card {
	background-color: var(--color-surface);
	border: 1px solid var(--color-border);
	padding: 30px;
	border-radius: 16px;
	text-align: left;
	transition: all 0.3s ease;
}

.ui-card.is-interactive {
	cursor: pointer;
}

.ui-card.is-interactive:hover {
	border-color: var(--color-accent);
	transform: translateY(-5px);
	box-shadow: 0 10px 30px var(--color-accent-shadow);
}

.card-icon {
	font-size: 2.5rem;
	margin-bottom: 20px;
	color: var(--color-accent);
}

.card-icon :slotted(svg) {
	stroke: currentColor;
}

.card-title {
	color: var(--color-text-primary);
	font-size: 1.3rem;
	margin-bottom: 15px;
	font-weight: 600;
}

.card-content {
	color: var(--color-text-secondary);
	line-height: 1.5;
	font-size: 0.95rem;
}
</style>
