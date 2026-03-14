<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiSocialLink.vue

const props = defineProps<{
	href?: string;
}>();
</script>

<template>
	<component
		:is="href ? 'a' : 'button'"
		:href="href"
		class="ui-social-link"
	>
		<slot></slot>
	</component>
</template>

<style scoped>
.ui-social-link {
	width: 40px;
	height: 40px;
	background-color: var(--color-bg);
	border: 1px solid var(--color-border);
	border-radius: 50%;
	display: inline-flex;
	align-items: center;
	justify-content: center;
	cursor: pointer;
	transition: background-color 0.2s, border-color 0.2s, color 0.2s, transform 0.2s;
	color: var(--color-text-primary);
	text-decoration: none;
	outline: none;
	padding: 0;
}

.ui-social-link:hover {
	background-color: var(--color-surface);
	border-color: var(--color-accent);
	color: var(--color-accent);
	transform: translateY(-2px);
}

.ui-social-link:active {
	transform: translateY(0);
}

/* Base style for icons passed into the slot */
:slotted(svg),
:slotted(img) {
	width: 20px;
	height: 20px;
	object-fit: contain;
}
</style>
