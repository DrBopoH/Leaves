<script setup lang="ts">
	// Copyright (C) 2026 MorangTong Creative Studio
	// SPDX-License-Identifier: AGPL-3.0-or-later

	// source/pages/debug/AdminPage.vue
	import { ref } from 'vue';
	import { useRouter } from 'vue-router';

	const router = useRouter();

	const isDarkTheme = ref(true);
	const toggleTheme = () => { isDarkTheme.value = !isDarkTheme.value; };

	const isAuthenticated = ref(false);
	const adminKeyInput = ref('');
	const authError = ref('');

	const SECRET_ADMIN_KEY = '777'; 

	const handleLogin = () => {
		if (adminKeyInput.value === SECRET_ADMIN_KEY) {
			isAuthenticated.value = true;
			authError.value = '';
			loadAdminData();
		} else {
			authError.value = 'Неверный ключ доступа';
			adminKeyInput.value = '';
		}
	};

	const totalRegistered = ref(0);
	const onlineCount = ref(0);
	const totalMessages = ref(0);

	interface User {
		id: string;
		username: string;
		isOnline: boolean;
		isBanned: boolean;
	}

	const users = ref<User[]>([]);

	const loadAdminData = async () => {
		totalRegistered.value = 142;
		onlineCount.value = 12;
		totalMessages.value = 15843;
		
		users.value = [
			{ id: '1', username: 'admin', isOnline: true, isBanned: false },
			{ id: '2', username: 'johndoe', isOnline: true, isBanned: false },
			{ id: '3', username: 'spammer99', isOnline: false, isBanned: true },
			{ id: '4', username: 'toxic_guy', isOnline: false, isBanned: false },
			{ id: '5', username: 'neon_ninja', isOnline: true, isBanned: false },
		];
	};

	const handleBanUser = (user: User) => {
		const action = user.isBanned ? 'разбанить' : 'забанить';
		if (!confirm(`Точно хочешь ${action} пользователя ${user.username}?`)) return;
		
		user.isBanned = !user.isBanned;
	};

	const goBack = () => {
		router.push('/'); 
	};
</script>

<template>
	<div class="app-layout" :class="{ 'light-theme': !isDarkTheme }">
		
		<div v-if="!isAuthenticated" class="auth-overlay">
			<button class="theme-btn absolute-theme-btn" @click="toggleTheme" title="Переключить тему">
				<svg v-if="isDarkTheme" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="5"></circle><line x1="12" y1="1" x2="12" y2="3"></line><line x1="12" y1="21" x2="12" y2="23"></line><line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line><line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line><line x1="1" y1="12" x2="3" y2="12"></line><line x1="21" y1="12" x2="23" y2="12"></line><line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line><line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line></svg>
				<svg v-else width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path></svg>
			</button>
			<button class="back-btn absolute-back-btn" @click="goBack">На главную</button>

			<div class="auth-box">
				<h2>Доступ ограничен</h2>
				<p>Введите секретный ключ администратора</p>
				<input 
					type="password" 
					v-model="adminKeyInput" 
					@keyup.enter="handleLogin"
					placeholder="Секретный ключ..."
					class="auth-input"
				>
				<span class="error-msg" v-if="authError">{{ authError }}</span>
				<button class="auth-btn" @click="handleLogin">Войти</button>
			</div>
		</div>

		<template v-else>
			<div class="chat-header">
				<div class="header-left">
					<button class="back-btn" @click="goBack">
						<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
							<line x1="19" y1="12" x2="5" y2="12"></line>
							<polyline points="12 19 5 12 12 5"></polyline>
						</svg>
						<span class="hide-mobile">Назад в чат</span>
					</button>
				</div>
				<div class="header-center brand">
					<span class="brand-text">Admin Panel</span>
				</div>
				<div class="header-right">
					<button class="theme-btn" @click="toggleTheme" title="Переключить тему">
						<svg v-if="isDarkTheme" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="5"></circle><line x1="12" y1="1" x2="12" y2="3"></line><line x1="12" y1="21" x2="12" y2="23"></line><line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line><line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line><line x1="1" y1="12" x2="3" y2="12"></line><line x1="21" y1="12" x2="23" y2="12"></line><line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line><line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line></svg>
						<svg v-else width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path></svg>
					</button>
				</div>
			</div>

			<div class="admin-content">
				<div class="stats-grid">
					<div class="stat-card">
						<div class="stat-title">Пользователей</div>
						<div class="stat-value">{{ totalRegistered }}</div>
					</div>
					<div class="stat-card">
						<div class="stat-title">Сейчас онлайн</div>
						<div class="stat-value text-accent">{{ onlineCount }}</div>
					</div>
					<div class="stat-card">
						<div class="stat-title">Всего сообщений</div>
						<div class="stat-value text-blue">{{ totalMessages }}</div>
					</div>
				</div>

				<div class="users-section">
					<h2>Управление пользователями</h2>
					<div class="table-container">
						<table class="users-table">
							<thead>
								<tr>
									<th>Пользователь</th>
									<th>Статус</th>
									<th>Действия</th>
								</tr>
							</thead>
							<tbody>
								<tr v-for="user in users" :key="user.id" :class="{ 'is-banned': user.isBanned }">
									<td class="user-name">
										<span class="avatar">{{ user.username.charAt(0).toUpperCase() }}</span>
										<span class="u-name-text">{{ user.username }}</span>
									</td>
									<td>
										<div class="status-badge" :class="user.isOnline ? 'online' : 'offline'">
											{{ user.isOnline ? 'Онлайн' : 'Оффлайн' }}
										</div>
										<div v-if="user.isBanned" class="banned-text">Забанен</div>
									</td>
									<td>
										<button 
											class="action-btn" 
											:class="user.isBanned ? 'unban-btn' : 'ban-btn'"
											@click="handleBanUser(user)"
										>
											{{ user.isBanned ? 'Разбанить' : 'Бан' }}
										</button>
									</td>
								</tr>
							</tbody>
						</table>
					</div>
				</div>
			</div>
		</template>

	</div>
</template>

<style scoped>
/* ПЕРЕМЕННЫЕ */
.app-layout {
	--bg-base: #050807;
	--bg-surface: #080b0a;
	--border: #0f1714;
	--bg-active: #1a231e;
	--text-primary: #c8c2b8;
	--text-secondary: #64615c;
	--text-muted: #8a867f;
	--accent: #5fca08;
	--accent-hover: #4da806;
	--danger: #dc3545;
	--blue: #3b82f6;
	
	display: flex;
	flex-direction: column;
	height: 100dvh;
	background-color: var(--bg-base);
	color: var(--text-primary);
	transition: all 0.3s ease;
	overflow: hidden;
}

.app-layout.light-theme {
	--bg-base: #f4f7f6;
	--bg-surface: #ffffff;
	--border: #e2e8e4;
	--bg-active: #e8f5e9;
	--text-primary: #1a231e;
	--text-secondary: #6b7280;
	--text-muted: #4b5563;
	--accent: #5fca08; 
	--blue: #2563eb;
}

.auth-overlay {
	flex: 1;
	display: flex;
	align-items: center;
	justify-content: center;
	position: relative;
	padding: 20px;
}
.absolute-theme-btn { position: absolute; top: 20px; right: 20px; }
.absolute-back-btn { position: absolute; top: 20px; left: 20px; }

.auth-box {
	background: var(--bg-surface);
	border: 1px solid var(--border);
	padding: 40px;
	border-radius: 16px;
	width: 100%;
	max-width: 400px;
	text-align: center;
	box-shadow: 0 10px 30px rgba(0,0,0,0.2);
}
.app-layout.light-theme .auth-box { box-shadow: 0 10px 30px rgba(0,0,0,0.05); }
.auth-box h2 { margin-bottom: 10px; color: var(--accent); }
.auth-box p { color: var(--text-secondary); margin-bottom: 24px; font-size: 14px; }

.auth-input {
	width: 100%;
	padding: 12px 16px;
	background: var(--bg-base);
	border: 1px solid var(--border);
	color: var(--text-primary);
	border-radius: 8px;
	margin-bottom: 16px;
	font-size: 16px;
	outline: none;
	box-sizing: border-box;
}
.auth-input:focus { border-color: var(--accent); }

.error-msg { display: block; color: var(--danger); font-size: 13px; margin-top: -10px; margin-bottom: 16px; }

.auth-btn {
	width: 100%;
	padding: 12px;
	background: var(--accent);
	color: white;
	border: none;
	border-radius: 8px;
	font-size: 16px;
	font-weight: bold;
	cursor: pointer;
	transition: background 0.2s;
}
.auth-btn:hover { background: var(--accent-hover); }

.chat-header {
	height: 70px; padding: 0 24px; border-bottom: 1px solid var(--border);
	display: flex; justify-content: space-between; align-items: center;
	background-color: var(--bg-surface); flex-shrink: 0;
}
.header-left, .header-right { flex: 1; display: flex; align-items: center; }
.header-right { justify-content: flex-end; }
.header-center { flex: 1; display: flex; justify-content: center; }

.back-btn {
	background: transparent; border: none; color: var(--text-secondary);
	display: flex; align-items: center; gap: 8px; cursor: pointer;
	font-size: 15px; font-weight: 500; transition: color 0.2s;
}
.back-btn:hover { color: var(--text-primary); }
.brand-text { font-size: 20px; font-weight: 700; color: var(--text-primary); }

.theme-btn {
	background: transparent; border: none; cursor: pointer; padding: 6px;
	border-radius: 8px; transition: all 0.2s ease; display: flex;
	align-items: center; justify-content: center; color: #64615c; 
}
.theme-btn:hover { background: var(--border); color: var(--accent); }
.app-layout.light-theme .theme-btn { color: #8a867f; }
.app-layout.light-theme .theme-btn:hover { color: var(--accent); background: #e2e8e4; }

.admin-content {
	padding: 30px 24px;
	max-width: 1000px;
	margin: 0 auto;
	width: 100%;
	overflow-y: auto;
	box-sizing: border-box;
}

.stats-grid {
	display: grid;
	grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
	gap: 20px;
	margin-bottom: 40px;
}
.stat-card {
	background: var(--bg-surface); border: 1px solid var(--border);
	border-radius: 12px; padding: 24px; display: flex; flex-direction: column;
	align-items: center; justify-content: center; text-align: center;
}
.stat-title { color: var(--text-secondary); font-size: 13px; text-transform: uppercase; font-weight: 600; margin-bottom: 8px; }
.stat-value { font-size: 32px; font-weight: 700; }
.text-accent { color: var(--accent); }
.text-blue { color: var(--blue); }

.users-section h2 { margin-bottom: 20px; font-size: 20px; }

.table-container {
	background: var(--bg-surface);
	border: 1px solid var(--border);
	border-radius: 12px;
	overflow-x: auto;
}

.users-table { width: 100%; border-collapse: collapse; text-align: left; min-width: 500px; }
.users-table th, .users-table td { padding: 16px; border-bottom: 1px solid var(--border); }
.users-table th { color: var(--text-secondary); font-weight: 600; font-size: 14px; white-space: nowrap; }
.users-table tr:last-child td { border-bottom: none; }
.users-table tr:hover { background: var(--bg-active); }
.users-table tr.is-banned td { opacity: 0.6; }

.user-name { display: flex; align-items: center; gap: 12px; font-weight: 500; }
.avatar {
	width: 32px; height: 32px; border-radius: 8px; background: var(--bg-active);
	display: flex; align-items: center; justify-content: center;
	font-weight: bold; border: 1px solid var(--border); flex-shrink: 0;
}
.u-name-text { word-break: break-all; }

.status-badge { display: inline-flex; padding: 4px 8px; border-radius: 6px; font-size: 12px; font-weight: 600; }
.status-badge.online { background: rgba(95, 202, 8, 0.1); color: var(--accent); }
.status-badge.offline { background: var(--bg-active); color: var(--text-muted); }
.banned-text { color: var(--danger); font-size: 12px; font-weight: bold; margin-top: 4px; }

.action-btn { padding: 8px 16px; border-radius: 8px; border: none; font-weight: 600; cursor: pointer; transition: all 0.2s; white-space: nowrap; }
.ban-btn { background: rgba(220, 53, 69, 0.1); color: var(--danger); }
.ban-btn:hover { background: var(--danger); color: white; }
.unban-btn { background: var(--border); color: var(--text-primary); }
.unban-btn:hover { background: var(--text-secondary); color: white; }

@media (max-width: 600px) {
	.chat-header { padding: 0 16px; }
	.hide-mobile { display: none; }
	
	.admin-content { padding: 20px 16px; }
	
	.stats-grid { grid-template-columns: 1fr; gap: 12px; }
	.stat-card { padding: 16px; }
	.stat-value { font-size: 28px; }

	.users-table th, .users-table td { padding: 12px; }
	.action-btn { padding: 8px 12px; font-size: 13px; }
	
	.auth-box { padding: 30px 20px; border-radius: 12px; }
}
</style>