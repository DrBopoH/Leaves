<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/pages/AppPage.vue
import { ref, computed, onMounted, onUnmounted, nextTick, watch } from 'vue';
import { useRouter } from 'vue-router';
import { useUserStore } from '../stores/user';
import { fetchChatHistory, fetchUsers } from '../api/auth';
import { useTheme } from '../composables/useTheme';

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
const visibleUsers = ref<{username: string, online: boolean, last_seen?: string}[]>([]);
const usersPage = ref(1);
const isLoadingUsers = ref(false);
let loadUsersTimeout: ReturnType<typeof setTimeout> | null = null;
let lastLoadTime = 0;
const PAGE_SIZE = 30;

const generateAndSortUsers = async () => {
    try {
        const users = await fetchUsers();
        users.sort((a: any, b: any) => {
            if(a.online && !b.online) return -1;
            if(!a.online && b.online) return 1;
            return a.username.localeCompare(b.username);
        });
        allUsers.value = users;
        updateVisibleUsers();
    } catch (e) {
        console.error("Failed to load users:", e);
    }
};

const updateVisibleUsers = () => {
    const end = usersPage.value * PAGE_SIZE;
    const start = Math.max(0, end - 90); // Keep max 90 users in DOM
    visibleUsers.value = allUsers.value.slice(start, end);
};

const handleUsersScroll = (e: Event) => {
    const target = e.target as HTMLElement;

    if (target.scrollHeight - target.scrollTop <= target.clientHeight + 10) {
        if (usersPage.value * PAGE_SIZE < allUsers.value.length && !isLoadingUsers.value) {
            const now = Date.now();
            const timeSinceLastLoad = now - lastLoadTime;
            const delay = timeSinceLastLoad > 800 ? 0 : 800 - timeSinceLastLoad;

            isLoadingUsers.value = true;

            if (loadUsersTimeout) clearTimeout(loadUsersTimeout);
            loadUsersTimeout = setTimeout(() => {
                usersPage.value++;
                updateVisibleUsers();
                isLoadingUsers.value = false;
                lastLoadTime = Date.now();
            }, delay);
        }
    }

    if (target.scrollTop <= 10 && usersPage.value > 3) {
        usersPage.value--;
        updateVisibleUsers();
        target.scrollTop = 50;
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

const formatRelativeTime = (isoString: string | undefined) => {
    if (!isoString) return 'Long ago';
    const date = new Date(isoString);
    const diffSeconds = Math.floor((Date.now() - date.getTime()) / 1000);

    if (diffSeconds < 60) return 'Just now';
    if (diffSeconds < 3600) return `${Math.floor(diffSeconds / 60)} mins ago`;
    if (diffSeconds < 86400) return `${Math.floor(diffSeconds / 3600)} hours ago`;
    return `${Math.floor(diffSeconds / 86400)} days ago`;
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
            case '&': return '&';
            case '<': return '<';
            case '>': return '>';
            case '"': return '"';
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
                    updateVisibleUsers();
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
    const msgData = { text: newMessage.value.trim() };
    ws.value.send(JSON.stringify(msgData));
    newMessage.value = '';
    scrollToBottom();
};

let lastTypingTime = 0;
const handleInput = (e: Event) => {
    const now = Date.now();
    if (now - lastTypingTime > 2000) {
        if (ws.value && isConnected.value) {
            console.log("[WS] Sending typing event...");
            ws.value.send(JSON.stringify({ type: 'typing', username: userStore.currentUser?.username }));
        }
        lastTypingTime = now;
    }

    const target = e.target as HTMLTextAreaElement;
    target.style.height = 'auto';
    target.style.height = Math.min(target.scrollHeight, 120) + 'px';
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

        <div class="top-header">
            <div class="header-left brand">
            </div>

            <div class="header-center brand" @click="router.push('/')" style="cursor: pointer; user-select: none;" title="На главную">
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
            </div>
        </div>

        <div class="main-content">
            <div class="sidebar-overlay" :class="{ 'is-open': isSidebarOpen }" @click="toggleSidebar"></div>
            <div class="sidebar-overlay right-overlay" :class="{ 'is-open': isRightSidebarOpen && windowWidth <= 768 }" @click="toggleRightSidebar"></div>

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
                            <span class="online-text">Online: {{ allUsers.filter(u => u.online).length }}</span>
                        </button>
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
                            <span class="typing-text">
                                {{ typingText }}
                            </span>
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
                                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="22" y1="2" x2="11" y2="13"></line><polygon points="22 2 15 22 11 13 2 9 22 2"></polygon></svg>
                            </button>
                        </div>
                    </div>
				</div>
			</div>

			<div class="right-sidebar" :class="{ 'is-open': isRightSidebarOpen }">
                <div class="members-list" @scroll="handleUsersScroll">
                    <div v-for="u in visibleUsers" :key="u.username" class="user-profile member-item">
                        <div class="avatar" :style="{
                            borderColor: userStore.getUserColor(u.username),
                            color: userStore.getUserColor(u.username)
                        }">
                            {{ u.username.charAt(0).toUpperCase() }}
                        </div>
                        <div class="user-info">
                            <span class="username" :style="{ color: userStore.getUserColor(u.username) }">
                                {{ u.username }}
                            </span>
                            <div class="status-indicator">
                                <span class="status-dot" :class="{ online: u.online }"></span>
                                <span class="status-text">{{ u.online ? 'Online' : formatRelativeTime(u.last_seen) }}</span>
                            </div>
                        </div>
                    </div>

                    <div v-if="isLoadingUsers" class="loading-users">
                        <span class="dot"></span><span class="dot"></span><span class="dot"></span>
                    </div>
                </div>
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
    transition: background-color 0.3s ease, color 0.3s ease;
}

.app-layout *, .app-layout *::before, .app-layout *::after {
    transition-duration: 0.3s;
    transition-timing-function: ease;
    transition-property: background-color, border-color, color, fill, stroke, box-shadow;
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

.top-header {
    height: 36px;
    padding: 0 16px;
    border-bottom: 1px solid var(--border);
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: var(--bg-surface);
    z-index: 10;
    flex-shrink: 0;
}
.header-left, .header-right { display: flex; align-items: center; gap: 10px; flex: 1; }
.header-right { justify-content: flex-end; }
.header-center { display: flex; justify-content: center; align-items: center; }

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
.brand-logo { width: 22px; height: 22px; }
.brand-text {
    font-size: 15px;
    font-weight: 700;
    color: var(--text-secondary);
}
.header-center.brand {
    display: flex;
    align-items: center;
    gap: 8px;
}

.theme-btn {
    background: transparent;
    border: none;
    cursor: pointer;
    margin-right: 5px;
    padding: 6px;
    border-radius: 8px;
    transition: background 0.2s, color 0.2s, transform 0.1s;
    display: flex;
    align-items: center;
    justify-content: center;
    color: var(--text-secondary);
}

.theme-btn:hover {
    background: var(--border);
    color: var(--accent);
    transform: scale(1.05);
}

.theme-btn:active,
.theme-btn:focus {
    color: var(--accent);
    transform: scale(0.95);
}

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
    transition-property: margin-left, background-color, border-color, color;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1), ease, ease, ease;
}
.sidebar:not(.is-open) {
    margin-left: -261px;
}
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
    transition-property: background-color, border-color, box-shadow;
    transition-duration: 0.3s;
}
.user-profile .avatar {
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
.user-info {
    display: flex;
    flex-direction: column;
    min-width: 0;
}
.user-info .username {
    font-size: 14px;
    font-weight: 600;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}
.status-indicator {
    display: flex;
    align-items: center;
    gap: 6px;
    height: 14px;
    margin-top: 2px;
}
.status-dot { width: 8px; height: 8px; border-radius: 50%; background-color: #dc3545; flex-shrink: 0;}
.status-dot.online { background-color: var(--accent); box-shadow: 0 0 8px rgba(95, 202, 8, 0.4); }
.status-indicator .status-text {
    font-size: 11px;
    color: var(--text-secondary);
    white-space: nowrap;
}
.fade-status-enter-active, .fade-status-leave-active {
    transition: opacity 0.5s ease;
}
.fade-status-enter-from, .fade-status-leave-to {
    opacity: 0;
}

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

.online-users-btn {
    margin-left: 0;
}
.online-icon {
    color: var(--text-secondary);
}
.online-users-btn:hover .online-icon {
    color: var(--text-primary);
}
.online-text {
    font-size: 14px;
    font-weight: 600;
}

.chat-area { flex: 1; display: flex; flex-direction: column; min-height: 0; position: relative; }
.messages-container {
    flex: 1;
    padding: 8px 12px;
    padding-bottom: 72px;
    overflow-y: auto;
    display: flex;
    flex-direction: column;
    scroll-behavior: smooth;
}
.empty-state { text-align: center; color: var(--text-secondary); margin-top: auto; margin-bottom: auto; font-size: 15px; }

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

.msg-text {
    font-size: 15px;
    line-height: 1.5;
    color: var(--text-primary);
    margin: 0;
    word-break: break-word;
}

.app-layout.light-theme .msg-bubble:not(.is-mine) .msg-text {
    color: #1a231e;
}

.app-layout:not(.light-theme) .msg-bubble:not(.is-mine) .msg-text {
    color: #c8c2b8;
}

:deep(.msg-image) { max-width: 100%; max-height: 250px; border-radius: 8px; margin-top: 8px; border: 1px solid var(--border); display: block; }
:deep(.msg-link) { color: var(--link-color); text-decoration: underline; text-underline-offset: 2px; }
:deep(.msg-link:hover) { color: var(--link-hover); }

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

::-webkit-scrollbar {
    width: 6px;
    height: 6px;
}
::-webkit-scrollbar-track {
    background: transparent;
}
::-webkit-scrollbar-thumb {
    background: var(--border-active);
    border-radius: 3px;
}
::-webkit-scrollbar-thumb:hover {
    background: var(--text-muted);
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
    background-color: var(--color-text-secondary);
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
    color: var(--color-text-secondary);
}

.right-sidebar {
    width: 240px;
    background-color: var(--bg-surface);
    border-left: 1px solid var(--border);
    display: flex;
    flex-direction: column;
    flex-shrink: 0;
    transition-property: margin-right, background-color, border-color, color;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1), ease, ease, ease;
}
.right-sidebar:not(.is-open) {
    margin-right: -241px;
}
.members-list {
    flex: 1;
    overflow-y: auto;
    padding: 16px;
    display: flex;
    flex-direction: column;
    gap: 12px;
}
.member-item {
    cursor: pointer;
    padding: 6px 8px;
    border-radius: 8px;
    transition-property: background-color;
    transition-duration: 0.2s;
}
.member-item:hover {
    background-color: var(--border);
}

.loading-users {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 15px;
    gap: 4px;
}
.loading-users .dot {
    width: 6px;
    height: 6px;
    background-color: var(--text-secondary);
    border-radius: 50%;
    animation: bounce 1.4s infinite ease-in-out both;
}
.loading-users .dot:nth-child(1) { animation-delay: -0.32s; }
.loading-users .dot:nth-child(2) { animation-delay: -0.16s; }

@keyframes bounce {
    0%, 80%, 100% { transform: scale(0); }
    40% { transform: scale(1); }
}

@media (max-width: 768px) {
    .chat-inner-header { padding: 0 16px; }

    .sidebar {
        position: fixed;
        top: 48px;
        left: 0;
        bottom: 0;
        z-index: 100;
        margin-left: 0 !important;
        transform: translateX(-100%);
        box-shadow: 5px 0 15px rgba(0,0,0,0.5);
    }
    .right-sidebar {
        position: fixed;
        top: 48px;
        right: 0;
        bottom: 0;
        z-index: 100;
        margin-right: 0 !important;
        transform: translateX(100%);
        box-shadow: -5px 0 15px rgba(0,0,0,0.5);
    }
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
        transition-property: opacity, background-color, backdrop-filter;
        transition-duration: 0.3s;
    }
    .app-layout.light-theme .sidebar-overlay {
        background: rgba(255, 255, 255, 0.6);
        backdrop-filter: blur(2px);
    }
    .sidebar-overlay.is-open {
        opacity: 1;
        pointer-events: auto;
    }
    .right-overlay { z-index: 99; }

    .messages-container { padding: 16px; }
    .chat-input-area {
		padding: 0 16px 16px 16px;
		padding-bottom: calc(16px + env(safe-area-inset-bottom));
	}
}
</style>
