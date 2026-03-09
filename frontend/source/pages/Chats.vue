<template>
  <div class="app-container">
    <aside class="sidebar" :class="{ 'mobile-hidden': activeChatId }">
      <div class="sidebar-search">
        <input type="text" placeholder="Поиск" class="neon-input">
      </div>

      <div class="chat-list">
        <div v-for="chat in chats" :key="chat.id" class="chat-item" 
             :class="{ active: activeChatId === chat.id }" @click="activeChatId = chat.id">
          <div class="avatar"></div>
          <div class="chat-info">
            <div class="chat-name-row">
              <span class="name">{{ chat.name }}</span>
              <span class="time">15:04</span>
            </div>
            <p class="last-msg">Нажми, чтобы написать...</p>
          </div>
        </div>
      </div>

      <div class="add-container">
        <transition name="pop">
          <div v-if="showAddMenu" class="add-menu">
            <button @click="openModal('contact')">👤 Добавить контакт</button>
            <button @click="openModal('group')">👥 Создать группу</button>
          </div>
        </transition>
        <button class="fab-add" @click="showAddMenu = !showAddMenu" :class="{ active: showAddMenu }">+</button>
      </div>

      <div class="bottom-profile" @click="showSettings = true">
        <div class="user-wrap">
          <div class="avatar online"></div>
          <div class="user-id">
            <span class="nick">{{ user.name }}</span>
            <span class="hash">#0001</span>
          </div>
        </div>
        <div class="gear">⚙️</div>
      </div>
    </aside>

    <main class="chat-main" :class="{ 'mobile-visible': activeChatId }">
      <template v-if="activeChatId">
        <header class="chat-header">
          <button class="back-btn" @click="activeChatId = null">←</button>
          <div class="chat-title-wrap">
            <span class="chat-title">{{ activeChatName }}</span>
            <span class="status">online</span>
          </div>
        </header>

        <div class="messages-scroll" ref="messageBox">
          <div v-for="(msg, i) in messages" :key="i" class="msg-row" :class="{ 'is-own': msg.own }">
            <div class="bubble">
              {{ msg.text }}
              <span class="msg-time">{{ msg.time }}</span>
            </div>
          </div>
        </div>

        <footer class="input-footer">
          <div class="input-group">
            <input v-model="currentText" @keyup.enter="sendMsg" placeholder="Написать сообщение..." />
            <button class="send-btn" @click="sendMsg">▲</button>
          </div>
        </footer>
      </template>
      <div v-else class="empty-screen">Выберите чат, чтобы начать общение</div>
    </main>

    <transition name="fade">
      <div v-if="showSettings" class="modal-overlay" @click.self="showSettings = false">
        <div class="modal-content">
          <div class="modal-header">
            <h3>Мой профиль</h3>
            <button @click="showSettings = false">✕</button>
          </div>
          <div class="modal-body">
            <div class="profile-pic-big"></div>
            <div class="form-group">
              <label>Имя пользователя</label>
              <input v-model="user.name" type="text" />
            </div>
            <div class="form-group">
              <label>Email / Ник</label>
              <input v-model="user.email" type="text" />
            </div>
            <button class="action-btn" @click="showSettings = false">Сохранить</button>
          </div>
        </div>
      </div>
    </transition>

    <transition name="fade">
      <div v-if="activeModal" class="modal-overlay" @click.self="activeModal = null">
        <div class="modal-content">
          <div class="modal-header">
            <h3>{{ activeModal === 'contact' ? 'Добавить контакт' : 'Новая группа' }}</h3>
            <button @click="activeModal = null">✕</button>
          </div>
          <div class="modal-body">
            <input type="text" :placeholder="activeModal === 'contact' ? 'Введите email или ник' : 'Название группы'" class="modal-input" />
            <button class="action-btn" @click="activeModal = null">Подтвердить</button>
          </div>
        </div>
      </div>
    </transition>
  </div>
</template>

<script setup>
import { ref, computed, nextTick } from 'vue';

const activeChatId = ref(null);
const currentText = ref('');
const showSettings = ref(false);
const showAddMenu = ref(false);
const activeModal = ref(null);
const messageBox = ref(null);

const user = ref({ name: 'DrBopoH', email: '@voron' });
const chats = ref([{ id: 1, name: 'Main Dev Group' }, { id: 2, name: 'Project Manager' }]);
const messages = ref([{ text: 'Всё вернул как было, только функции добавил!', own: false, time: '15:10' }]);

const activeChatName = computed(() => chats.value.find(c => c.id === activeChatId.value)?.name || '');

const sendMsg = async () => {
  if (!currentText.value.trim()) return;
  messages.value.push({
    text: currentText.value,
    own: true,
    time: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
  });
  currentText.value = '';
  await nextTick();
  messageBox.value.scrollTop = messageBox.value.scrollHeight;
};

const openModal = (type) => {
  activeModal.value = type;
  showAddMenu.value = false;
};
</script>

<style scoped>
/* ФУЛЛСКРИН БАЗА */
.app-container { display: flex; width: 100vw; height: 100vh; background: #09090b; color: #fff; overflow: hidden; }

/* САЙДБАР */
.sidebar { width: 380px; background: #111113; border-right: 1px solid #27272a; display: flex; flex-direction: column; position: relative; }
.sidebar-search { padding: 15px; }
.neon-input { width: 100%; background: #1c1c1f; border: 1px solid #333; padding: 10px; border-radius: 8px; color: #fff; outline: none; }
.neon-input:focus { border-color: #22c55e; }

.chat-list { flex: 1; overflow-y: auto; }
.chat-item { display: flex; padding: 12px 15px; gap: 12px; cursor: pointer; }
.chat-item:hover { background: #1a1a1c; }
.chat-item.active { background: #22c55e; }

.avatar { width: 48px; height: 48px; background: #27272a; border-radius: 50%; }
.online { position: relative; }
.online::after { content: ''; position: absolute; bottom: 2px; right: 2px; width: 12px; height: 12px; background: #22c55e; border-radius: 50%; border: 2px solid #111; }

/* ПЛЮСИК И МЕНЮ */
.add-container { position: absolute; bottom: 85px; right: 20px; display: flex; flex-direction: column; align-items: flex-end; gap: 10px; }
.add-menu { background: #1c1c1f; border: 1px solid #333; border-radius: 12px; padding: 8px; box-shadow: 0 10px 30px #000; }
.add-menu button { background: none; border: none; color: #fff; padding: 10px; text-align: left; cursor: pointer; width: 100%; border-radius: 6px; }
.add-menu button:hover { background: #27272a; color: #22c55e; }
.fab-add { width: 56px; height: 56px; background: #22c55e; border: none; border-radius: 50%; font-size: 32px; cursor: pointer; transition: 0.3s; }
.fab-add.active { transform: rotate(45deg); background: #ef4444; }

/* ПРОФИЛЬ */
.bottom-profile { padding: 10px 15px; background: #09090b; border-top: 1px solid #27272a; display: flex; justify-content: space-between; align-items: center; cursor: pointer; }
.user-id { display: flex; flex-direction: column; gap: 2px; margin-left: 10px; }
.hash { font-size: 11px; color: #71717a; }

/* ЧАТ */
.chat-main { flex: 1; display: flex; flex-direction: column; background: #0a0a0a; }
.chat-header { padding: 15px 25px; background: #111113; border-bottom: 1px solid #27272a; display: flex; gap: 15px; align-items: center; }
.status { color: #22c55e; font-size: 12px; }

.messages-scroll { flex: 1; padding: 20px; overflow-y: auto; display: flex; flex-direction: column; gap: 10px; }
.msg-row { display: flex; width: 100%; }
.is-own { justify-content: flex-end; }
.bubble { max-width: 65%; padding: 10px 15px; border-radius: 15px; background: #27272a; position: relative; }
.is-own .bubble { background: #166534; border: 1px solid #22c55e; }
.msg-time { font-size: 10px; opacity: 0.5; margin-left: 8px; }

.input-footer { padding: 20px; }
.input-group { display: flex; background: #1c1c1f; border: 1px solid #333; border-radius: 25px; padding: 5px 15px; align-items: center; }
.input-group input { flex: 1; background: none; border: none; color: #fff; padding: 10px; outline: none; }
.send-btn { width: 40px; height: 40px; background: #22c55e; border: none; border-radius: 50%; cursor: pointer; }

/* МОДАЛКИ */
.modal-overlay { position: fixed; inset: 0; background: rgba(0,0,0,0.8); display: flex; align-items: center; justify-content: center; z-index: 1000; }
.modal-content { background: #121214; border: 1px solid #333; width: 360px; border-radius: 20px; padding: 25px; }
.modal-header { display: flex; justify-content: space-between; margin-bottom: 20px; }
.profile-pic-big { width: 80px; height: 80px; background: #27272a; border-radius: 50%; border: 2px solid #22c55e; margin: 0 auto 15px; }
.form-group { margin-bottom: 15px; }
.form-group label { font-size: 12px; color: #71717a; margin-bottom: 5px; display: block; }
.form-group input, .modal-input { width: 100%; background: #1c1c1f; border: 1px solid #333; padding: 10px; border-radius: 8px; color: #fff; }
.action-btn { width: 100%; background: #22c55e; border: none; padding: 12px; border-radius: 8px; font-weight: bold; cursor: pointer; margin-top: 10px; }

/* АНИМАЦИИ */
.pop-enter-active { transition: all 0.2s ease-out; }
.pop-enter-from { opacity: 0; transform: scale(0.9) translateY(10px); }
.fade-enter-active { transition: opacity 0.3s; }
.fade-enter-from { opacity: 0; }

@media (max-width: 768px) {
  .sidebar.mobile-hidden { display: none; }
  .chat-main { display: none; }
  .chat-main.mobile-visible { display: flex; position: fixed; inset: 0; }
  .back-btn { display: block; background: none; border: none; color: #22c55e; font-size: 24px; }
}
@media (min-width: 769px) { .back-btn { display: none; } }
</style>