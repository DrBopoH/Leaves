<script setup lang="ts">
	// Copyright (C) 2026 MorangTong Creative Studio
	// SPDX-License-Identifier: AGPL-3.0-or-later

	// source/components/ui-kit/UiBoard.vue
	import { ref, onUnmounted } from 'vue';

	const props = defineProps({
		minZoom: { type: Number, default: 0.1 },
		maxZoom: { type: Number, default: 5 },
		zoomSpeed: { type: Number, default: 0.1 },
		gridSize: { type: Number, default: 20 },
	});

	const wrapperRef = ref<HTMLElement | null>(null);
	const isDragging = ref(false);
	const boardX = ref(0);
	const boardY = ref(0);
	const scale = ref(1);

	const onWheel = (event: WheelEvent) => {
		if (!wrapperRef.value) return;

		const rect = wrapperRef.value.getBoundingClientRect();
		const mouseX = event.clientX - rect.left;
		const mouseY = event.clientY - rect.top;

		const wheel = event.deltaY < 0 ? 1 : -1;
		let newScale = scale.value * Math.exp(wheel * props.zoomSpeed);
		
		newScale = Math.max(props.minZoom, Math.min(newScale, props.maxZoom));

		const ratio = newScale / scale.value;

		boardX.value = mouseX - (mouseX - boardX.value) * ratio;
		boardY.value = mouseY - (mouseY - boardY.value) * ratio;
		scale.value = newScale;
	};

	const onMouseMove = (event: MouseEvent) => {
		if (!isDragging.value) return;
		boardX.value += event.movementX;
		boardY.value += event.movementY;
	};

	const onMouseUp = () => {
		isDragging.value = false;
		window.removeEventListener('mousemove', onMouseMove);
		window.removeEventListener('mouseup', onMouseUp);
	};

	const onMouseDown = () => {
		isDragging.value = true;
		window.addEventListener('mousemove', onMouseMove);
		window.addEventListener('mouseup', onMouseUp);
	};

	onUnmounted(() => {
		window.removeEventListener('mousemove', onMouseMove);
		window.removeEventListener('mouseup', onMouseUp);
	});
</script>

<template>
	<div
		ref="wrapperRef"
		class="ui-board-wrapper"
		:class="{ dragging: isDragging }"
		:style="{
			backgroundPosition: `${boardX}px ${boardY}px`,
			backgroundSize: `${props.gridSize*scale}px ${props.gridSize*scale}px`
		}"
		@mousedown.left.prevent="onMouseDown"
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