<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/AuthHeader.vue

import UiLogo from './ui-kit/UiLogo.vue';
</script>

<template>
	<div class="brand-header">
		<UiLogo size="large" text="Leaves" class="auth-logo" />
		<p class="brand-subtitle">Create your account and start building your conversations today.</p>
	</div>
</template>

<style scoped>
.brand-header {
	text-align: center;
	margin-bottom: 24px;
}

.auth-logo {
	justify-content: center;
	margin-bottom: 12px;
}

.brand-subtitle {
	font-size: 13px;
	color: var(--color-text-secondary);
	margin: 0;
	line-height: 1.5;
	transition: color 0.3s ease;
}
</style>
