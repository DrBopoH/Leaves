// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/stores/user.ts
import { ref } from 'vue';
import { defineStore } from 'pinia';

export const useUserStore = defineStore('user', () => {
	const currentUser = ref<{ id: number; username: string } | null>(null);

	function setUser(user: { id: number; username: string } | null) {
		currentUser.value = user;
	}

	const getUserColor = (username: string) => {
		if (!username) return 'var(--color-accent)';
		let hash = 0;
		for (let i = 0; i < username.length; i++) {
			hash = username.charCodeAt(i) + ((hash << 5) - hash);
		}
		const hue = Math.abs(hash % 360);
		return `hsl(${hue}, 80%, var(--avatar-lightness, 65%))`;
	};

	return { currentUser, setUser, getUserColor };
});
