<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiLogo.vue

const props = defineProps<{
	size?: 'small' | 'medium' | 'large';
	text?: string;
}>();
</script>

<template>
	<div class="ui-logo" :class="`size-${size || 'medium'}`">
		<img src="/Logo.svg" alt="Logo" class="logo-icon" />
		<span v-if="text" class="logo-text">{{ text }}</span>
	</div>
</template>

<style scoped>
.ui-logo {
	display: inline-flex;
	align-items: center;
	gap: 10px;
	user-select: none;
}

.logo-text {
	font-weight: 700;
	color: var(--color-text-primary);
	letter-spacing: -0.5px;
}

.ui-logo.size-small .logo-icon { width: 22px; height: 22px; }
.ui-logo.size-small .logo-text { font-size: 15px; }

.ui-logo.size-medium .logo-icon { width: 28px; height: 28px; }
.ui-logo.size-medium .logo-text { font-size: 1.4rem; }

.ui-logo.size-large .logo-icon { width: 32px; height: 32px; }
.ui-logo.size-large .logo-text { font-size: 28px; }
</style>
