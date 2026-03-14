<script setup lang="ts">
	// Copyright (C) 2026 MorangTong Creative Studio
	// SPDX-License-Identifier: AGPL-3.0-or-later

	// source/components/debug/DebugGrid.vue
	import { ref } from 'vue';
	import UiBoard from '../ui-kit/UiBoard.vue';
	import UiAlert from '../ui-kit/UiAlert.vue';

	const isAlertVisible = ref(false);

	const toggleAlert = () => {
		isAlertVisible.value = !isAlertVisible.value;
	};
</script>

<template>
	<UiBoard>
		<div class="debug-content">
			<UiAlert show type="info">Сдвинь меня средней кнопкой мыши!</UiAlert>
			<UiAlert :show="isAlertVisible" type="error">Критическая ошибка!</UiAlert>
			<button @click="toggleAlert">Показать / Скрыть ошибку</button>
		</div>
	</UiBoard>
</template>

<style scoped>
	.debug-content {
		padding: 40px;
		display: flex;
		flex-direction: column;
		gap: 20px;
	}

	.debug-content button, 
	.debug-content .ui-alert {
		cursor: default;
	}
</style>