<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiLoader.vue

const props = defineProps<{
	size?: 'small' | 'medium' | 'large';
	color?: string;
}>();
</script>

<template>
	<div
		class="ui-loader"
		:class="`size-${size || 'medium'}`"
		:style="{ color: color || 'currentColor' }"
	>
		<svg viewBox="0 0 24 24" fill="currentColor">
			<circle cx="4" cy="12" r="3">
				<animate id="a1" begin="0;a3.end-0.25s" attributeName="r" dur="0.75s" values="3;.2;3"/>
			</circle>
			<circle cx="12" cy="12" r="3">
				<animate begin="a1.begin+0.15s" attributeName="r" dur="0.75s" values="3;.2;3"/>
			</circle>
			<circle cx="20" cy="12" r="3">
				<animate id="a3" begin="a1.begin+0.3s" attributeName="r" dur="0.75s" values="3;.2;3"/>
			</circle>
		</svg>
	</div>
</template>

<style scoped>
.ui-loader {
	display: inline-flex;
	align-items: center;
	justify-content: center;
}

.ui-loader.size-small {
	width: 16px;
	height: 16px;
}

.ui-loader.size-medium {
	width: 24px;
	height: 24px;
}

.ui-loader.size-large {
	width: 32px;
	height: 32px;
}

.ui-loader svg {
	width: 100%;
	height: 100%;
}
</style>
