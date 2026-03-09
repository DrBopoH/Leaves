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

const getUserColor = (username: string) => {
    if (!username) return '#5fca08';
    let hash = 0;
    for (let i = 0; i < username.length; i++) {
        hash = username.charCodeAt(i) + ((hash << 5) - hash);
    }
    const hue = Math.abs(hash % 360);
    return `hsl(${hue}, 80%, 65%)`;
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
        <div class="sidebar">
            <div class="brand">
                <img src="/Logo.svg" alt="Leaves" class="brand-logo" />
                <span class="brand-text">Leaves</span>
            </div>
            <div class="channels">
                <div class="channel active"># general</div>
            </div>
        </div>

        <div class="chat-area">
            <div class="chat-header">
                <h2># general</h2>
                <div class="header-status">
                    <span class="status-dot" :class="{ online: isConnected }"></span>
                    {{ isConnected ? 'Connected' : 'Reconnecting...' }}
                </div>
            </div>

            <div class="messages-container" ref="messagesContainer">
                <div v-if="messages.length === 0" class="empty-state">No messages yet. Say hi!</div>
                
                <div v-for="(msg, index) in messages" :key="msg.id" 
                     :class="['message', { 'my-message': msg.user === userStore.currentUser?.username }]"
                     :style="{ marginBottom: (index < messages.length - 1 && messages[index + 1].user === msg.user) ? '4px' : '24px' }">
                    
                    <div class="msg-avatar" 
                         v-if="msg.user !== userStore.currentUser?.username"
                         :style="{ 
                             borderColor: getUserColor(msg.user), 
                             color: getUserColor(msg.user),
                             visibility: (index > 0 && messages[index - 1].user === msg.user) ? 'hidden' : 'visible' 
                         }">
                        {{ msg.user.charAt(0).toUpperCase() }}
                    </div>

                    <div class="msg-content">
                        <div class="msg-header" 
                             v-if="msg.user !== userStore.currentUser?.username && !(index > 0 && messages[index - 1].user === msg.user)">
                            <span class="msg-username" :style="{ color: getUserColor(msg.user) }">{{ msg.user }}</span>
                            <span class="msg-time">{{ msg.time }}</span>
                        </div>
                        
                        <div class="msg-bubble">
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
                    ></textarea>
                    <button @click="sendMessage" class="send-btn" :disabled="!newMessage.trim() || !isConnected">
                        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="22" y1="2" x2="11" y2="13"></line><polygon points="22 2 15 22 11 13 2 9 22 2"></polygon></svg>
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.app-layout { display: flex; height: 100vh; background-color: #050807; color: #c8c2b8; overflow: hidden; }
.sidebar { width: 260px; background-color: #080b0a; border-right: 1px solid #0f1714; display: flex; flex-direction: column; }
.brand { padding: 20px; display: flex; align-items: center; gap: 12px; border-bottom: 1px solid #0f1714; }
.brand-logo { width: 32px; height: 32px; }
.brand-text { font-size: 20px; font-weight: 700; color: #5fca08; }
.channels { padding: 15px 10px; flex: 1; overflow-y: auto; }
.channel { padding: 10px 15px; border-radius: 8px; color: #8a867f; cursor: pointer; transition: all 0.2s; font-weight: 500; }
.channel:hover { background-color: #0f1714; color: #c8c2b8; }
.channel.active { background-color: #1a231e; color: #5fca08; }

.chat-area { flex: 1; display: flex; flex-direction: column; background-color: #050807; min-width: 0; }
.chat-header { padding: 20px 24px; border-bottom: 1px solid #0f1714; display: flex; justify-content: space-between; align-items: center; background-color: #080b0a; z-index: 10; }
.chat-header h2 { margin: 0; font-size: 18px; color: #c8c2b8; font-weight: 600; }
.header-status { display: flex; align-items: center; gap: 8px; font-size: 13px; color: #64615c; }
.status-dot { width: 8px; height: 8px; border-radius: 50%; background-color: #dc3545; }
.status-dot.online { background-color: #5fca08; box-shadow: 0 0 8px rgba(95, 202, 8, 0.4); }

.messages-container { flex: 1; padding: 24px; overflow-y: auto; display: flex; flex-direction: column; scroll-behavior: smooth; }
.empty-state { text-align: center; color: #64615c; margin-top: auto; margin-bottom: auto; font-size: 15px; }

/* Сообщения */
.message { display: flex; gap: 16px; align-items: flex-start; max-width: 85%; transition: margin 0.2s ease; }
.my-message { align-self: flex-end; flex-direction: row-reverse; }

.msg-avatar { 
    width: 40px; height: 40px; border-radius: 12px; 
    background-color: #080b0a; border: 1px solid;
    display: flex; justify-content: center; align-items: center; 
    font-weight: bold; font-size: 18px; flex-shrink: 0; 
}

.msg-header { display: flex; align-items: baseline; gap: 10px; margin-bottom: 6px; }
.msg-username { font-weight: 600; font-size: 15px; }
.msg-time { font-size: 12px; color: #64615c; }
.msg-bubble { background-color: transparent; }
.my-message .msg-bubble { background-color: #1a231e; border: 1px solid #233129; padding: 10px 16px; border-radius: 16px; border-bottom-right-radius: 4px; }
.msg-text { font-size: 15px; line-height: 1.5; color: #c8c2b8; margin: 0; word-break: break-word; }

:deep(.msg-image) {
    max-width: 100%;
    max-height: 250px;
    border-radius: 8px;
    margin-top: 8px;
    border: 1px solid #0f1714;
    display: block;
}
:deep(.msg-link) {
    color: #5fca08;
    text-decoration: underline;
    text-underline-offset: 2px;
}
:deep(.msg-link:hover) {
    color: #88ffb4;
}

.chat-input-area { flex-shrink: 0; padding: 0 24px 30px 24px; background-color: #050807; }
.input-wrapper { background-color: #080b0a; border: 1px solid #0f1714; border-radius: 16px; display: flex; align-items: flex-end; padding: 12px 20px; transition: border-color 0.2s ease; }
.input-wrapper:focus-within { border-color: rgba(95, 202, 8, 0.4); }

.chat-input { 
    flex: 1; background: transparent; border: none; color: #c8c2b8; 
    font-size: 15px; outline: none; resize: none; max-height: 120px; 
    font-family: inherit; line-height: 1.5; padding-top: 2px;
}
.chat-input::placeholder { color: #64615c; }
.send-btn { background: #5fca08; color: #050807; border: none; width: 36px; height: 36px; border-radius: 10px; display: flex; justify-content: center; align-items: center; cursor: pointer; transition: all 0.2s; flex-shrink: 0; margin-left: 12px; }
.send-btn:hover:not(:disabled) { background: #4da806; transform: scale(1.05); }
.send-btn:disabled { background: #1a231e; color: #64615c; cursor: not-allowed; }
</style>