<script setup lang="ts">
	// Copyright (C) 2026 MorangTong Creative Studio
	// SPDX-License-Identifier: AGPL-3.0-or-later

	import { ref, watch, onMounted, provide } from 'vue';

	interface Section {
		id: string;
		title: string;
	}

	const sections = ref<Section[]>([]);
	const bookmarks = ref<string[]>([]);

	onMounted(() => {
		const saved = localStorage.getItem('debug-bookmarks');
		if (saved) {
			try {
				bookmarks.value = JSON.parse(saved);
			} catch (e) {
				console.error('Failed to parse bookmarks', e);
			}
		}
	});

	watch(bookmarks, (newVal) => {
			localStorage.setItem('debug-bookmarks', JSON.stringify(newVal));
		}, 
		{ 
			deep: true 
		}
	);

	const registerSection = (id: string, title: string) => {
		if (!sections.value.find(s => s.id === id)) {
			sections.value.push({ id, title });
		}
	};

	const unregisterSection = (id: string) => {
		sections.value = sections.value.filter(s => s.id !== id);
	};

	const toggleBookmark = (id: string) => {
		const index = bookmarks.value.indexOf(id);
		if (index === -1) {
			bookmarks.value.push(id);
		} else {
			bookmarks.value.splice(index, 1);
		}
	};

	provide('debugShowcase', {
		bookmarks,
		registerSection,
		unregisterSection,
		toggleBookmark
	});

	const scrollToSection = (id: string) => {
		const el = document.getElementById(`debug-section-${id}`);
		if (el) {
			el.scrollIntoView({ behavior: 'smooth', block: 'center', inline: 'center' });
			el.classList.add('pulse-highlight');
			setTimeout(() => el.classList.remove('pulse-highlight'), 1000);
		}
	};
</script>

<template>
    <div class="debug-showcase">
        <div class="toc-card">
            <div class="toc-column toc-all">
                <h3 class="toc-title">All Components</h3>
                <ul class="toc-list">
                    <li v-for="section in sections" :key="section.id">
                        <a href="#" @click.prevent="scrollToSection(section.id)" class="toc-link">
                            {{ section.title }}
                        </a>
                    </li>
                </ul>
            </div>

            <div class="toc-separator"></div>

            <div class="toc-column toc-bookmarks">
                <h3 class="toc-title">Bookmarks</h3>
                <ul class="toc-list">
                    <li v-for="id in bookmarks" :key="`bm-${id}`">
                        <a 
                            href="#" 
                            @click.prevent="scrollToSection(id)" 
                            class="toc-link bookmark-link"
                        >
                            {{ sections.find(s => s.id === id)?.title || 'Unknown section' }}
                        </a>
                    </li>
                    <li v-if="bookmarks.length === 0" class="toc-empty">
                        No bookmarks
                    </li>
                </ul>
            </div>
        </div>

        <div class="showcase-content">
            <slot />
        </div>
    </div>
</template>

<style scoped>
    .debug-showcase {
        display: flex;
        flex-direction: column;
        gap: 40px;
        padding: 80px;
    }

    .toc-card {
        display: flex;
        background-color: var(--color-surface);
        border: 1px solid var(--color-border);
        border-radius: 16px;
        padding: 24px;
        width: max-content;
        min-width: 500px;
        box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    }

    .toc-column {
        flex: 1;
        display: flex;
        flex-direction: column;
        gap: 12px;
    }

    .toc-separator {
        width: 1px;
        background-color: var(--color-border);
        margin: 0 24px;
    }

    .toc-title {
        margin: 0;
        font-size: var(--font-size-sm);
        text-transform: uppercase;
        letter-spacing: var(--letter-spacing-wide);
        color: var(--color-text-secondary); 
    }

    .toc-list {
        padding: 0;
        margin: 0;
        display: flex;
        flex-direction: column;
        gap: 8px;
    }

    .toc-link {
        color: var(--color-text-primary);
        text-decoration: none;
        font-size: var(--font-size-lg);
        cursor: pointer;
        display: inline-block;
    }

    .toc-link:hover {
        color: var(--color-accent);
        transform: translateX(4px);
    }

    .bookmark-link {
        color: var(--color-accent);
    }

    .toc-empty {
        color: var(--color-text-secondary);
        font-size: var(--font-size-base);
        font-style: italic;
        opacity: 0.7;
    }

    .showcase-content {
        display: flex;
        flex-direction: column;
        gap: 32px;
    }
</style>