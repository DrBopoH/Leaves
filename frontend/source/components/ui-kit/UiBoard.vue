<script setup lang="ts">
	// Copyright (C) 2026 MorangTong Creative Studio
	// SPDX-License-Identifier: AGPL-3.0-or-later

	// source/components/ui-kit/UiBoard.vue
	import { ref, onMounted, onUnmounted } from 'vue';

	const isDragging = ref(false);
	const boardX = ref(0);
	const boardY = ref(0);
	const scale = ref(1);

	const onWheel = (event: WheelEvent) => {
		const wheel = event.deltaY < 0 ? 1 : -1;
		let newScale = scale.value * Math.exp(wheel * 0.1);
		newScale = Math.max(0.1, Math.min(newScale, 5));

		const ratio = newScale / scale.value;

		boardX.value = event.clientX - (event.clientX - boardX.value) * ratio;
		boardY.value = event.clientY - (event.clientY - boardY.value) * ratio;
		scale.value = newScale;
	};

	const onMouseDown = () => {
		isDragging.value = true;
	};

	const onMouseMove = (event: MouseEvent) => {
		if (!isDragging.value) return;

		boardX.value += event.movementX;
		boardY.value += event.movementY;
	};

	const onMouseUp = () => {
		isDragging.value = false;
	};

	onMounted(() => {
		window.addEventListener('mousemove', onMouseMove);
		window.addEventListener('mouseup', onMouseUp);
	});

	onUnmounted(() => {
		window.removeEventListener('mousemove', onMouseMove);
		window.removeEventListener('mouseup', onMouseUp);
	});
</script>

<template>
	<div
		class="ui-board-wrapper"
		:class="{ dragging: isDragging }"
		:style="{ 
			backgroundPosition: `${boardX}px ${boardY}px`, 
			backgroundSize: `${20 * scale}px ${20 * scale}px` 
		}"
		@mousedown.middle.prevent="onMouseDown"
		@wheel.prevent="onWheel"
	>
		<div
			class="ui-board"
			:style="{ 
				transform: `translate(${boardX}px, ${boardY}px) scale(${scale})` 
			}"
		>
			<slot />
		</div>
	</div>
</template>

<style scoped>
	.ui-board-wrapper {
		flex: 1;
		overflow: hidden;
		position: relative;
		background-color: var(--color-surface);
		background-image: radial-gradient(var(--color-accent-shadow) 1px, transparent 1px);
		cursor: grab;
	}

	.ui-board-wrapper.dragging {
		cursor: grabbing;
	}

	.ui-board {
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		transform-origin: 0 0;
		will-change: transform;
		transition: none !important; 
	}
</style>