<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/pages/AppPage.vue
import { ref, computed, onMounted, onUnmounted, nextTick, watch } from 'vue';
import { useRouter } from 'vue-router';
import { useUserStore } from '../stores/user';
import { fetchChatHistory, fetchUsers } from '../api/chat';
import { useTheme } from '../composables/useTheme';
import ChatMessage from '../components/ChatMessage.vue';
import UsersList from '../components/UsersList.vue';
import AppHeader from '../components/AppHeader.vue';
import ChatInput from '../components/ChatInput.vue';
import UiSidebar from '../components/ui-kit/UiSidebar.vue';
import UiOverlay from '../components/ui-kit/UiOverlay.vue';
import UiMenuLink from '../components/ui-kit/UiMenuLink.vue';
import UiButton from '../components/ui-kit/UiButton.vue';
import UiIconButton from '../components/ui-kit/UiIconButton.vue';
import UiUserItem from '../components/ui-kit/UiUserItem.vue';

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
	if (count === 1) return `${users[0]} is typing...`;
	if (count === 2) return `${users[0]} and ${users[1]} are typing...`;
	if (count === 3) return `${users[0]}, ${users[1]} and ${users[2]} are typing...`;
	if (count === 4) return `${users[0]}, ${users[1]}, ${users[2]} and ${users[3]} are typing...`;

	const remaining = count - 3;
	return `${users[0]}, ${users[1]}, ${users[2]} and ${remaining} others are typing...`;
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
            <UiOverlay :show="isSidebarOpen" @click="toggleSidebar" class="mobile-only" />
            <UiOverlay :show="isRightSidebarOpen" @click="toggleRightSidebar" class="mobile-only right-overlay" />

            <UiSidebar position="left" :isOpen="isSidebarOpen" class="app-sidebar">
                <div class="channels">
                    <UiMenuLink
                        v-for="chan in channels"
                        :key="chan"
                        :active="activeChannel === chan"
                        @click="selectChannel(chan)"
                        @dblclick="handleChannelDblClick(chan)"
                    >
                        <template #icon>
                            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                                <line x1="4" y1="9" x2="20" y2="9"></line>
                                <line x1="4" y1="15" x2="20" y2="15"></line>
                                <line x1="10" y1="3" x2="8" y2="21"></line>
                                <line x1="16" y1="3" x2="14" y2="21"></line>
                            </svg>
                        </template>
                        {{ chan }}
                    </UiMenuLink>
                </div>

                <template #footer>
                    <div class="sidebar-profile-wrap">
                        <UiUserItem
                            v-if="userStore.currentUser"
                            :username="userStore.currentUser.username"
                            :color="userStore.getUserColor(userStore.currentUser.username)"
                            :status="isConnected ? 'online' : 'offline'"
                            :subtitle="showStatusText ? (isConnected ? 'Connected' : 'Reconnecting...') : ''"
                            class="app-user-profile"
                        />
                    </div>
                </template>
            </UiSidebar>

            <div class="chat-container">
                <div class="chat-inner-header">
                    <div class="header-left">
                        <UiButton variant="ghost" @click="toggleSidebar" class="channel-toggle-btn">
                            <template #prepend>
                                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round" class="channel-hash">
                                    <line x1="4" y1="9" x2="20" y2="9"></line>
                                    <line x1="4" y1="15" x2="20" y2="15"></line>
                                    <line x1="10" y1="3" x2="8" y2="21"></line>
                                    <line x1="16" y1="3" x2="14" y2="21"></line>
                                </svg>
                            </template>
                            <span class="channel-name">{{ activeChannel }}</span>
                        </UiButton>
                    </div>

                    <div class="header-right">
                        <UiButton variant="ghost" @click="toggleRightSidebar" class="channel-toggle-btn">
                            <template #prepend>
                                <svg class="online-icon" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                                    <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
                                    <circle cx="9" cy="7" r="4"></circle>
                                    <path d="M23 21v-2a4 4 0 0 0-3-3.87"></path>
                                    <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
                                </svg>
                            </template>
                            Online: {{ allUsers.filter(u => u.online).length }}
                        </UiButton>
                    </div>
                </div>

                <div class="chat-area">
                    <div class="messages-container" ref="messagesContainer">
                        <div v-if="messages.length === 0" class="empty-state">No messages yet. Say hi!</div>

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

            <UiSidebar position="right" :isOpen="isRightSidebarOpen" class="app-right-sidebar">
                <div class="right-sidebar-header">
                    <span class="right-sidebar-title">Online</span>
                    <UiIconButton @click="toggleRightSidebar">
                        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                            <line x1="18" y1="6" x2="6" y2="18"></line>
                            <line x1="6" y1="6" x2="18" y2="18"></line>
                        </svg>
                    </UiIconButton>
                </div>
                <UsersList :users="allUsers" />
            </UiSidebar>
        </div>
    </div>
</template>

<style scoped>
.app-layout {
	display: flex;
	flex-direction: column;
	height: 100dvh;
	overflow: hidden;
}

/* Variables removed in favor of global themes.css */

.channel-toggle-btn {
    margin-left: -10px;
}

.channel-hash { color: var(--color-text-secondary); display: flex; align-items: center; justify-content: center; }
.channel-name { font-size: var(--font-size-xl); font-weight: var(--font-weight-semibold); letter-spacing: var(--letter-spacing-base); }

.main-content {
    display: flex;
    flex: 1;
    min-height: 0;
    position: relative;
}

.channels { padding: 10px 10px; flex: 1; overflow-y: auto; }

.sidebar-profile-wrap {
    padding: 12px 8px;
    padding-bottom: calc(8px + env(safe-area-inset-bottom));
}

.app-user-profile {
    background-color: var(--bg-active);
    border: 1px solid var(--border);
    border-radius: 14px;
    padding: 12px;
}

.mobile-only { display: none; }

.chat-container {
    flex: 1;
    display: flex;
    flex-direction: column;
    min-width: 0;
}
.chat-inner-header {
    height: 60px;
    padding: 0 24px;
    border-bottom: 1px solid var(--border);
    display: flex;
    align-items: center;
    justify-content: space-between;
    background-color: var(--color-bg);
    flex-shrink: 0;
}

.online-users-btn { margin-left: 0; }
.online-icon { color: var(--color-text-secondary); }
.online-users-btn:hover .online-icon { color: var(--color-text-primary); }
.online-text { font-size: var(--font-size-base); font-weight: var(--font-weight-semibold); }

.chat-area {
    flex: 1;
    display: flex;
    flex-direction: column;
    min-height: 0;
    position: relative;
    /* Account for safe areas (iOS/Android notches and home bars) */
    padding-bottom: env(safe-area-inset-bottom);
}

.messages-container {
    flex: 1;
    padding: 8px 12px;
    padding-bottom: 72px;
    overflow-y: auto;
    display: flex;
    flex-direction: column;
    scroll-behavior: smooth;
    background-color: var(--color-bg);
    transition: background-color var(--theme-transition-duration) ease;
}
.empty-state { text-align: center; color: var(--color-text-secondary); margin-top: auto; margin-bottom: auto; font-size: var(--font-size-md); }

::-webkit-scrollbar { width: 6px; height: 6px; }
::-webkit-scrollbar-track { background: transparent; }
::-webkit-scrollbar-thumb { background: var(--color-border-hover); border-radius: 3px; }
::-webkit-scrollbar-thumb:hover { background: var(--color-text-secondary); }

.right-sidebar-header {
    display: none;
    padding: 16px;
    border-bottom: 1px solid var(--color-border);
    align-items: center;
    justify-content: space-between;
}
.right-sidebar-title { font-weight: var(--font-weight-semibold); font-size: var(--font-size-lg); color: var(--color-text-primary); }

@media (max-width: 768px) {
    .chat-inner-header { padding: 0 16px; }

    .app-sidebar, .app-right-sidebar {
        position: fixed;
        top: 48px; bottom: 0;
        z-index: 100;
        margin: 0 !important;
        box-shadow: 5px 0 15px rgba(0,0,0,0.5);
    }

    .app-sidebar {
        left: 0;
        transform: translateX(-100%);
    }

    .app-right-sidebar {
        right: 0;
        transform: translateX(100%);
        box-shadow: -5px 0 15px rgba(0,0,0,0.5);
    }

    .right-sidebar-header { display: flex; }

    .app-layout.light-theme .app-sidebar, .app-layout.light-theme .app-right-sidebar {
        box-shadow: 5px 0 15px rgba(0,0,0,0.1);
    }
    .app-layout.light-theme .app-right-sidebar {
        box-shadow: -5px 0 15px rgba(0,0,0,0.1);
    }

    .app-sidebar.is-open { transform: translateX(0); }
    .app-right-sidebar.is-open { transform: translateX(0); }

    .mobile-only {
        display: block;
        top: 48px;
    }

    .mobile-only.right-overlay:not(.is-visible) { display: none; }

    .messages-container { padding: 16px;padding-bottom: 64px; }
}
</style>
