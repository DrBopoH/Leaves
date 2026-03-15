// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/composables/useTheme.ts
import { ref, watch, computed } from 'vue';

type Theme = 'dark' | 'light';

const currentTheme = ref<Theme>((localStorage.getItem('theme') as Theme) || 'dark');

export function useTheme() {
	const isDark = computed(
		() => currentTheme.value === 'dark'
	);

	const toggleTheme = () => {
		currentTheme.value = isDark.value ? 'light' : 'dark';
	};

	watch(currentTheme, (newTheme) => {
			if (newTheme === 'dark') {
				document.documentElement.removeAttribute('data-theme');
			} else {
				document.documentElement.setAttribute('data-theme', newTheme);
			}
			localStorage.setItem('theme', newTheme);
		}, 
		{ 
			immediate: true 
		}
	);

	return {
		isDark,
		toggleTheme
	};
}
