<script setup lang="ts">
import { ref, nextTick, watch } from 'vue';

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
        <div class="typing-indicator" v-show="activeTypingUsers.length > 0">
            <span class="typing-dots">
                <span class="dot"></span><span class="dot"></span><span class="dot"></span>
            </span>
            <span class="typing-text">
                {{ typingText }}
            </span>
        </div>

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
            <button @click="onSend" class="send-btn" :disabled="!modelValue.trim() || !isConnected">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="22" y1="2" x2="11" y2="13"></line><polygon points="22 2 15 22 11 13 2 9 22 2"></polygon></svg>
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

.input-wrapper {
	background-color: var(--bg-surface);
	border: 1px solid var(--border);
	border-radius: 16px;
	display: flex;
	align-items: flex-end;
	padding: 12px 20px;
}
.input-wrapper:focus-within { border-color: rgba(95, 202, 8, 0.4); }

.chat-input {
    flex: 1;
    background: transparent;
    border: none;
    color: var(--text-primary) !important;
    font-size: 14px;
    outline: none;
    resize: none;
    max-height: 120px;
    font-family: inherit;
    line-height: 1.4;
    padding: 4px 0;
    word-wrap: break-word;
    white-space: pre-wrap;
    overflow-y: auto;
}
.chat-input::placeholder { color: var(--text-secondary) !important; }

.send-btn { background: var(--accent); color: #ffffff; border: none; width: 32px; height: 32px; border-radius: 8px; display: flex; justify-content: center; align-items: center; cursor: pointer; transition-property: background-color, transform, color; transition-duration: 0.2s; flex-shrink: 0; margin-left: 10px; margin-bottom: 2px; }
.send-btn:hover:not(:disabled) { background: var(--accent-hover); transform: scale(1.05); }
.send-btn:disabled { background: var(--bg-active); color: var(--text-secondary); cursor: not-allowed; }

.typing-indicator {
    height: 20px;
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 0 4px;
    background-color: transparent;
    font-size: 13px;
    font-weight: 500;
    pointer-events: none;
}
.typing-dots {
    display: flex;
    gap: 3px;
}
.typing-dots .dot {
    width: 4px;
    height: 4px;
    background-color: var(--text-secondary);
    border-radius: 50%;
    animation: typing 1.4s infinite ease-in-out both;
}
.typing-dots .dot:nth-child(1) { animation-delay: -0.32s; }
.typing-dots .dot:nth-child(2) { animation-delay: -0.16s; }
@keyframes typing {
    0%, 80%, 100% { transform: scale(0); }
    40% { transform: scale(1); }
}
.typing-text {
    color: var(--text-secondary);
}

@media (max-width: 768px) {
    .chat-input-area {
		padding: 0 16px 16px 16px;
		padding-bottom: calc(16px + env(safe-area-inset-bottom));
	}
}
</style>
