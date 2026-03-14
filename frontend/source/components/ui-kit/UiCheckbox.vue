<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiCheckbox.vue

const props = defineProps<{
	modelValue: boolean;
	label?: string;
	disabled?: boolean;
}>();

const emit = defineEmits<{
	(e: 'update:modelValue', value: boolean): void;
}>();
</script>

<template>
	<label class="ui-checkbox" :class="{ 'is-disabled': disabled }">
		<input
			type="checkbox"
			:checked="modelValue"
			@change="emit('update:modelValue', ($event.target as HTMLInputElement).checked)"
			:disabled="disabled"
			class="hidden-input"
		/>

		<div class="box">
			<svg class="check-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round">
				<polyline points="20 6 9 17 4 12"></polyline>
			</svg>
		</div>

		<span v-if="label || $slots.default" class="label-text">
			<slot>{{ label }}</slot>
		</span>
	</label>
</template>

<style scoped>
.ui-checkbox {
	display: inline-flex;
	align-items: center;
	gap: 8px;
	cursor: pointer;
	user-select: none;
}

.ui-checkbox.is-disabled {
	opacity: 0.5;
	cursor: not-allowed;
}

.hidden-input {
	display: none;
}

.box {
	width: 18px;
	height: 18px;
	background-color: var(--color-bg);
	border: 1px solid var(--color-border);
	border-radius: 4px;
	display: flex;
	align-items: center;
	justify-content: center;
	transition: all 0.2s ease;
}

.check-icon {
	width: 12px;
	height: 12px;
	color: var(--color-text-inverse);
	opacity: 0;
	transform: scale(0.5);
	transition: all 0.2s cubic-bezier(0.175, 0.885, 0.32, 1.275);
}

.hidden-input:checked + .box {
	background-color: var(--color-accent);
	border-color: var(--color-accent);
}

.hidden-input:checked + .box .check-icon {
	opacity: 1;
	transform: scale(1);
}

.ui-checkbox:not(.is-disabled):hover .box {
	border-color: var(--color-accent);
}

.label-text {
	color: var(--color-text-primary);
	font-size: 14px;
}
</style>
