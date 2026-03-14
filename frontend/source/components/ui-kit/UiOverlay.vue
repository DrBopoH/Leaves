<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiOverlay.vue

const props = defineProps<{
	show: boolean;
	zIndex?: number;
}>();

const emit = defineEmits<{
	(e: 'click', event: MouseEvent): void;
}>();
</script>

<template>
	<div
		class="ui-overlay"
		:class="{ 'is-visible': show }"
		:style="{ zIndex: zIndex || 99 }"
		@click="emit('click', $event)"
	></div>
</template>

<style scoped>
.ui-overlay {
	position: fixed;
	top: 0;
	left: 0;
	right: 0;
	bottom: 0;
	background-color: var(--color-overlay);
	backdrop-filter: blur(2px);
	opacity: 0;
	pointer-events: none;
	transition: opacity 0.3s ease;
}

.ui-overlay.is-visible {
	opacity: 1;
	pointer-events: auto;
}
</style>
