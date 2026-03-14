<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiMenuLink.vue

const props = defineProps<{
	active?: boolean;
}>();
</script>

<template>
	<div
		class="ui-menu-link"
		:class="{ 'is-active': active }"
	>
		<div v-if="$slots.icon" class="link-icon">
			<slot name="icon"></slot>
		</div>

		<span class="link-text">
			<slot></slot>
		</span>
	</div>
</template>

<style scoped>
.ui-menu-link {
	display: flex;
	align-items: center;
	gap: 8px;
	padding: 10px 15px;
	border-radius: 8px;
	color: var(--color-text-secondary);
	cursor: pointer;
	font-weight: 500;
	user-select: none;
	transition: background-color var(--theme-transition-duration) ease, color var(--theme-transition-duration) ease;
}

.ui-menu-link:hover {
	background-color: var(--color-surface-hover);
	color: var(--color-text-primary);
}

.ui-menu-link.is-active {
	background-color: var(--color-surface-hover);
	color: var(--color-accent);
}

.link-icon {
	display: flex;
	align-items: center;
	justify-content: center;
}

.link-text {
	flex: 1;
	white-space: nowrap;
	overflow: hidden;
	text-overflow: ellipsis;
}
</style>
