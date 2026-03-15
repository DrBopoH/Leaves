<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiGlowBackground.vue
</script>

<template>
	<div class="ui-glow-background">
		<div class="glow-orb top-left"></div>
		<div class="glow-orb bottom-right"></div>
	</div>
</template>

<style scoped>
.ui-glow-background {
	position: absolute;
	top: 0;
	left: 0;
	right: 0;
	bottom: 0;
	z-index: 0;
	pointer-events: none;
	overflow: hidden;
}

.glow-orb {
	position: absolute;
	width: 400px;
	height: 400px;
	background: radial-gradient(circle, var(--color-surface-hover) 0%, transparent 70%);
	border-radius: 50%;
	transition: background 0.3s ease;
}

.top-left {
	top: -100px;
	left: -150px;
}

.bottom-right {
	bottom: -100px;
	right: -150px;
}
</style>
