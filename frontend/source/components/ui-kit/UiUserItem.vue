<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiUserItem.vue

import UiAvatar from './UiAvatar.vue';
import UiStatusDot from './UiStatusDot.vue';

const props = defineProps<{
	username: string;
	color?: string;
	status?: 'online' | 'offline' | 'away' | 'dnd';
	subtitle?: string;
	interactive?: boolean;
}>();
</script>

<template>
	<div
		class="ui-user-item"
		:class="{ 'is-interactive': interactive }"
	>
		<UiAvatar
			:name="username"
			:color="color"
			size="small"
			class="user-avatar"
		/>

		<div class="user-info">
			<span class="user-name" :style="{ color: color || 'inherit' }">
				{{ username }}
			</span>

			<div v-if="subtitle || status" class="user-subtitle-row">
				<UiStatusDot v-if="status" :status="status" />
				<span v-if="subtitle" class="user-subtitle">{{ subtitle }}</span>
			</div>
		</div>
	</div>
</template>

<style scoped>
.ui-user-item {
	display: flex;
	align-items: center;
	gap: 10px;
	padding: 6px 8px;
	border-radius: 8px;
	transition: background-color 0.2s ease;
	background-color: transparent;
}

.ui-user-item.is-interactive {
	cursor: pointer;
}

.ui-user-item.is-interactive:hover {
	background-color: var(--color-surface-hover);
}

.user-avatar {
	flex-shrink: 0;
}

.user-info {
	display: flex;
	flex-direction: column;
	min-width: 0;
}

.user-name {
	font-size: 14px;
	font-weight: 600;
	white-space: nowrap;
	overflow: hidden;
	text-overflow: ellipsis;
}

.user-subtitle-row {
	display: flex;
	align-items: center;
	gap: 6px;
	margin-top: 2px;
}

.user-subtitle {
	font-size: 11px;
	color: var(--color-text-secondary);
	white-space: nowrap;
	overflow: hidden;
	text-overflow: ellipsis;
}
</style>
