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
		:style="color ? { color } : {}"
	>
		<div class="dot"></div>
		<div class="dot"></div>
		<div class="dot"></div>
	</div>
</template>

<style scoped>
	.ui-loader {
		display: inline-flex;
		align-items: center;
		justify-content: center;
	}

	.size-small { 
		gap: 3px; 
	}
	.size-small 
	.dot { 
		width: 4px; 
		height: 4px; 
	}

	.size-medium { 
		gap: 4px; 
	}
	.size-medium 
	.dot { 
		width: 6px; 
		height: 6px; 
	}

	.size-large { 
		gap: 6px; 
	}
	.size-large 
	.dot { 
		width: 8px; 
		height: 8px; 
	}

	.dot {
		background-color: currentColor;
		border-radius: 50%;
		animation: loader-pulse 1.4s infinite ease-in-out both;
	}

	.dot:nth-child(1) { 
		animation-delay: -0.32s; 
	}
	.dot:nth-child(2) { 
		animation-delay: -0.16s; 
	}
	.dot:nth-child(3) { 
		animation-delay: 0s; 
	}

	@keyframes loader-pulse {
		0%, 80%, 100% {
			transform: scale(0);
			opacity: 0.3;
		}
		40% {
			transform: scale(1);
			opacity: 1;
		}
	}
</style>