<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiSidebar.vue

const props = defineProps<{
	isOpen: boolean;
	position?: 'left' | 'right';
}>();
</script>

<template>
	<div
		class="ui-sidebar"
		:class="[
			`position-${position || 'left'}`,
			{ 'is-open': isOpen }
		]"
	>
		<div class="sidebar-content">
			<slot></slot>
		</div>

		<div v-if="$slots.footer" class="sidebar-footer">
			<slot name="footer"></slot>
		</div>
	</div>
</template>

<style scoped>
.ui-sidebar {
	background-color: var(--color-surface);
	border-color: var(--color-border);
	display: flex;
	flex-direction: column;
	flex-shrink: 0;
	transition-property: margin-left, margin-right, background-color, border-color, color, transform;
	transition-duration: 0.3s;
	transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1), ease, ease, ease, ease;
	position: relative;
	z-index: 100;
}

.position-left {
	border-right-width: 1px;
	border-right-style: solid;
	width: 260px;
}

.position-left:not(.is-open) {
	margin-left: -261px;
}

.position-right {
	border-left-width: 1px;
	border-left-style: solid;
	width: 240px;
}

.position-right:not(.is-open) {
	margin-right: -241px;
}

.sidebar-content {
	flex: 1;
	overflow-y: auto;
	display: flex;
	flex-direction: column;
}

.sidebar-footer {
	border-top: 1px solid var(--color-border);
	background-color: transparent;
	flex-shrink: 0;
}

/* Mobile overrides assuming we pass global classes or handle inside parent,
   but it's good to provide a generic mobile approach if desired.
   However, atomic ui-sidebar shouldn't dictate breakpoint layout too strictly,
   so we keep it basic and rely on parent container logic for mobile fixed positioning,
   or add a specific .is-mobile class later if needed. */
</style>
