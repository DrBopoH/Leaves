<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ChatMessage.vue
import { useUserStore } from '../stores/user';
import UiMessageBubble from './ui-kit/UiMessageBubble.vue';

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

    safeText = safeText.replace(/^(---+) *\r?$/gm, (match, p1) => {
        let pct = Math.min(100, p1.length * 10);
        return `<hr class="msg-divider" style="width: ${pct}%" />`;
    });

    safeText = safeText.replace(/^(---+)(.+?)(---+) *\r?$/gm, (match, p1, p2, p3) => {
        return `<div class="msg-text-divider"><div class="msg-text-divider-line" style="flex-grow: ${p1.length}"></div><span>${p2.trim()}</span><div class="msg-text-divider-line" style="flex-grow: ${p3.length}"></div></div>`;
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
    <UiMessageBubble
        :username="message.user"
        :color="userStore.getUserColor(message.user)"
        :time="message.time"
        :is-mine="isMine"
        :show-header="showHeader"
    >
        <p class="msg-text" v-html="formatMessage(message.text)"></p>
    </UiMessageBubble>
</template>

<style scoped>
</style>
