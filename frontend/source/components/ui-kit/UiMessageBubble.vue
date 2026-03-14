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
	font-weight: var(--font-weight-semibold);
	font-size: var(--font-size-md);
}

.message-time {
	font-size: var(--font-size-xs);
	color: var(--color-text-secondary);
}

.message-content :deep(.msg-text) {
	font-size: var(--font-size-md);
	line-height: var(--line-height-md);
	color: var(--color-text-primary);
	margin: 0;
	word-break: break-word;
}

.message-content :deep(h1), .message-content :deep(h2), .message-content :deep(h3) {
    margin: 4px 0;
    line-height: 1.2;
}
.message-content :deep(.msg-quote) {
	border-left: 3px solid var(--color-accent);
	margin: 4px 0;
	padding: 4px 10px;
	color: var(--color-text-secondary);
	background-color: var(--color-accent-alpha);
	border-radius: 0 4px 4px 0;
}

.message-content :deep(.msg-divider) {
	border: none;
	border-top: 2px solid var(--color-border-hover);
	border-radius: 2px;
	transition: width var(--theme-transition-duration) ease;
}

.message-content :deep(.msg-text-divider) {
	display: flex;
	align-items: center;
	justify-content: center;
	gap: 12px;
	margin: 8px 0;
	color: var(--color-text-secondary);
	font-weight: var(--font-weight-semibold);
	font-size: var(--font-size-sm);
	letter-spacing: var(--letter-spacing-wide);
	width: 100%;
}

.message-content :deep(.msg-text-divider span) {
	flex-shrink: 0;
	max-width: calc(100% - 24px);
	text-align: center;
	word-break: break-word;
}

.message-content :deep(.msg-text-divider-line) {
	height: 2px;
	background-color: var(--color-border-hover);
	border-radius: 2px;
	flex-basis: 0;
}


.message-content :deep(.msg-image) {
	max-width: 100%;
	max-height: 250px;
	border-radius: 8px;
	margin-top: 8px;
	border: 1px solid var(--color-border);
	display: block;
}

.message-content :deep(.msg-link) {
	color: var(--color-accent);
	text-decoration: underline;
	text-underline-offset: 2px;
}

.message-content :deep(.msg-link:hover) {
	color: var(--color-accent-hover);
}

.message-bubble {
	background-color: transparent;
	border: 1px solid transparent;
	padding: 10px 16px;
	border-radius: 8px;
	color: var(--color-text-primary);
	font-size: var(--font-size-md);
	line-height: var(--line-height-md);
	word-break: break-word;
}

.ui-message.is-mine .message-bubble {
	background-color: var(--color-surface);
	border: 1px solid var(--color-border);
	padding: 10px 16px;
	border-radius: 16px;
	border-top-left-radius: 8px;
}
</style>
