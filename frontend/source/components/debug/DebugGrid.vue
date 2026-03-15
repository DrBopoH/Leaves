<script setup lang="ts">
	// Copyright (C) 2026 MorangTong Creative Studio
	// SPDX-License-Identifier: AGPL-3.0-or-later

	// source/components/debug/DebugGrid.vue
	import { ref } from 'vue';
	import DebugShowcase from './DebugShowcase.vue';
	import DebugSection from './DebugSection.vue';
	import UiBoard from '../ui-kit/UiBoard.vue';
	

	import UiAlert from '../ui-kit/UiAlert.vue';	
	
	const isAlertVisible = ref(false);
	
	const toggleAlert = () => {
		isAlertVisible.value = !isAlertVisible.value;
	};
</script>

<template>
	<UiBoard>
		<DebugShowcase>
			<DebugSection id="alerts" title="UiAlert">
				<div>
					<UiAlert show type="info">Move me with the middle mouse button!</UiAlert>
					<UiAlert :show="isAlertVisible" type="error">Critical error!</UiAlert>
					<button @click="toggleAlert">Show / Hide Error</button>
				</div>
			</DebugSection>

			<DebugSection id="buttons" title="UiButton">
				<div>Buttons will be here...</div>
			</DebugSection>

			<DebugSection id="inputs" title="UiInput">
				<div>Inputs will be here...</div>
			</DebugSection>
		</DebugShowcase>
	</UiBoard>
</template>