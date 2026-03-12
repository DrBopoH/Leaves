<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ChatMessages.vue
import { useUserStore } from '../stores/user';

const props = defineProps<{
    message: {
        id: string | number;
        user: string;
        text: string;
        time: string;
    };
    isMine: boolean;
    showHeader: boolean;
}>();

const userStore = useUserStore();

const formatMessage = (text: string) => {
    if (!text) return '';

    let safeText = text.replace(/[&<"'>]/g, (m) => {
        switch (m) {
            case '&': return '&';
            case '<': return '<';
            case '>': return '>';
            case '"': return '"';
            case "'": return '&#039;';
            default: return m;
        }
    });

    safeText = safeText.replace(/\\\\n/g, '&#92;n');
    safeText = safeText.replace(/\\\\t/g, '&#92;t');
    safeText = safeText.replace(/\\\\\*/g, '&#42;');

    safeText = safeText.replace(/^### (.*?)\r?$/gm, '<h3>$1</h3>');
    safeText = safeText.replace(/^## (.*?)\r?$/gm, '<h2>$1</h2>');
    safeText = safeText.replace(/^# (.*?)\r?$/gm, '<h1>$1</h1>');

    safeText = safeText.replace(/^> (.*?)\r?$/gm, '<blockquote class="msg-quote">$1</blockquote>');

    safeText = safeText.replace(/^(---+)(.+?)(---+) *\r?$/gm, (match, p1, p2, p3) => {
        return `<div class="msg-text-divider" style="--grow-left: ${p1.length}; --grow-right: ${p3.length};"><span>${p2.trim()}</span></div>`;
    });

    safeText = safeText.replace(/^(---+) *\r?$/gm, (match, p1) => {
        let pct = Math.min(100, p1.length * 10);
        return `<hr class="msg-divider" style="width: ${pct}%" />`;
    });

    safeText = safeText.replace(/\\n/g, '<br>'); // текстовый \n
    safeText = safeText.replace(/\\t/g, '&nbsp;&nbsp;&nbsp;&nbsp;'); // текстовый \t

    safeText = safeText.replace(/\n/g, '<br>');

    safeText = safeText.replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>');
    safeText = safeText.replace(/__(.*?)__/g, '<u>$1</u>'); // Подчеркивание
    safeText = safeText.replace(/\*(.*?)\*/g, '<em>$1</em>');
    safeText = safeText.replace(/~~(.*?)~~/g, '<del>$1</del>');

    const urlRegex = /(https?:\/\/[^\s<]+)/g;
    safeText = safeText.replace(urlRegex, (url) => {
        if (url.match(/\.(jpeg|jpg|gif|png|webp)(\?.*)?$/i)) {
            return `<img src="${url}" class="msg-image" alt="$media" loading="lazy" />`;
        }
        return `<a href="${url}" target="_blank" class="msg-link">${url}</a>`;
    });

    return safeText;
};
</script>

<template>
    <div class="message">
        <div class="msg-avatar"
             :style="{
                 borderColor: userStore.getUserColor(message.user),
                 color: userStore.getUserColor(message.user),
                 visibility: !showHeader ? 'hidden' : 'visible'
             }">
            {{ message.user.charAt(0).toUpperCase() }}
        </div>

        <div class="msg-content">
            <div class="msg-header" v-if="showHeader">
                <span class="msg-username" :style="{ color: userStore.getUserColor(message.user) }">{{ message.user }}</span>
                <span class="msg-time">{{ message.time }}</span>
            </div>

            <div class="msg-bubble" :class="{ 'is-mine': isMine }">
                <p class="msg-text" v-html="formatMessage(message.text)"></p>
            </div>
        </div>
    </div>
</template>

<style scoped>
.message { display: flex; gap: 16px; align-items: flex-start; max-width: 100%; transition-property: margin, background-color, color; transition-duration: 0.2s, 0.3s, 0.3s; }
.msg-avatar {
    width: 40px; height: 40px; border-radius: 12px;
    background-color: var(--bg-surface); border: 1px solid;
    display: flex; justify-content: center; align-items: center;
    font-weight: bold; font-size: 18px; flex-shrink: 0;
}
.msg-content { max-width: calc(100% - 56px); }
.msg-header { display: flex; align-items: baseline; gap: 10px; margin-bottom: 4px; }
.msg-username { font-weight: 600; font-size: 15px; }
.msg-time { font-size: 12px; color: var(--text-secondary); }

:global(.app-layout.light-theme .msg-username),
:global(.app-layout.light-theme .msg-avatar) {
    filter: brightness(0.75) contrast(1.2);
}

.msg-bubble { background-color: transparent; border: 1px solid transparent; padding: 2px 0; border-radius: 8px; }
.msg-bubble.is-mine {
    background-color: var(--bg-active);
    border: 1px solid var(--border-active);
    padding: 10px 16px;
    border-radius: 16px;
    border-top-left-radius: 4px;
}

.msg-text {
    font-size: 15px;
    line-height: 1.5;
    color: var(--text-primary);
    margin: 0;
    word-break: break-word;
}

:deep(h1), :deep(h2), :deep(h3) {
    margin: 4px 0;
    line-height: 1.2;
}
:deep(.msg-quote) {
    border-left: 3px solid var(--accent);
    margin: 4px 0;
    padding: 4px 10px;
    color: var(--text-secondary);
    background-color: rgba(95, 202, 8, 0.05);
    border-radius: 0 4px 4px 0;
}
:deep(.msg-divider) {
    border: none;
    border-top: 2px solid var(--text-muted);
    border-radius: 2px;
    transition: width 0.3s ease;
}

:deep(.msg-text-divider) {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 12px;
    margin: 8px 0;
    color: var(--text-secondary);
    font-weight: 600;
    font-size: 13px;
    letter-spacing: 0.5px;
    width: 100%;
}
:deep(.msg-text-divider span) {
    flex-shrink: 0;
    max-width: calc(100% - 24px);
    text-align: center;
    word-break: break-word;
}
:deep(.msg-text-divider::before),
:deep(.msg-text-divider::after) {
    content: '';
    height: 2px;
    background-color: var(--text-muted);
    border-radius: 2px;
    flex-basis: 0;
}
:deep(.msg-text-divider::before) {
    flex-grow: var(--grow-left);
}
:deep(.msg-text-divider::after) {
    flex-grow: var(--grow-right);
}

:global(.app-layout.light-theme .msg-bubble:not(.is-mine) .msg-text) {
    color: #1a231e;
}

:global(.app-layout:not(.light-theme) .msg-bubble:not(.is-mine) .msg-text) {
    color: #c8c2b8;
}

:deep(.msg-image) { max-width: 100%; max-height: 250px; border-radius: 8px; margin-top: 8px; border: 1px solid var(--border); display: block; }
:deep(.msg-link) { color: var(--link-color); text-decoration: underline; text-underline-offset: 2px; }
:deep(.msg-link:hover) { color: var(--link-hover); }
</style>
