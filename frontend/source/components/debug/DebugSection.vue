<script setup lang="ts">
	// Copyright (C) 2026 MorangTong Creative Studio
	// SPDX-License-Identifier: AGPL-3.0-or-later

	import { computed, inject, onMounted, onUnmounted } from 'vue';

	const props = defineProps<{
		id: string;
		title: string;
	}>();

	const showcase = inject('debugShowcase') as any;

	onMounted(() => {
		showcase?.registerSection(props.id, props.title);
	});

	onUnmounted(() => {
		showcase?.unregisterSection(props.id);
	});

	const isBookmarked = computed(() => showcase?.bookmarks.value.includes(props.id));

	const handleBookmark = () => {
		showcase?.toggleBookmark(props.id);
	};
</script>

<template>
    <div :id="`debug-section-${id}`" class="debug-section">
        <div class="section-header">
            <h4 class="section-title">{{ title }}</h4>
            <button 
                class="bookmark-btn" 
                :class="{ active: isBookmarked }"
                @click="handleBookmark"
                :title="isBookmarked ? 'Remove bookmark' : 'Add bookmark'"
            >
                <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
                </svg>
            </button>
        </div>
        <div class="section-body">
            <slot />
        </div>
    </div>
</template>
<style scoped>
	.debug-section {
		background: rgba(100, 100, 100, 0.1);
		border: 1px solid rgba(255, 255, 255, 0.05);
		border-radius: 12px;
		overflow: hidden;
		transition: background-color 0.3s ease;
	}

	.section-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 12px 20px;
		background: rgba(0, 0, 0, 0.2);
		border-bottom: 1px solid rgba(255, 255, 255, 0.05);
	}

	.section-title {
		margin: 0;
		color: #fff;
		font-weight: 500;
	}

	.bookmark-btn {
		background: none;
		border: none;
		color: rgba(255, 255, 255, 0.3);
		cursor: pointer;
		padding: 4px;
		border-radius: 4px;
		transition: all 0.2s ease;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.bookmark-btn svg {
		width: 18px;
		height: 18px;
	}

	.bookmark-btn:hover {
		color: #fff;
		background: rgba(255, 255, 255, 0.1);
	}

	.bookmark-btn.active {
		color: var(--color-accent, #ffb86c);
		fill: var(--color-accent, #ffb86c);
	}

	.section-body {
		padding: 20px;
	}
</style>