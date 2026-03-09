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

    ws.value.onopen = () => {
        console.log('🔗 WebSocket подключен!');
        isConnected.value = true;
    };

    ws.value.onmessage = (event) => {
        try {
            const data = JSON.parse(event.data);
            messages.value.push({
                id: data.id,
                user: data.username,
                text: data.text,
                time: formatTime(data.timestamp)
            });
            scrollToBottom();
        } catch (e) {
            console.error('Ошибка парсинга сообщения:', e);
        }
    };

    ws.value.onclose = () => {
        console.log('❌ WebSocket отключен. Пробуем переподключиться через 3 сек...');
        isConnected.value = false;
        setTimeout(connectWebSocket, 3000);
    };
};

onMounted(async () => {
    if (!userStore.currentUser) {
        router.push('/auth');
        return;
    }
    
    await loadHistory();
    connectWebSocket();
});

onUnmounted(() => {
    if (ws.value) {
        ws.value.close();
    }
});

const sendMessage = () => {
    if (!newMessage.value.trim() || !isConnected.value) return;
    
    const payload = {
        text: newMessage.value,
    };

    ws.value?.send(JSON.stringify(payload));
    
    newMessage.value = '';
};

const logout = () => {
    userStore.setUser(null);
    if (ws.value) ws.value.close();
    router.push('/');
};
</script>

<template>
    <div class="chat-layout" v-if="userStore.currentUser">
        
        <header class="chat-header">
            <div class="chat-title">
                <img src="/Logo.svg" alt="Leaves" class="header-logo" />
                <h2># global</h2>
                <div :class="['status-dot', { online: isConnected }]"></div>
            </div>
            
            <div class="user-control">
                <span class="username">{{ userStore.currentUser.username }}</span>
                <div class="avatar">{{ userStore.currentUser.username.charAt(0).toUpperCase() }}</div>
                <button @click="logout" class="logout-btn">
                    <svg viewBox="0 0 24 24" width="18" height="18" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"></path><polyline points="16 17 21 12 16 7"></polyline><line x1="21" y1="12" x2="9" y2="12"></line></svg>
                </button>
            </div>
        </header>

        <main class="chat-messages" ref="messagesContainer">
            <div v-if="messages.length === 0" class="empty-state">
                <div class="welcome-orb"></div>
                <h3>Welcome to #global</h3>
                <p>This is the start of the conversation.</p>
            </div>

            <div class="message-wrapper" v-for="msg in messages" :key="msg.id">
                <div class="msg-avatar">{{ msg.user ? msg.user.charAt(0).toUpperCase() : '?' }}</div>
                <div class="msg-content">
                    <div class="msg-header">
                        <span class="msg-username">{{ msg.user }}</span>
                        <span class="msg-time">{{ msg.time }}</span>
                    </div>
                    <p class="msg-text">{{ msg.text }}</p>
                </div>
            </div>
        </main>

        <footer class="chat-input-area">
            <div class="input-wrapper">
                <input 
                    type="text" 
                    v-model="newMessage" 
                    @keyup.enter="sendMessage"
                    :placeholder="isConnected ? 'Message #global' : 'Connecting to server...'" 
                    :disabled="!isConnected"
                />
                <button class="send-btn" @click="sendMessage" :disabled="!isConnected || !newMessage.trim()">
                    <svg viewBox="0 0 24 24" width="20" height="20" fill="currentColor"><path d="M2.01 21L23 12 2.01 3 2 10l15 2-15 2z"></path></svg>
                </button>
            </div>
        </footer>

    </div>
</template>

<style scoped>
.chat-layout {
    display: flex;
    flex-direction: column;
    height: 100vh;
    width: 100%; /* Убрали 100vw, чтобы не было горизонтального скролла */
    background-color: #050807; 
    color: #c8c2b8;
    overflow: hidden; /* Защита от любых выползаний */
}

/* Header */
.chat-header {
    flex-shrink: 0; /* Не дает шапке сжиматься */
    height: 60px;
    padding: 0 24px;
    background-color: #080b0a;
    border-bottom: 1px solid #0f1714;
    display: flex;
    justify-content: space-between;
    align-items: center;
    z-index: 10;
}

.chat-title { display: flex; align-items: center; gap: 12px; }
.header-logo { width: 28px; height: 28px; }
.chat-title h2 { font-size: 18px; font-weight: 600; color: #c8c2b8; margin: 0; }

.status-dot { width: 8px; height: 8px; border-radius: 50%; background-color: #ff5555; transition: background-color 0.3s; }
.status-dot.online { background-color: #5fca08; box-shadow: 0 0 8px rgba(95, 202, 8, 0.5); }

.user-control { display: flex; align-items: center; gap: 16px; }
.username { font-size: 14px; font-weight: 500; color: #c8c2b8; }

.avatar { width: 36px; height: 36px; background-color: #0f1714; color: #5fca08; border: 2px solid #080b0a; border-radius: 50%; display: flex; justify-content: center; align-items: center; font-weight: bold; font-size: 14px; }

.logout-btn { background: transparent; border: none; color: #8a867f; cursor: pointer; padding: 8px; border-radius: 6px; transition: all 0.2s; display: flex; align-items: center; justify-content: center; }
.logout-btn:hover { color: #ff5555; background: #0f1714; }

/* Messages Area */
.chat-messages {
    flex: 1; /* Занимает все свободное пространство */
    overflow-y: auto;
    padding: 24px 20px;
    display: flex;
    flex-direction: column;
    gap: 20px;
    /* Красивый скроллбар для Webkit */
    scrollbar-width: thin;
    scrollbar-color: #0f1714 transparent;
}
.chat-messages::-webkit-scrollbar { width: 6px; }
.chat-messages::-webkit-scrollbar-track { background: transparent; }
.chat-messages::-webkit-scrollbar-thumb { background-color: #0f1714; border-radius: 10px; }

.empty-state { display: flex; flex-direction: column; justify-content: flex-end; height: 100%; padding-bottom: 40px; }
.welcome-orb { width: 48px; height: 48px; background-color: #5fca08; border-radius: 50%; margin-bottom: 16px; filter: blur(10px); opacity: 0.5; }
.empty-state h3 { font-size: 24px; color: #c8c2b8; margin: 0 0 8px 0; }
.empty-state p { color: #8a867f; margin: 0; }

.message-wrapper { display: flex; gap: 16px; padding: 8px 16px; border-radius: 12px; transition: background-color 0.2s; }
.message-wrapper:hover { background-color: #080b0a; }

.msg-avatar { width: 42px; height: 42px; border-radius: 50%; background-color: #0f1714; border: 1px solid #5fca08; display: flex; justify-content: center; align-items: center; font-weight: bold; font-size: 18px; color: #5fca08; flex-shrink: 0; }

.msg-header { display: flex; align-items: baseline; gap: 10px; margin-bottom: 6px; }
.msg-username { font-weight: 600; color: #c8c2b8; font-size: 15px; }
.msg-time { font-size: 12px; color: #64615c; }
.msg-text { font-size: 15px; line-height: 1.5; color: #8a867f; margin: 0; word-break: break-word; }

/* Input Area */
.chat-input-area {
    flex-shrink: 0; /* Чтобы инпут не ужимался */
    padding: 0 24px 30px 24px;
    background-color: #050807; /* Чтобы сообщения уходили под него плавно */
}

.input-wrapper { background-color: #080b0a; border: 1px solid #0f1714; border-radius: 16px; display: flex; align-items: center; padding: 12px 20px; transition: border-color 0.2s ease; }
.input-wrapper:focus-within { border-color: rgba(95, 202, 8, 0.5); }

.input-wrapper input { flex: 1; background: transparent; border: none; outline: none; color: #c8c2b8; font-size: 15px; }
.input-wrapper input::placeholder { color: #64615c; }
.input-wrapper input:disabled { opacity: 0.5; cursor: not-allowed; }

.send-btn { background: transparent; border: none; color: #5fca08; cursor: pointer; display: flex; align-items: center; justify-content: center; padding: 4px; transition: transform 0.2s, color 0.2s; }
.send-btn:hover:not(:disabled) { transform: scale(1.1); color: #4da806; }
.send-btn:disabled { color: #64615c; cursor: not-allowed; }
</style>