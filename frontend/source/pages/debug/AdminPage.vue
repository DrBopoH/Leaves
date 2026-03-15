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
			authError.value = 'Invalid access key';
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
		const action = user.isBanned ? 'unban' : 'ban';
		if (!confirm(`Are you sure you want to ${action} user ${user.username}?`)) return;

		user.isBanned = !user.isBanned;
	};

	const goBack = () => {
		router.push('/');
	};
</script>

<template>
	<div class="app-layout" :class="{ 'light-theme': !isDarkTheme }">
		<div v-if="!isAuthenticated" class="auth-overlay">
			<button class="theme-btn absolute-theme-btn" @click="toggleTheme" title="Toggle theme">
				<svg v-if="isDarkTheme" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="5"></circle><line x1="12" y1="1" x2="12" y2="3"></line><line x1="12" y1="21" x2="12" y2="23"></line><line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line><line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line><line x1="1" y1="12" x2="3" y2="12"></line><line x1="21" y1="12" x2="23" y2="12"></line><line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line><line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line></svg>
				<svg v-else width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path></svg>
			</button>
			<button class="back-btn absolute-back-btn" @click="goBack">Go to Home</button>

			<div class="auth-box">
				<h2>Access Restricted</h2>
				<p>Enter the secret admin key</p>
				<input
					type="password"
					v-model="adminKeyInput"
					@keyup.enter="handleLogin"
					placeholder="Secret key..."
					class="auth-input"
				/>
				<span class="error-msg" v-if="authError">{{ authError }}</span>
				<button class="auth-btn" @click="handleLogin">Login</button>
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
						<span class="hide-mobile">Back to Chat</span>
					</button>
				</div>
				<div class="header-center brand">
					<span class="brand-text">Admin Panel</span>
				</div>
				<div class="header-right">
					<button class="theme-btn" @click="toggleTheme" title="Toggle theme">
						<svg v-if="isDarkTheme" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="5"></circle><line x1="12" y1="1" x2="12" y2="3"></line><line x1="12" y1="21" x2="12" y2="23"></line><line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line><line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line><line x1="1" y1="12" x2="3" y2="12"></line><line x1="21" y1="12" x2="23" y2="12"></line><line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line><line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line></svg>
						<svg v-else width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path></svg>
					</button>
				</div>
			</div>

			<div class="admin-content">
				<div class="stats-grid">
					<div class="stat-card">
						<div class="stat-title">Users</div>
						<div class="stat-value">{{ totalRegistered }}</div>
					</div>
					<div class="stat-card">
						<div class="stat-title">Online Now</div>
						<div class="stat-value text-accent">{{ onlineCount }}</div>
					</div>
					<div class="stat-card">
						<div class="stat-title">Total Messages</div>
						<div class="stat-value text-blue">{{ totalMessages }}</div>
					</div>
				</div>

				<div class="users-section">
					<h2>User Management</h2>
					<div class="table-container">
						<table class="users-table">
							<thead>
								<tr>
									<th>User</th>
									<th>Status</th>
									<th>Actions</th>
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
											{{ user.isOnline ? 'Online' : 'Offline' }}
										</div>
										<div v-if="user.isBanned" class="banned-text">Banned</div>
									</td>
									<td>
										<button
											class="action-btn"
											:class="user.isBanned ? 'unban-btn' : 'ban-btn'"
											@click="handleBanUser(user)"
										>
											{{ user.isBanned ? 'Unban' : 'Ban' }}
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
	.app-layout {
		display: flex;
		height: 100dvh;
		flex-direction: column;
		overflow: hidden;
		color: var(--color-text-primary);
		background-color: var(--color-bg-primary);
		transition: all 0.3s ease;
	}

	.auth-overlay {
		position: relative;
		display: flex;
		flex: 1;
		align-items: center;
		justify-content: center;
		padding: 20px;
	}
	.absolute-theme-btn {
		position: absolute;
		top: 20px;
		right: 20px;
	}
	.absolute-back-btn {
		position: absolute;
		top: 20px;
		left: 20px;
	}

	.auth-box {
		width: 100%;
		max-width: 400px;
		padding: 40px;
		text-align: center;
		background: var(--color-bg-secondary);
		border: 1px solid var(--color-border);
		border-radius: 16px;
		box-shadow: var(--shadow-xl);
	}
	.auth-box h2 {
		margin-bottom: 10px;
		color: var(--color-primary);
	}
	.auth-box p {
		margin-bottom: 24px;
		color: var(--color-text-secondary);
		font-size: var(--font-size-sm);
	}

	.auth-input {
		width: 100%;
		padding: 12px 16px;
		margin-bottom: 16px;
		color: var(--color-text-primary);
		background: var(--color-bg-primary);
		border: 1px solid var(--color-border);
		border-radius: 8px;
		font-size: var(--font-size-lg);
		outline: none;
		box-sizing: border-box;
	}
	.auth-input:focus {
		border-color: var(--color-primary);
	}

	.error-msg {
		display: block;
		margin-top: -10px;
		margin-bottom: 16px;
		color: var(--color-error);
		font-size: var(--font-size-sm);
	}

	.auth-btn {
		width: 100%;
		padding: 12px;
		color: var(--color-text-contrast);
		background: var(--color-primary);
		border: none;
		border-radius: 8px;
		font-size: var(--font-size-lg);
		font-weight: var(--font-weight-bold);
		cursor: pointer;
		transition: background 0.2s;
	}
	.auth-btn:hover {
		background: var(--color-primary-hover);
	}

	.chat-header {
		display: flex;
		height: 70px;
		flex-shrink: 0;
		align-items: center;
		justify-content: space-between;
		padding: 0 24px;
		background-color: var(--color-bg-secondary);
		border-bottom: 1px solid var(--color-border);
	}
	.header-left,
	.header-right {
		display: flex;
		flex: 1;
		align-items: center;
	}
	.header-right {
		justify-content: flex-end;
	}
	.header-center {
		display: flex;
		flex: 1;
		justify-content: center;
	}

	.back-btn {
		display: flex;
		align-items: center;
		gap: 8px;
		color: var(--color-text-secondary);
		font-size: 15px;
		font-weight: var(--font-weight-medium);
		background: transparent;
		border: none;
		cursor: pointer;
		transition: color 0.2s;
	}
	.back-btn:hover {
		color: var(--color-text-primary);
	}
	.brand-text {
		color: var(--color-text-primary);
		font-size: 20px;
		font-weight: var(--font-weight-bold);
	}

	.theme-btn {
		display: flex;
		align-items: center;
		justify-content: center;
		padding: 6px;
		color: var(--color-text-tertiary);
		background: transparent;
		border: none;
		border-radius: 8px;
		cursor: pointer;
		transition: all 0.2s ease;
	}
	.theme-btn:hover {
		color: var(--color-primary);
		background: var(--color-border);
	}

	.admin-content {
		width: 100%;
		max-width: 1000px;
		margin: 0 auto;
		padding: 30px 24px;
		overflow-y: auto;
		box-sizing: border-box;
	}

	.stats-grid {
		display: grid;
		gap: 20px;
		grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
		margin-bottom: 40px;
	}
	.stat-card {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		padding: 24px;
		text-align: center;
		background: var(--color-bg-secondary);
		border: 1px solid var(--color-border);
		border-radius: 12px;
	}
	.stat-title {
		margin-bottom: 8px;
		color: var(--color-text-secondary);
		font-size: var(--font-size-sm);
		font-weight: var(--font-weight-semibold);
		text-transform: uppercase;
	}
	.stat-value {
		font-size: 32px;
		font-weight: var(--font-weight-bold);
	}
	.text-accent {
		color: var(--color-primary);
	}
	.text-blue {
		color: var(--color-info);
	}

	.users-section h2 {
		margin-bottom: 20px;
		font-size: 20px;
	}

	.table-container {
		overflow-x: auto;
		background: var(--color-bg-secondary);
		border: 1px solid var(--color-border);
		border-radius: 12px;
	}

	.users-table {
		width: 100%;
		min-width: 500px;
		text-align: left;
		border-collapse: collapse;
	}
	.users-table th,
	.users-table td {
		padding: 16px;
		border-bottom: 1px solid var(--color-border);
	}
	.users-table th {
		color: var(--color-text-secondary);
		font-size: var(--font-size-sm);
		font-weight: var(--font-weight-semibold);
		white-space: nowrap;
	}
	.users-table tr:last-child td {
		border-bottom: none;
	}
	.users-table tr:hover {
		background: var(--color-bg-tertiary);
	}
	.users-table tr.is-banned td {
		opacity: 0.6;
	}

	.user-name {
		display: flex;
		align-items: center;
		gap: 12px;
		font-weight: var(--font-weight-medium);
	}
	.avatar {
		display: flex;
		width: 32px;
		height: 32px;
		flex-shrink: 0;
		align-items: center;
		justify-content: center;
		background: var(--color-bg-tertiary);
		border: 1px solid var(--color-border);
		border-radius: 8px;
		font-weight: var(--font-weight-bold);
	}
	.u-name-text {
		word-break: break-all;
	}

	.status-badge {
		display: inline-flex;
		padding: 4px 8px;
		font-size: var(--font-size-xs);
		font-weight: var(--font-weight-semibold);
		border-radius: 6px;
	}
	.status-badge.online {
		color: var(--color-success);
		background: rgba(34, 197, 94, 0.1);
	}
	.status-badge.offline {
		color: var(--color-text-tertiary);
		background: var(--color-bg-tertiary);
	}
	.banned-text {
		margin-top: 4px;
		color: var(--color-error);
		font-size: var(--font-size-xs);
		font-weight: var(--font-weight-bold);
	}

	.action-btn {
		padding: 8px 16px;
		font-weight: var(--font-weight-semibold);
		border: none;
		border-radius: 8px;
		cursor: pointer;
		white-space: nowrap;
		transition: all 0.2s;
	}
	.ban-btn {
		color: var(--color-error);
		background: rgba(239, 68, 68, 0.1);
	}
	.ban-btn:hover {
		color: white;
		background: var(--color-error);
	}
	.unban-btn {
		color: var(--color-text-primary);
		background: var(--color-border);
	}
	.unban-btn:hover {
		color: white;
		background: var(--color-text-secondary);
	}

	@media (max-width: 600px) {
		.chat-header {
			padding: 0 16px;
		}
		.hide-mobile {
			display: none;
		}

		.admin-content {
			padding: 20px 16px;
		}

		.stats-grid {
			gap: 12px;
			grid-template-columns: 1fr;
		}
		.stat-card {
			padding: 16px;
		}
		.stat-value {
			font-size: 28px;
		}

		.users-table th,
		.users-table td {
			padding: 12px;
		}
		.action-btn {
			padding: 8px 12px;
			font-size: var(--font-size-sm);
		}

		.auth-box {
			padding: 30px 20px;
			border-radius: 12px;
		}
	}
</style>
