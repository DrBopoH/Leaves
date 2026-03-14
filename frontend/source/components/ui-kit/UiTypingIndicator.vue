<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiTypingIndicator.vue

const props = defineProps<{
	text?: string;
}>();
</script>

<template>
	<div class="ui-typing-indicator">
		<span class="typing-dots">
			<span class="dot"></span>
			<span class="dot"></span>
			<span class="dot"></span>
		</span>

		<span v-if="text || $slots.default" class="typing-text">
			<slot>{{ text }}</slot>
		</span>
	</div>
</template>

<style scoped>
.ui-typing-indicator {
	display: inline-flex;
	align-items: center;
	gap: 8px;
	height: 20px;
	font-size: 13px;
	font-weight: 500;
	pointer-events: none;
}

.typing-dots {
	display: flex;
	gap: 3px;
}

.dot {
	width: 4px;
	height: 4px;
	background-color: var(--color-text-secondary);
	border-radius: 50%;
	animation: typing 1.4s infinite ease-in-out both;
}

.dot:nth-child(1) {
	animation-delay: -0.32s;
}

.dot:nth-child(2) {
	animation-delay: -0.16s;
}

@keyframes typing {
	0%, 80%, 100% {
		transform: scale(0);
	}
	40% {
		transform: scale(1);
	}
}

.typing-text {
	color: var(--color-text-secondary);
}
</style>
