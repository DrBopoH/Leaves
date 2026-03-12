<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/pages/AppPage.vue
import { ref, computed, onMounted, onUnmounted, nextTick, watch } from 'vue';
import { useRouter } from 'vue-router';
import { useUserStore } from '../stores/user';
import { fetchChatHistory, fetchUsers } from '../api/auth';
import { useTheme } from '../composables/useTheme';
import ChatMessage from '../components/ChatMessage.vue';
import UsersList from '../components/UsersList.vue';
import AppHeader from '../components/AppHeader.vue';
import ChatInput from '../components/ChatInput.vue';

const router = useRouter();
const userStore = useUserStore();

const { isDark: isDarkTheme, toggleTheme } = useTheme();

const newMessage = ref('');
const messagesContainer = ref<HTMLElement | null>(null);
const ws = ref<WebSocket | null>(null);
const isConnected = ref(false);

const messages = ref<any[]>([]);

const activeTypingUsers = ref<string[]>([]);
const typingTimers = new Map<string, ReturnType<typeof setTimeout>>();

const allUsers = ref<{username: string, online: boolean, last_seen?: string}[]>([]);

const generateAndSortUsers = async () => {
    try {
        const users = await fetchUsers();
        users.sort((a: any, b: any) => {
            if(a.online && !b.online) return -1;
            if(!a.online && b.online) return 1;
            return a.username.localeCompare(b.username);
        });
        allUsers.value = users;
    } catch (e) {
        console.error("Failed to load users:", e);
    }
};

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

const isRightSidebarOpen = ref(false);
const toggleRightSidebar = () => {
    isRightSidebarOpen.value = !isRightSidebarOpen.value;
};

const channels = ref(['general']);
const activeChannel = ref('general');
const selectChannel = (channel: string) => {
    activeChannel.value = channel;
};
const handleChannelDblClick = (channel: string) => {
    activeChannel.value = channel;
    isSidebarOpen.value = false;
};

const showStatusText = ref(false);
let statusTimer: ReturnType<typeof setTimeout> | null = null;
watch(isConnected, (newVal) => {
    showStatusText.value = true;
    if (statusTimer) clearTimeout(statusTimer);
    statusTimer = setTimeout(() => {
        showStatusText.value = false;
    }, 3000);
});

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
                handleTypingEvent(data.username);
                return;
            }
            if (data.type === 'status') {
                const userIndex = allUsers.value.findIndex(u => u.username === data.username);
                if (userIndex !== -1) {
                    allUsers.value[userIndex].online = data.online;
                    allUsers.value[userIndex].last_seen = data.last_seen;

                    allUsers.value.sort((a: any, b: any) => {
                        if(a.online && !b.online) return -1;
                        if(!a.online && b.online) return 1;
                        return a.username.localeCompare(b.username);
                    });
                }
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

    let processedText = newMessage.value.trim();

    processedText = processedText.replace(/(^|[^\\])\\n/g, '$1\n');
    processedText = processedText.replace(/(^|[^\\])\\t/g, '$1\t');

    processedText = processedText.replace(/\\\\n/g, '\\n');
    processedText = processedText.replace(/\\\\t/g, '\\t');

    const msgData = { text: processedText };
    ws.value.send(JSON.stringify(msgData));

    newMessage.value = '';

    scrollToBottom();
};

let lastTypingTime = 0;
const emitTypingEvent = () => {
    const now = Date.now();
    if (now - lastTypingTime > 2000) {
        if (ws.value && isConnected.value) {
            console.log("[WS] Sending typing event...");
            ws.value.send(JSON.stringify({ type: 'typing', username: userStore.currentUser?.username }));
        }
        lastTypingTime = now;
    }
};

onMounted(() => {
    if (!userStore.currentUser) {
        router.push('/auth');
        return;
    }
    generateAndSortUsers();
    loadHistory();
    connectWebSocket();
});

onUnmounted(() => {
    if (ws.value) ws.value.close();
});
</script>

<template>
    <div class="app-layout" :class="{ 'light-theme': !isDarkTheme }">

        <AppHeader />

        <div class="main-content">
            <div class="sidebar-overlay" :class="{ 'is-open': isSidebarOpen }" @click="toggleSidebar"></div>
            <div class="sidebar-overlay right-overlay" :class="{ 'is-open': isRightSidebarOpen }" @click="toggleRightSidebar"></div>

            <div class="sidebar" :class="{ 'is-open': isSidebarOpen }">
                <div class="channels">
                    <div v-for="chan in channels" :key="chan"
                         class="channel"
                         :class="{ active: activeChannel === chan }"
                         @click="selectChannel(chan)"
                         @dblclick="handleChannelDblClick(chan)">
                         # {{ chan }}
                    </div>
                </div>

                <div class="sidebar-footer">
                    <div class="user-profile">
                        <div class="avatar" :style="{
                            borderColor: userStore.getUserColor(userStore.currentUser?.username || ''),
                            color: userStore.getUserColor(userStore.currentUser?.username || '')
                        }">
                            {{ userStore.currentUser?.username.charAt(0).toUpperCase() }}
                        </div>
                        <div class="user-info">
                            <span class="username" :style="{ color: userStore.getUserColor(userStore.currentUser?.username || '') }">
                                {{ userStore.currentUser?.username }}
                            </span>
                            <div class="status-indicator">
                                <span class="status-dot" :class="{ online: isConnected }"></span>
                                <transition name="fade-status">
                                    <span class="status-text" v-if="showStatusText">{{ isConnected ? 'Connected' : 'Reconnecting...' }}</span>
                                </transition>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <div class="chat-container">
                <div class="chat-inner-header">
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
                            <span class="channel-name">{{ activeChannel }}</span>
                        </button>
                    </div>

                    <div class="header-right">
                        <button class="channel-toggle-btn online-users-btn" @click="toggleRightSidebar">
                            <svg class="online-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                                <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
                                <circle cx="9" cy="7" r="4"></circle>
                                <path d="M23 21v-2a4 4 0 0 0-3-3.87"></path>
                                <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
                            </svg>
                            <span class="online-text">В сети: {{ allUsers.filter(u => u.online).length }}</span>
                        </button>
                    </div>
                </div>

                <div class="chat-area">
                    <div class="messages-container" ref="messagesContainer">
                        <div v-if="messages.length === 0" class="empty-state">Пока нет сообщений. Поздоровайтесь!</div>

                        <ChatMessage
                            v-for="(msg, index) in messages"
                            :key="msg.id"
                            :message="msg"
                            :is-mine="msg.user === userStore.currentUser?.username"
                            :show-header="!(index > 0 && messages[index - 1].user === msg.user)"
                            :style="{ marginBottom: (index < messages.length - 1 && messages[index + 1].user === msg.user) ? '4px' : '24px' }"
                        />
                    </div>

                    <ChatInput
                        v-model="newMessage"
                        :is-connected="isConnected"
                        :active-typing-users="activeTypingUsers"
                        :typing-text="typingText"
                        @send="sendMessage"
                        @typing="emitTypingEvent"
                    />
                </div>
            </div>

            <div class="right-sidebar" :class="{ 'is-open': isRightSidebarOpen }">
                <div class="right-sidebar-header">
                    <span class="right-sidebar-title">В сети</span>
                    <button class="close-sidebar-btn" @click="toggleRightSidebar">
                        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <line x1="18" y1="6" x2="6" y2="18"></line>
                            <line x1="6" y1="6" x2="18" y2="18"></line>
                        </svg>
                    </button>
                </div>
                <UsersList :users="allUsers" />
            </div>
        </div>
    </div>
</template>

<style scoped>
.app-layout {
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
    /* Убрали тяжелую анимацию всех элементов, оставили только плавную смену фона и цвета */
    transition: background-color 0.2s ease, color 0.2s ease;
}

.app-layout.light-theme {
    --bg-base: #f4f7f6;
    --bg-surface: #ebf0ee;
    --border: #e2e8e4;
    --bg-active: #e2e8e4;
    --border-active: #c8e6c9;
    --text-primary: #080b0a;
    --text-secondary: #4b5563;
    --text-muted: #6b7280;
    --accent: #4da806;
    --accent-hover: #3d8c00;
    --link-color: #3d8c00;
    --link-hover: #5fca08;
}

/* Прицельная анимация только для блоков, которым она реально нужна */
.sidebar, .right-sidebar, .chat-inner-header, .user-profile, .channel {
    transition: background-color 0.2s ease, border-color 0.2s ease, color 0.2s ease, transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

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
    transition: background-color 0.2s ease;
}
.channel-toggle-btn:hover { background-color: var(--border); }
.channel-hash { color: var(--text-secondary); display: flex; align-items: center; justify-content: center; }
.channel-name { font-size: 17px; font-weight: 600; letter-spacing: 0.3px; }

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
}
.sidebar:not(.is-open) { margin-left: -261px; }

.channels { padding: 10px 10px; flex: 1; overflow-y: auto; }
.channel { padding: 10px 15px; border-radius: 8px; color: var(--text-muted); cursor: pointer; font-weight: 500; user-select: none; }
.channel:hover { background-color: var(--border); color: var(--text-primary); }
.channel.active { background-color: var(--bg-active); color: var(--accent); }

.sidebar-footer {
    padding: 12px 8px;
    padding-bottom: calc(8px + env(safe-area-inset-bottom));
    border-top: 1px solid var(--border);
    background-color: transparent;
    flex-shrink: 0;
}
.user-profile {
    display: flex;
    align-items: center;
    gap: 10px;
    background-color: var(--bg-active);
    border: 1px solid var(--border);
    border-radius: 14px;
    padding: 12px 12px;
}
.avatar {
    width: 28px;
    height: 28px;
    border-radius: 50%;
    border: 2px solid;
    display: flex;
    align-items: center;
    justify-content: center;
    font-weight: bold;
    font-size: 14px;
    background-color: var(--bg-surface);
    flex-shrink: 0;
}
.user-info { display: flex; flex-direction: column; min-width: 0; }
.username { font-size: 14px; font-weight: 600; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }
.status-indicator { display: flex; align-items: center; gap: 6px; height: 14px; margin-top: 2px; }
.status-dot { width: 8px; height: 8px; border-radius: 50%; background-color: #dc3545; flex-shrink: 0;}
.status-dot.online { background-color: var(--accent); box-shadow: 0 0 8px rgba(95, 202, 8, 0.4); }
.status-text { font-size: 11px; color: var(--text-secondary); white-space: nowrap; }
.fade-status-enter-active, .fade-status-leave-active { transition: opacity 0.5s ease; }
.fade-status-enter-from, .fade-status-leave-to { opacity: 0; }

.sidebar-overlay { display: none; }

.chat-container {
    flex: 1;
    display: flex;
    flex-direction: column;
    min-width: 0;
    background-color: transparent;
}
.chat-inner-header {
    height: 60px;
    padding: 0 24px;
    border-bottom: 1px solid var(--border);
    display: flex;
    align-items: center;
    justify-content: space-between;
    background-color: var(--bg-base);
    flex-shrink: 0;
}

.online-users-btn { margin-left: 0; }
.online-icon { color: var(--text-secondary); }
.online-users-btn:hover .online-icon { color: var(--text-primary); }
.online-text { font-size: 14px; font-weight: 600; }

.chat-area { 
    flex: 1; 
    display: flex; 
    flex-direction: column; 
    min-height: 0; 
    position: relative;
    /* Учитываем зону безопасности (челки и полоски навигации iOS/Android) */
    padding-bottom: env(safe-area-inset-bottom);
}

.messages-container {
    flex: 1;
    padding: 8px 12px;
    padding-bottom: 72px; /* Место под инпут */
    overflow-y: auto;
    display: flex;
    flex-direction: column;
    scroll-behavior: smooth;
}
.empty-state { text-align: center; color: var(--text-secondary); margin-top: auto; margin-bottom: auto; font-size: 15px; }

::-webkit-scrollbar { width: 6px; height: 6px; }
::-webkit-scrollbar-track { background: transparent; }
::-webkit-scrollbar-thumb { background: var(--border-active); border-radius: 3px; }
::-webkit-scrollbar-thumb:hover { background: var(--text-muted); }

.right-sidebar {
    width: 240px;
    background-color: var(--bg-surface);
    border-left: 1px solid var(--border);
    display: flex;
    flex-direction: column;
    flex-shrink: 0;
}
.right-sidebar:not(.is-open) { margin-right: -241px; }

/* Скрываем шапку правой панели на ПК */
.right-sidebar-header {
    display: none;
    padding: 16px;
    border-bottom: 1px solid var(--border);
    align-items: center;
    justify-content: space-between;
}
.right-sidebar-title { font-weight: 600; font-size: 16px; color: var(--text-primary); }
.close-sidebar-btn {
    background: transparent;
    border: none;
    color: var(--text-secondary);
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 4px;
}
.close-sidebar-btn:hover { color: var(--text-primary); }

@media (max-width: 768px) {
    .chat-inner-header { padding: 0 16px; }

    .sidebar {
        position: fixed;
        top: 48px; left: 0; bottom: 0;
        z-index: 100;
        margin-left: 0 !important;
        transform: translateX(-100%);
        box-shadow: 5px 0 15px rgba(0,0,0,0.5);
    }
    .right-sidebar {
        position: fixed;
        top: 48px; right: 0; bottom: 0;
        z-index: 100;
        margin-right: 0 !important;
        transform: translateX(100%);
        box-shadow: -5px 0 15px rgba(0,0,0,0.5);
    }
    
    .right-sidebar-header { display: flex; } /* Показываем шапку с крестиком на мобилке */

    .app-layout.light-theme .sidebar, .app-layout.light-theme .right-sidebar {
        box-shadow: 5px 0 15px rgba(0,0,0,0.1);
    }
    .app-layout.light-theme .right-sidebar {
        box-shadow: -5px 0 15px rgba(0,0,0,0.1);
    }
    .sidebar.is-open { transform: translateX(0); }
    .right-sidebar.is-open { transform: translateX(0); }

    .sidebar-overlay {
        display: block;
        position: fixed;
        top: 48px; left: 0; right: 0; bottom: 0;
        background: rgba(0, 0, 0, 0.6);
        z-index: 99;
        opacity: 0;
        pointer-events: none;
        transition: opacity 0.3s ease, background-color 0.3s ease;
    }
    
    /* Скрываем overlay правой панели, если мы на десктопе, чтобы не перекрывало чат */
    .right-overlay:not(.is-open) { display: none; }
    
    .app-layout.light-theme .sidebar-overlay {
        background: rgba(255, 255, 255, 0.6);
        backdrop-filter: blur(2px);
    }
    .sidebar-overlay.is-open {
        display: block;
        opacity: 1;
        pointer-events: auto;
    }

    .messages-container { padding: 16px;padding-bottom: 64px; }
}
</style>