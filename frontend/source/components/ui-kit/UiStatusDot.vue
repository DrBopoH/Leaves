<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiStatusDot.vue

const props = defineProps<{
	status?: 'online' | 'offline' | 'away' | 'dnd';
}>();
</script>

<template>
	<span class="ui-status-dot" :class="`status-${status || 'offline'}`"></span>
</template>

<style scoped>
.ui-status-dot {
	width: 8px;
	height: 8px;
	border-radius: 50%;
	flex-shrink: 0;
	display: inline-block;
	transition: background-color 0.3s ease, box-shadow 0.3s ease;
}

.status-offline {
	background-color: var(--color-status-offline);
}

.status-online {
	background-color: var(--color-status-online);
	box-shadow: 0 0 8px var(--color-accent-shadow);
}

.status-away {
	background-color: var(--color-status-away);
	box-shadow: 0 0 8px var(--color-status-away);
}

.status-dnd {
	background-color: var(--color-status-dnd);
	box-shadow: 0 0 8px var(--color-status-dnd);
}
</style>
