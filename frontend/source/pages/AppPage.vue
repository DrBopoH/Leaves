<script setup lang="ts"> // source/pages/AppPage.vue
import { ref, onMounted, onUnmounted, nextTick } from 'vue';
import { useRouter } from 'vue-router';
import { useUserStore } from '../stores/user';
import { fetchChatHistory } from '../api/auth';

const router = useRouter();
const userStore = useUserStore();

const newMessage = ref('');
const messagesContainer = ref<HTMLElement | null>(null);
const ws = ref<WebSocket | null>(null);
const isConnected = ref(false);

const messages = ref<any[]>([]);

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
    <div class="app-layout">

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
                    <div class="input-wrapper">
                        <textarea
                            v-model="newMessage"
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
.app-layout {
    display: flex;
    flex-direction: column;
    height: 100dvh;
    background-color: #050807;
    color: #c8c2b8;
    overflow: hidden;
}

/* Общий Header */
.chat-header {
    height: 70px;
    padding: 0 24px;
    border-bottom: 1px solid #0f1714;
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: #080b0a;
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
    transition: background-color 0.2s;
    color: #c8c2b8;
    margin-left: -10px;
}
.channel-toggle-btn:hover {
    background-color: #0f1714;
}
.channel-hash {
    color: #64615c;
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
.brand-text { font-size: 20px; font-weight: 700; color: #5fca08; }

.header-status { gap: 8px; font-size: 13px; color: #64615c; }
.status-dot { width: 8px; height: 8px; border-radius: 50%; background-color: #dc3545; }
.status-dot.online { background-color: #5fca08; box-shadow: 0 0 8px rgba(95, 202, 8, 0.4); }

.main-content {
    display: flex;
    flex: 1;
    min-height: 0;
    position: relative;
}

.sidebar {
    width: 260px;
    background-color: #080b0a;
    border-right: 1px solid #0f1714;
    display: flex;
    flex-direction: column;
    flex-shrink: 0;
    transition: margin-left 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}
.sidebar:not(.is-open) {
    margin-left: -261px;
}
.channels { padding: 15px 10px; flex: 1; overflow-y: auto; }
.channel { padding: 10px 15px; border-radius: 8px; color: #8a867f; cursor: pointer; transition: all 0.2s; font-weight: 500; }
.channel:hover { background-color: #0f1714; color: #c8c2b8; }
.channel.active { background-color: #1a231e; color: #5fca08; }

.sidebar-overlay { display: none; }

.chat-area { flex: 1; display: flex; flex-direction: column; background-color: #050807; min-width: 0; }
.messages-container { flex: 1; padding: 24px; overflow-y: auto; display: flex; flex-direction: column; scroll-behavior: smooth; }
.empty-state { text-align: center; color: #64615c; margin-top: auto; margin-bottom: auto; font-size: 15px; }

.message { display: flex; gap: 16px; align-items: flex-start; max-width: 100%; transition: margin 0.2s ease; }
.msg-avatar {
    width: 40px; height: 40px; border-radius: 12px;
    background-color: #080b0a; border: 1px solid;
    display: flex; justify-content: center; align-items: center;
    font-weight: bold; font-size: 18px; flex-shrink: 0;
}
.msg-content { max-width: calc(100% - 56px); }
.msg-header { display: flex; align-items: baseline; gap: 10px; margin-bottom: 4px; }
.msg-username { font-weight: 600; font-size: 15px; }
.msg-time { font-size: 12px; color: #64615c; }

.msg-bubble { background-color: transparent; border: 1px solid transparent; padding: 2px 0; border-radius: 8px; }
.msg-bubble.is-mine {
    background-color: #1a231e;
    border: 1px solid #233129;
    padding: 10px 16px;
    border-radius: 16px;
    border-top-left-radius: 4px;
}
.msg-text { font-size: 15px; line-height: 1.5; color: #c8c2b8; margin: 0; word-break: break-word; }

:deep(.msg-image) { max-width: 100%; max-height: 250px; border-radius: 8px; margin-top: 8px; border: 1px solid #0f1714; display: block; }
:deep(.msg-link) { color: #5fca08; text-decoration: underline; text-underline-offset: 2px; }
:deep(.msg-link:hover) { color: #88ffb4; }

.chat-input-area {
    flex-shrink: 0;
    padding: 0 24px 24px 24px;
    padding-bottom: calc(24px + env(safe-area-inset-bottom));
    background-color: #050807;
}
.input-wrapper { background-color: #080b0a; border: 1px solid #0f1714; border-radius: 16px; display: flex; align-items: flex-end; padding: 12px 20px; transition: border-color 0.2s ease; }
.input-wrapper:focus-within { border-color: rgba(95, 202, 8, 0.4); }
.chat-input { flex: 1; background: transparent; border: none; color: #c8c2b8; font-size: 15px; outline: none; resize: none; max-height: 120px; font-family: inherit; line-height: 1.5; padding-top: 2px; }
.chat-input::placeholder { color: #64615c; }
.send-btn { background: #5fca08; color: #050807; border: none; width: 36px; height: 36px; border-radius: 10px; display: flex; justify-content: center; align-items: center; cursor: pointer; transition: all 0.2s; flex-shrink: 0; margin-left: 12px; }
.send-btn:hover:not(:disabled) { background: #4da806; transform: scale(1.05); }
.send-btn:disabled { background: #1a231e; color: #64615c; cursor: not-allowed; }

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
    .sidebar-overlay.is-open {
        opacity: 1;
        pointer-events: auto;
    }

    .messages-container { padding: 16px; }
    .chat-input-area { padding: 0 16px 16px 16px; padding-bottom: calc(16px + env(safe-area-inset-bottom)); }
}
</style>
