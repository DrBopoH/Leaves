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
        background-color: var(--color-surface);
        overflow: hidden;
    }

    .section-header {
        display: flex;
        justify-content: space-between;
        background-color: var(--color-bg); 
    }

    .section-title {
        color: var(--color-text-primary);
        font-weight: var(--font-weight-medium);
    }

    .bookmark-btn {
        background: none;
        border: none;
        color: var(--color-text-secondary);
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .bookmark-btn svg {
        width: 18px;
        height: 18px;
    }

    .bookmark-btn:hover {
        color: var(--color-text-primary);
        background-color: var(--color-surface-hover);
    }

    .bookmark-btn.active {
        color: var(--color-accent); 
    }
</style>