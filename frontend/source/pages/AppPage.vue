<script setup lang="ts"> // source/pages/AppPage.vue
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useUserStore } from '../stores/user';

const router = useRouter();
const userStore = useUserStore();
const newMessage = ref('');

const messages = ref([
    { id: 1, user: 'terminal_ninja', text: 'Всем привет! Тестируем новый чат 🚀', time: '14:23' },
    { id: 2, user: 'system', text: 'Добро пожаловать в глобальный канал.', time: '14:24' },
]);

onMounted(() => {
    // Если юзер не в Pinia, отправляем логиниться
    if (!userStore.currentUser) {
        router.push('/signin');
    }
});

const sendMessage = () => {
    if (!newMessage.value.trim() || !userStore.currentUser) return;
    
    messages.value.push({
        id: Date.now(),
        user: userStore.currentUser.username,
        text: newMessage.value,
        time: new Date().toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' })
    });
    
    newMessage.value = '';
};

const logout = () => {
    userStore.setUser(null);
    router.push('/');
};
</script>

<template>
    <div class="chat-layout" v-if="userStore.currentUser">
        
        <header class="chat-header">
            <div class="chat-title">
                <img src="/Logo.svg" alt="Leaves" class="header-logo" />
                <h2># global</h2>
            </div>
            
            <div class="user-control">
                <span class="username">{{ userStore.currentUser.username }}</span>
                <div class="avatar">{{ userStore.currentUser.username.charAt(0).toUpperCase() }}</div>
                <button @click="logout" class="logout-btn">
                    <svg viewBox="0 0 24 24" width="18" height="18" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"></path><polyline points="16 17 21 12 16 7"></polyline><line x1="21" y1="12" x2="9" y2="12"></line></svg>
                </button>
            </div>
        </header>

        <main class="chat-messages">
            <div class="message-wrapper" v-for="msg in messages" :key="msg.id">
                <div class="msg-avatar">{{ msg.user.charAt(0).toUpperCase() }}</div>
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
                    placeholder="Message #global" 
                />
                <button class="send-btn" @click="sendMessage">
                    <svg viewBox="0 0 24 24" width="20" height="20" fill="currentColor"><path d="M2.01 21L23 12 2.01 3 2 10l15 2-15 2z"></path></svg>
                </button>
            </div>
        </footer>

    </div>
</template>

<style scoped>
.chat-layout {
    display: flex; flex-direction: column; height: 100vh; width: 100vw;
    background-color: #313338; color: #dbdee1;
}

/* Header */
.chat-header {
    height: 48px; padding: 0 16px; background-color: #313338;
    border-bottom: 1px solid #1e1f22; box-shadow: 0 1px 2px rgba(0,0,0,0.2);
    display: flex; justify-content: space-between; align-items: center; z-index: 10;
}
.chat-title { display: flex; align-items: center; gap: 10px; }
.header-logo { width: 24px; height: 24px; }
.chat-title h2 { font-size: 16px; font-weight: 600; color: #fff; margin: 0; }

.user-control { display: flex; align-items: center; gap: 12px; }
.username { font-size: 14px; font-weight: 500; }
.avatar {
    width: 32px; height: 32px; background-color: #88ffb4; color: #000;
    border-radius: 50%; display: flex; justify-content: center; align-items: center;
    font-weight: bold; font-size: 14px;
}
.logout-btn {
    background: transparent; border: none; color: #b5bac1; cursor: pointer;
    padding: 6px; border-radius: 4px; transition: all 0.2s;
}
.logout-btn:hover { color: #f23f42; background: #2b2d31; }

/* Messages */
.chat-messages {
    flex: 1; overflow-y: auto; padding: 20px 10px;
    display: flex; flex-direction: column; gap: 16px;
}
.message-wrapper { display: flex; gap: 16px; padding: 4px 20px; }
.message-wrapper:hover { background-color: #2e3035; }
.msg-avatar {
    width: 40px; height: 40px; border-radius: 50%; background-color: #5865F2;
    display: flex; justify-content: center; align-items: center;
    font-weight: bold; font-size: 18px; color: #fff; flex-shrink: 0;
}
.msg-header { display: flex; align-items: baseline; gap: 8px; margin-bottom: 4px; }
.msg-username { font-weight: 500; color: #fff; font-size: 15px; }
.msg-time { font-size: 12px; color: #949ba4; }
.msg-text { font-size: 15px; line-height: 1.4; color: #dbdee1; margin: 0; }

/* Input */
.chat-input-area { padding: 0 16px 24px 16px; }
.input-wrapper { background-color: #383a40; border-radius: 8px; display: flex; align-items: center; padding: 10px 16px; }
.input-wrapper input { flex: 1; background: transparent; border: none; outline: none; color: #dbdee1; font-size: 15px; }
.input-wrapper input::placeholder { color: #80848e; }
.send-btn { background: transparent; border: none; color: #88ffb4; cursor: pointer; display: flex; align-items: center; transition: transform 0.2s; }
.send-btn:hover { transform: scale(1.1); color: #6ce09b; }
</style>