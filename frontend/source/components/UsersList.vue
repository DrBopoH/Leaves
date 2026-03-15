<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/UsersList.vue
import { ref, computed } from 'vue';
import { useUserStore } from '../stores/user';
import UiUserItem from './ui-kit/UiUserItem.vue';

const props = defineProps<{
    users: { username: string; online: boolean; last_seen?: string }[];
}>();

const userStore = useUserStore();

const usersPage = ref(1);
const isLoadingUsers = ref(false);
let loadUsersTimeout: ReturnType<typeof setTimeout> | null = null;
let lastLoadTime = 0;
const PAGE_SIZE = 30;

// Использование computed вместо ref + watch избавляет от рассинхрона
// и гарантирует сохранение глубокой реактивности массива.
const visibleUsers = computed(() => {
    const end = usersPage.value * PAGE_SIZE;
    const start = Math.max(0, end - 90); // Keep max 90 users in DOM
    return props.users.slice(start, end);
});

const handleUsersScroll = (e: Event) => {
    const target = e.target as HTMLElement;

    if (target.scrollHeight - target.scrollTop <= target.clientHeight + 10) {
        if (usersPage.value * PAGE_SIZE < props.users.length && !isLoadingUsers.value) {
            const now = Date.now();
            const timeSinceLastLoad = now - lastLoadTime;
            const delay = timeSinceLastLoad > 800 ? 0 : 800 - timeSinceLastLoad;

            isLoadingUsers.value = true;

            if (loadUsersTimeout) clearTimeout(loadUsersTimeout);
            loadUsersTimeout = setTimeout(() => {
                usersPage.value++;
                isLoadingUsers.value = false;
                lastLoadTime = Date.now();
            }, delay);
        }
    }

    if (target.scrollTop <= 10 && usersPage.value > 3) {
        usersPage.value--;
        target.scrollTop = 50;
    }
};
const isUserOnline = (u: { username: string; online: boolean }) => {
    return u.online || u.username === (userStore as any).user?.username;
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
</script>

<template>
	<div class="members-list" @scroll="handleUsersScroll">
		<UiUserItem
			v-for="u in visibleUsers"
			:key="u.username"
			:username="u.username"
			:color="userStore.getUserColor(u.username)"
			:status="isUserOnline(u) ? 'online' : 'offline'"
			:subtitle="isUserOnline(u) ? 'Online' : formatRelativeTime(u.last_seen)"
			interactive
		/>

		<div v-if="isLoadingUsers" class="loading-users">
			<span class="dot"></span>
			<span class="dot"></span>
			<span class="dot"></span>
		</div>
	</div>
</template>

<style scoped>
.members-list {
	flex: 1;
	overflow-y: auto;
	padding: 16px;
	display: flex;
	flex-direction: column;
	gap: 12px;
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
	background-color: var(--color-text-secondary);
	border-radius: 50%;
	animation: bounce 1.4s infinite ease-in-out both;
}

.loading-users .dot:nth-child(1) { animation-delay: -0.32s; }
.loading-users .dot:nth-child(2) { animation-delay: -0.16s; }

@keyframes bounce {
	0%, 80%, 100% { transform: scale(0); }
	40% { transform: scale(1); }
}
</style>
