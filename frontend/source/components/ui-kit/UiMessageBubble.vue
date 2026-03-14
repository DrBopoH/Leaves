<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiMessageBubble.vue

import UiAvatar from './UiAvatar.vue';

const props = defineProps<{
	username: string;
	color?: string;
	time?: string;
	isMine?: boolean;
	showHeader?: boolean;
}>();
</script>

<template>
	<div
		class="ui-message"
		:class="{ 'is-mine': isMine, 'no-header': !showHeader }"
	>
		<div
			class="message-avatar-container"
			:class="{ 'hidden': !showHeader }"
		>
			<UiAvatar
				:name="username"
				:color="color"
				size="medium"
				class="message-avatar"
			/>
		</div>

		<div class="message-content">
			<div v-if="showHeader" class="message-header">
				<span class="message-username" :style="{ color: color || 'inherit' }">
					{{ username }}
				</span>

				<span v-if="time" class="message-time">{{ time }}</span>
			</div>

			<div class="message-bubble">
				<slot></slot>
			</div>
		</div>
	</div>
</template>

<style scoped>
.ui-message {
	display: flex;
	gap: 16px;
	align-items: flex-start;
	max-width: 100%;
}

.message-avatar-container {
	flex-shrink: 0;
}

.message-avatar-container.hidden {
	visibility: hidden;
}

.message-avatar {
	border-radius: 12px;
}

.message-content {
	max-width: calc(100% - 56px); /* 40px avatar + 16px gap */
	display: flex;
	flex-direction: column;
}

.message-header {
	display: flex;
	align-items: baseline;
	gap: 10px;
	margin-bottom: 4px;
}

.message-username {
	font-weight: 600;
	font-size: 15px;
}

.message-time {
	font-size: 12px;
	color: var(--color-text-secondary);
}

.message-bubble {
	background-color: transparent;
	border: 1px solid transparent;
	padding: 2px 0;
	border-radius: 8px;
	color: var(--color-text-primary);
	font-size: 15px;
	line-height: 1.5;
	word-break: break-word;
}

.ui-message.is-mine .message-bubble {
	background-color: var(--color-surface);
	border: 1px solid var(--color-border);
	padding: 10px 16px;
	border-radius: 16px;
	border-top-left-radius: 4px;
}
</style>
