<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ChatInput.vue

import { ref, nextTick, watch } from 'vue';
import UiTypingIndicator from './ui-kit/UiTypingIndicator.vue';

const props = defineProps<{
	modelValue: string;
	isConnected: boolean;
	activeTypingUsers: string[];
	typingText: string;
}>();

const emit = defineEmits<{
	(e: 'update:modelValue', value: string): void;
	(e: 'send'): void;
	(e: 'typing'): void;
}>();

const chatInputRef = ref<HTMLTextAreaElement | null>(null);

const onInput = (e: Event) => {
	const target = e.target as HTMLTextAreaElement;
	emit('update:modelValue', target.value);

	target.style.height = 'auto';
	target.style.height = Math.min(target.scrollHeight, 120) + 'px';

	emit('typing');
};

const onKeydown = (e: KeyboardEvent) => {
	if (e.key === 'Enter' && !e.shiftKey) {
		e.preventDefault();
		onSend();
	}
};

const onSend = () => {
	if (!props.modelValue.trim() || !props.isConnected) return;
	emit('send');

	nextTick(() => {
		if (chatInputRef.value) {
			chatInputRef.value.style.height = 'auto';
		}
	});
};

watch(() => props.modelValue, (newVal) => {
	if (!newVal && chatInputRef.value) {
		nextTick(() => {
			if (chatInputRef.value) {
				chatInputRef.value.style.height = 'auto';
			}
		});
	}
});
</script>

<template>
	<div class="chat-input-area">
		<UiTypingIndicator
			v-show="activeTypingUsers.length > 0"
			:text="typingText"
			class="typing-container"
		/>

		<div class="input-wrapper">
			<textarea
				ref="chatInputRef"
				:value="modelValue"
				@input="onInput"
				@keydown="onKeydown"
				placeholder="Message #general... (Shift+Enter for new line)"
				class="chat-input"
				rows="1"
				maxlength="2000"
			></textarea>

			<button
				@click="onSend"
				class="send-btn"
				:disabled="!modelValue.trim() || !isConnected"
			>
				<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
					<line x1="22" y1="2" x2="11" y2="13"></line>
					<polygon points="22 2 15 22 11 13 2 9 22 2"></polygon>
				</svg>
			</button>
		</div>
	</div>
</template>

<style scoped>
.chat-input-area {
	position: absolute;
	bottom: 0;
	left: 0;
	right: 0;
	padding: 0 8px 12px 8px;
	padding-bottom: calc(8px + env(safe-area-inset-bottom));
	background: transparent;
	pointer-events: none;
	z-index: 10;
}

.chat-input-area > * {
	pointer-events: auto;
}

.typing-container {
	margin-left: 12px;
	margin-bottom: 4px;
}

.input-wrapper {
	background-color: var(--color-surface);
	border: 1px solid var(--color-border);
	border-radius: 16px;
	display: flex;
	align-items: flex-end;
	padding: 12px 20px;
	transition: border-color var(--theme-transition-duration) ease, box-shadow var(--theme-transition-duration) ease;
}

.input-wrapper:focus-within {
	border-color: var(--color-accent-shadow);
}

.chat-input {
	flex: 1;
	background: transparent;
	border: none;
	color: var(--color-text-primary);
	font-size: var(--font-size-base);
	outline: none;
	resize: none;
	max-height: 120px;
	line-height: var(--line-height-base);
	padding: 4px 0;
	word-wrap: break-word;
	white-space: pre-wrap;
	overflow-y: auto;
}

.chat-input::placeholder {
	color: var(--color-text-secondary);
}

.send-btn {
	background-color: var(--color-accent);
	color: var(--color-text-inverse);
	border: none;
	width: 32px;
	height: 32px;
	border-radius: 8px;
	display: flex;
	justify-content: center;
	align-items: center;
	cursor: pointer;
	transition: background-color var(--theme-transition-duration), transform 0.2s, color var(--theme-transition-duration), box-shadow var(--theme-transition-duration);
	flex-shrink: 0;
	margin-left: 10px;
	margin-bottom: 2px;
}

.send-btn:hover:not(:disabled) {
	background-color: var(--color-accent-hover);
	transform: scale(1.05);
}

.send-btn:disabled {
	background-color: var(--color-surface-hover);
	color: var(--color-text-secondary);
	cursor: not-allowed;
}

@media (max-width: 768px) {
	.chat-input-area {
		padding: 0 16px 16px 16px;
		padding-bottom: calc(16px + env(safe-area-inset-bottom));
	}
}
</style>
