<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/pages/AppPage.vue
import { ref, computed, onMounted, onUnmounted, nextTick } from 'vue';
import { useRouter } from 'vue-router';
import { useUserStore } from '../stores/user';
import { fetchChatHistory } from '../api/auth';

const router = useRouter();
const userStore = useUserStore();

// === ТЕМА ===
const isDarkTheme = ref(true);
const toggleTheme = () => {
    isDarkTheme.value = !isDarkTheme.value;
};

const newMessage = ref('');
const messagesContainer = ref<HTMLElement | null>(null);
const ws = ref<WebSocket | null>(null);
const isConnected = ref(false);

const messages = ref<any[]>([]);

const activeTypingUsers = ref<string[]>([]);
const typingTimers = new Map<string, ReturnType<typeof setTimeout>>();

const typingText = computed(() => {
    const users = activeTypingUsers.value;
    const count = users.length;
    if (count === 0) return '';
    if (count === 1) return `${users[0]} печатает...`;
    if (count === 2) return `${users[0]} и ${users[1]} печатают...`;
    if (count === 3) return `${users[0]}, ${users[1]} и ${users[2]} печатают...`;
    if (count === 4) return `${users[0]}, ${users[1]}, ${users[2]} и ${users[3]} печатают...`;

    const remaining = count - 3;
    return `${users[0]}, ${users[1]}, ${users[2]} и ещё ${remaining} человек печатают...`;
});

const handleTypingEvent = (username: string) => {
    if (!username || username === userStore.currentUser?.username) return;

    if (!activeTypingUsers.value.includes(username)) {
        activeTypingUsers.value.push(username);
    }

    if (typingTimers.has(username)) {
        clearTimeout(typingTimers.get(username));
    }

    const timer = setTimeout(() => {
        activeTypingUsers.value = activeTypingUsers.value.filter(u => u !== username);
        typingTimers.delete(username);
    }, 3000);

    typingTimers.set(username, timer);
};

const isSidebarOpen = ref(window.innerWidth > 768);
const toggleSidebar = () => {
    isSidebarOpen.value = !isSidebarOpen.value;
};

const formatTime = (isoString: string) => {
    if (!isoString) return '';
    const date = new Date(isoString);
    return date.toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit', hour12: false });
};

const scrollToBottom = async () => {
    await nextTick();
    if (messagesContainer.value) {
        messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight;
    }
};

const formatMessage = (text: string) => {
    if (!text) return '';
    let safeText = text.replace(/[&<"'>]/g, (m) => {
        switch (m) {
            case '&': return '&amp;';
            case '<': return '&lt;';
            case '>': return '&gt;';
            case '"': return '&quot;';
            case "'": return '&#039;';
            default: return m;
        }
    });
    safeText = safeText.replace(/\n/g, '<br>');
    const urlRegex = /(https?:\/\/[^\s<]+)/g;
    safeText = safeText.replace(urlRegex, (url) => {
        if (url.match(/\.(jpeg|jpg|gif|png|webp)(\?.*)?$/i)) {
            return `<img src="${url}" class="msg-image" alt="Вложение" loading="lazy" />`;
        }
        return `<a href="${url}" target="_blank" class="msg-link">${url}</a>`;
    });
    return safeText;
};

const loadHistory = async () => {
    try {
        const history = await fetchChatHistory();
        messages.value = history.map((msg: any) => ({
            id: msg.id,
            user: msg.username,
            text: msg.text,
            time: formatTime(msg.timestamp)
        }));
        scrollToBottom();
    } catch (error) {
        console.error("Failed to load chat history:", error);
    }
};

const connectWebSocket = () => {
    const wsBaseUrl = (import.meta.env.VITE_API_URL || 'http://localhost:8080').replace(/^http/, 'ws');
    ws.value = new WebSocket(`${wsBaseUrl}/ws`);

    ws.value.onopen = () => { isConnected.value = true; };
    ws.value.onclose = () => { isConnected.value = false; setTimeout(connectWebSocket, 3000); };
    ws.value.onmessage = (event) => {
        try {
            const data = JSON.parse(event.data);

            if (data.type === 'typing') {
                console.log("[WS] Received typing event from:", data.username);
                handleTypingEvent(data.username);
                return;
            }

            messages.value.push({
                id: data.id || Date.now(),
                user: data.username,
                text: data.text,
                time: formatTime(data.timestamp || new Date().toISOString())
            });
            scrollToBottom();
        } catch (e) {
            console.error("Error parsing WS message:", e);
        }
    };
};

const sendMessage = () => {
    if (!newMessage.value.trim() || !ws.value || !isConnected.value) return;
    const msgData = { text: newMessage.value.trim() };
    ws.value.send(JSON.stringify(msgData));
    newMessage.value = '';
    scrollToBottom();
};

let lastTypingTime = 0;
const handleInput = () => {
    const now = Date.now();
    if (now - lastTypingTime > 2000) {
        if (ws.value && isConnected.value) {
            console.log("[WS] Sending typing event...");
            ws.value.send(JSON.stringify({ type: 'typing', username: userStore.currentUser?.username }));
        }
        lastTypingTime = now;
    }
};

const handleKeydown = (e: KeyboardEvent) => {
    if (e.key === 'Enter' && !e.shiftKey) {
        e.preventDefault();
        sendMessage();
    }
};

onMounted(() => {
    if (!userStore.currentUser) {
        router.push('/auth');
        return;
    }
    loadHistory();
    connectWebSocket();
});

onUnmounted(() => {
    if (ws.value) ws.value.close();
});
</script>

<template>
    <div class="app-layout" :class="{ 'light-theme': !isDarkTheme }">

        <div class="chat-header">
            <div class="header-left">
                <button class="channel-toggle-btn" @click="toggleSidebar">
                    <div class="channel-hash">
                        <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                            <line x1="4" y1="9" x2="20" y2="9"></line>
                            <line x1="4" y1="15" x2="20" y2="15"></line>
                            <line x1="10" y1="3" x2="8" y2="21"></line>
                            <line x1="16" y1="3" x2="14" y2="21"></line>
                        </svg>
                    </div>
                    <span class="channel-name">general</span>
                </button>
            </div>

            <div class="header-center brand">
                <img src="/Logo.svg" alt="Leaves" class="brand-logo" />
                <span class="brand-text">Leaves</span>
            </div>

            <div class="header-right header-status">
                <button class="theme-btn" @click="toggleTheme" title="Переключить тему">
                    <svg v-if="isDarkTheme" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <circle cx="12" cy="12" r="5"></circle>
                        <line x1="12" y1="1" x2="12" y2="3"></line>
                        <line x1="12" y1="21" x2="12" y2="23"></line>
                        <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line>
                        <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line>
                        <line x1="1" y1="12" x2="3" y2="12"></line>
                        <line x1="21" y1="12" x2="23" y2="12"></line>
                        <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line>
                        <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line>
                    </svg>
                    <svg v-else width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path>
                    </svg>
                </button>

                <span class="status-dot" :class="{ online: isConnected }"></span>
                <span class="status-text">{{ isConnected ? 'Connected' : 'Reconnecting...' }}</span>
            </div>
        </div>

        <div class="main-content">
            <div class="sidebar-overlay" :class="{ 'is-open': isSidebarOpen }" @click="toggleSidebar"></div>

            <div class="sidebar" :class="{ 'is-open': isSidebarOpen }">
                <div class="channels">
                    <div class="channel active"># general</div>
                </div>
            </div>

            <div class="chat-area">
                <div class="messages-container" ref="messagesContainer">
                    <div v-if="messages.length === 0" class="empty-state">No messages yet. Say hi!</div>

                    <div v-for="(msg, index) in messages" :key="msg.id"
                         class="message"
                         :style="{ marginBottom: (index < messages.length - 1 && messages[index + 1].user === msg.user) ? '4px' : '24px' }">

                        <div class="msg-avatar"
                             :style="{
                                 borderColor: userStore.getUserColor(msg.user),
                                 color: userStore.getUserColor(msg.user),
                                 visibility: (index > 0 && messages[index - 1].user === msg.user) ? 'hidden' : 'visible'
                             }">
                            {{ msg.user.charAt(0).toUpperCase() }}
                        </div>

                        <div class="msg-content">
                            <div class="msg-header" v-if="!(index > 0 && messages[index - 1].user === msg.user)">
                                <span class="msg-username" :style="{ color: userStore.getUserColor(msg.user) }">{{ msg.user }}</span>
                                <span class="msg-time">{{ msg.time }}</span>
                            </div>

                            <div class="msg-bubble" :class="{ 'is-mine': msg.user === userStore.currentUser?.username }">
                                <p class="msg-text" v-html="formatMessage(msg.text)"></p>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="chat-input-area">
                    <div class="typing-indicator" v-show="activeTypingUsers.length > 0">
                        <span class="typing-dots">
                            <span class="dot"></span><span class="dot"></span><span class="dot"></span>
                        </span>
                        <span class="typing-text">{{ typingText }}</span>
                    </div>

                    <div class="input-wrapper">
                        <textarea
                            v-model="newMessage"
                            @input="handleInput"
                            @keydown="handleKeydown"
                            placeholder="Message #general... (Shift+Enter for new line)"
                            class="chat-input"
                            rows="1"
                            maxlength="2000"
                        ></textarea>
                        <button @click="sendMessage" class="send-btn" :disabled="!newMessage.trim() || !isConnected">
                            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="22" y1="2" x2="11" y2="13"></line><polygon points="22 2 15 22 11 13 2 9 22 2"></polygon></svg>
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
/* ПЕРЕМЕННЫЕ ТЕМ */
.app-layout {
    /* ТЁМНАЯ ТЕМА (ПО УМОЛЧАНИЮ) */
    --bg-base: #050807;
    --bg-surface: #080b0a;
    --border: #0f1714;
    --bg-active: #1a231e;
    --border-active: #233129;
    --text-primary: #c8c2b8;
    --text-secondary: #64615c;
    --text-muted: #8a867f;
    --accent: #5fca08;
    --accent-hover: #4da806;
    --link-color: #5fca08;
    --link-hover: #88ffb4;

    display: flex;
    flex-direction: column;
    height: 100dvh;
    background-color: var(--bg-base);
    color: var(--text-primary);
    overflow: hidden;
    transition: background-color 0.3s ease, color 0.3s ease;
}

/* СВЕТЛАЯ НЕОНОВАЯ ТЕМА */
.app-layout.light-theme {
    --bg-base: #f4f7f6;
    --bg-surface: #ffffff;
    --border: #e2e8e4;
    --bg-active: #e8f5e9;
    --border-active: #c8e6c9;
    --text-primary: #1a231e;
    --text-secondary: #6b7280;
    --text-muted: #4b5563;
    --accent: #5fca08; 
    --accent-hover: #4da806;
    --link-color: #3d8c00;
    --link-hover: #5fca08;
}

/* Плавные переходы для всех элементов */
.chat-header, .sidebar, .input-wrapper, .msg-bubble, .msg-avatar, .channel, .chat-input, .channel-toggle-btn {
    transition: background-color 0.3s ease, border-color 0.3s ease, color 0.3s ease, box-shadow 0.3s ease, filter 0.3s ease;
}

/* Общий Header */
.chat-header {
    height: 70px;
    padding: 0 24px;
    border-bottom: 1px solid var(--border);
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: var(--bg-surface);
    z-index: 10;
    flex-shrink: 0;
}
.header-left, .header-right { flex: 1; display: flex; align-items: center; }
.header-right { justify-content: flex-end; }
.header-center { flex: 1; display: flex; justify-content: center; align-items: center; gap: 10px; }

.channel-toggle-btn {
    background: transparent;
    border: none;
    display: flex;
    align-items: center;
    gap: 6px;
    cursor: pointer;
    padding: 6px 10px;
    border-radius: 8px;
    color: var(--text-primary);
    margin-left: -10px;
}
.channel-toggle-btn:hover {
    background-color: var(--border);
}
.channel-hash {
    color: var(--text-secondary);
    display: flex;
    align-items: center;
    justify-content: center;
}
.channel-name {
    font-size: 17px;
    font-weight: 600;
    letter-spacing: 0.3px;
}
.brand-logo { width: 28px; height: 28px; }
.brand-text { font-size: 20px; font-weight: 700; color: var(--accent); }

/* ======== КНОПКА СМЕНЫ ТЕМЫ ======== */
.theme-btn {
    background: transparent;
    border: none;
    cursor: pointer;
    margin-right: 15px;
    padding: 6px;
    border-radius: 8px;
    transition: background 0.2s, color 0.2s, transform 0.1s;
    display: flex;
    align-items: center;
    justify-content: center;
    color: var(--text-secondary); /* По умолчанию берет серый цвет из темы */
}

.theme-btn:hover {
    background: var(--border);
    color: var(--accent); /* При наведении берет зеленый цвет из темы */
    transform: scale(1.05); 
}

.theme-btn:active,
.theme-btn:focus {
    color: var(--accent); /* При фокусе/нажатии зеленый цвет */
    transform: scale(0.95); 
}
/* =================================== */

.header-status { gap: 8px; font-size: 13px; color: var(--text-secondary); }
.status-dot { width: 8px; height: 8px; border-radius: 50%; background-color: #dc3545; }
.status-dot.online { background-color: var(--accent); box-shadow: 0 0 8px rgba(95, 202, 8, 0.4); }

.main-content {
    display: flex;
    flex: 1;
    min-height: 0;
    position: relative;
}

.sidebar {
    width: 260px;
    background-color: var(--bg-surface);
    border-right: 1px solid var(--border);
    display: flex;
    flex-direction: column;
    flex-shrink: 0;
    transition: margin-left 0.3s cubic-bezier(0.4, 0, 0.2, 1), background-color 0.3s, border-color 0.3s;
}
.sidebar:not(.is-open) {
    margin-left: -261px;
}
.channels { padding: 15px 10px; flex: 1; overflow-y: auto; }
.channel { padding: 10px 15px; border-radius: 8px; color: var(--text-muted); cursor: pointer; font-weight: 500; }
.channel:hover { background-color: var(--border); color: var(--text-primary); }
.channel.active { background-color: var(--bg-active); color: var(--accent); }

.sidebar-overlay { display: none; }

.chat-area { flex: 1; display: flex; flex-direction: column; background-color: var(--bg-base); min-width: 0; }
.messages-container { flex: 1; padding: 24px; overflow-y: auto; display: flex; flex-direction: column; scroll-behavior: smooth; }
.empty-state { text-align: center; color: var(--text-secondary); margin-top: auto; margin-bottom: auto; font-size: 15px; }

.message { display: flex; gap: 16px; align-items: flex-start; max-width: 100%; transition: margin 0.2s ease; }
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

/* Фикс для ников и аватарок в светлой теме, чтобы ядовитые цвета читались на белом */
.app-layout.light-theme .msg-username,
.app-layout.light-theme .msg-avatar {
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

/* Общий текст сообщения */
.msg-text { 
    font-size: 15px; 
    line-height: 1.5; 
    color: var(--text-primary);
    margin: 0; 
    word-break: break-word; 
}

/* Цвет чужих сообщений жестко привязан к читаемым цветам в зависимости от темы */
.app-layout.light-theme .msg-bubble:not(.is-mine) .msg-text {
    color: #1a231e;
    font-weight: 500;
}

.app-layout:not(.light-theme) .msg-bubble:not(.is-mine) .msg-text {
    color: #c8c2b8;
}

:deep(.msg-image) { max-width: 100%; max-height: 250px; border-radius: 8px; margin-top: 8px; border: 1px solid var(--border); display: block; }
:deep(.msg-link) { color: var(--link-color); text-decoration: underline; text-underline-offset: 2px; }
:deep(.msg-link:hover) { color: var(--link-hover); }

.chat-input-area {
    flex-shrink: 0;
    padding: 0 24px 24px 24px;
    padding-bottom: calc(24px + env(safe-area-inset-bottom));
    background-color: var(--bg-base);
}
.input-wrapper { background-color: var(--bg-surface); border: 1px solid var(--border); border-radius: 16px; display: flex; align-items: flex-end; padding: 12px 20px; }
.input-wrapper:focus-within { border-color: rgba(95, 202, 8, 0.4); }

/* Поле ввода */
.chat-input { 
    flex: 1; 
    background: transparent; 
    border: none; 
    color: var(--text-primary) !important;
    font-size: 15px; 
    outline: none; 
    resize: none; 
    max-height: 120px; 
    font-family: inherit; 
    line-height: 1.5; 
    padding-top: 2px; 
}
.chat-input::placeholder { color: var(--text-secondary) !important; }

.send-btn { background: var(--accent); color: #ffffff; border: none; width: 36px; height: 36px; border-radius: 10px; display: flex; justify-content: center; align-items: center; cursor: pointer; transition: all 0.2s; flex-shrink: 0; margin-left: 12px; }
.send-btn:hover:not(:disabled) { background: var(--accent-hover); transform: scale(1.05); }
.send-btn:disabled { background: var(--bg-active); color: var(--text-secondary); cursor: not-allowed; }

.typing-indicator {
    height: 20px;
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 0 4px;
    margin-bottom: 8px;
    color: #64615c;
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
    background-color: #64615c;
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
    color: #8a867f;
}

/* ======== МОБИЛЬНАЯ АДАПТАЦИЯ ======== */
@media (max-width: 768px) {
    .status-text { display: none; }
    .chat-header { padding: 0 16px; }

    .sidebar {
        position: fixed;
        top: 70px;
        left: 0;
        bottom: 0;
        z-index: 100;
        margin-left: 0 !important;
        transform: translateX(-100%);
        transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
        box-shadow: 5px 0 15px rgba(0,0,0,0.5);
    }
    .app-layout.light-theme .sidebar {
        box-shadow: 5px 0 15px rgba(0,0,0,0.1);
    }
    .sidebar.is-open {
        transform: translateX(0);
    }

    .sidebar-overlay {
        display: block;
        position: fixed;
        top: 70px; left: 0; right: 0; bottom: 0;
        background: rgba(0, 0, 0, 0.6);
        z-index: 99;
        opacity: 0;
        pointer-events: none;
        transition: opacity 0.3s ease;
    }
    .app-layout.light-theme .sidebar-overlay {
        background: rgba(255, 255, 255, 0.6);
        backdrop-filter: blur(2px);
    }
    .sidebar-overlay.is-open {
        opacity: 1;
        pointer-events: auto;
    }

    .messages-container { padding: 16px; }
    .chat-input-area { padding: 0 16px 16px 16px; padding-bottom: calc(16px + env(safe-area-inset-bottom)); }
}
</style>