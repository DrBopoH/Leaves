<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/App.vue
import { ref, onMounted } from 'vue';
import { RouterView } from 'vue-router';
import { fetchMe } from './api/auth';
import { useUserStore } from './stores/user';
import { useTheme } from './composables/useTheme';
import UiLoader from './components/ui-kit/UiLoader.vue';

useTheme();

const userStore = useUserStore();
const isAuthReady = ref(false);

onMounted(async () => {
	try {
		const user = await fetchMe();
		if (user) {
			userStore.setUser(user);
		}
	} catch (error) {
		console.error("Failed to fetch user session", error);
	} finally {
		isAuthReady.value = true;
	}
});
</script>

<template>
	<div v-if="!isAuthReady" class="app-loader-screen">
		<UiLoader size="large" color="var(--color-accent)" />
	</div>
	<RouterView v-else />
</template>

<style scoped>
.app-loader-screen {
	display: flex;
	align-items: center;
	justify-content: center;
	min-height: 100vh;
	width: 100%;
	background-color: var(--color-bg);
}
</style>
