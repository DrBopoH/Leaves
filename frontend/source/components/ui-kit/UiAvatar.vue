<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiAvatar.vue

const props = defineProps<{
	name: string;
	color?: string;
	size?: 'small' | 'medium' | 'large';
}>();
</script>

<template>
	<div
		class="ui-avatar"
		:class="`size-${size || 'medium'}`"
		:style="{
			borderColor: color || 'var(--color-border)',
			color: color || 'var(--color-accent)'
		}"
	>
		{{ name ? name.charAt(0).toUpperCase() : '?' }}
	</div>
</template>

<style scoped>
.ui-avatar {
	background-color: var(--color-border);
	border-radius: 50%;
	display: inline-flex;
	align-items: center;
	justify-content: center;
	font-weight: bold;
	border: 2px solid var(--color-surface);
	transition: background-color 0.3s ease, border-color 0.3s ease;
}

.ui-avatar.size-small {
	width: 24px;
	height: 24px;
	font-size: 0.8rem;
}

.ui-avatar.size-medium {
	width: 36px;
	height: 36px;
	font-size: 1.1rem;
}

.ui-avatar.size-large {
	width: 48px;
	height: 48px;
	font-size: 1.5rem;
}
</style>
