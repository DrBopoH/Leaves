<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/AppHeader.vue
import { useRouter } from 'vue-router';
import { useTheme } from '../composables/useTheme';
import UiLogo from './ui-kit/UiLogo.vue';
import UiIconButton from './ui-kit/UiIconButton.vue';

const router = useRouter();
const { isDark: isDarkTheme, toggleTheme } = useTheme();
</script>

<template>
	<div class="top-header">
		<div class="header-left"></div>

		<div
			class="header-center"
			@click="router.push('/')"
			style="cursor: pointer;"
			title="Home"
		>
			<UiLogo size="small" text="Leaves" />
		</div>

		<div class="header-right">
			<UiIconButton
				@click="toggleTheme"
				title="Toggle Theme"
			>
				<svg v-if="isDarkTheme" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
					<circle cx="12" cy="12" r="5"></circle>
					<line x1="12" y1="1" x2="12" y2="3"></line>
					<line x1="12" y1="21" x2="12" y2="23"></line>
					<line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line>
					<line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line>
					<line x1="1" y1="12" x2="3" y2="12"></line>
					<line x1="21" y1="12" x2="23" y2="12"></line>
					<line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line>
					<line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line>
				</svg>
				<svg v-else viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
					<path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path>
				</svg>
			</UiIconButton>
		</div>
	</div>
</template>

<style scoped>
.top-header {
	height: 36px;
	padding: 0 16px;
	border-bottom: 1px solid var(--color-border);
	display: flex;
	justify-content: space-between;
	align-items: center;
	background-color: var(--color-surface);
	z-index: 10;
	flex-shrink: 0;
}

.header-left,
.header-right {
	display: flex;
	align-items: center;
	gap: 10px;
	flex: 1;
}

.header-right {
	justify-content: flex-end;
}

.header-center {
	display: flex;
	justify-content: center;
	align-items: center;
}
</style>
