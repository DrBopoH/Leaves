<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/ui-kit/UiInput.vue

const props = defineProps<{
	modelValue: string;
	type?: string;
	placeholder?: string;
	disabled?: boolean;
}>();

const emit = defineEmits<{
	(e: 'update:modelValue', value: string): void;
}>();
</script>

<template>
	<div class="ui-input-wrapper">
		<div v-if="$slots.prepend" class="ui-icon-prepend">
			<slot name="prepend"></slot>
		</div>

		<input
			:type="type || 'text'"
			:value="modelValue"
			@input="emit('update:modelValue', ($event.target as HTMLInputElement).value)"
			:placeholder="placeholder"
			:disabled="disabled"
			class="ui-input"
			:class="{ 'has-prepend': $slots.prepend, 'has-append': $slots.append }"
		/>

		<div v-if="$slots.append" class="ui-icon-append">
			<slot name="append"></slot>
		</div>
	</div>
</template>

<style scoped>
.ui-input-wrapper {
	position: relative;
	display: flex;
	align-items: center;
	width: 100%;
}

.ui-input {
	width: 100%;
	background-color: var(--color-bg);
	border: 1px solid var(--color-border);
	border-radius: 12px;
	padding: 12px 16px;
	color: var(--color-text-primary);
	font-size: 14px;
	outline: none;
	transition: border-color 0.2s, box-shadow 0.2s, background-color 0.2s;
}

.ui-input::placeholder {
	color: var(--color-text-secondary);
}

.ui-input:focus {
	background-color: var(--color-surface);
	border-color: var(--color-accent);
	box-shadow: 0 0 0 2px var(--color-accent-shadow);
}

.ui-input:disabled {
	opacity: 0.5;
	cursor: not-allowed;
}

.ui-input.has-prepend {
	padding-left: 40px;
}

.ui-input.has-append {
	padding-right: 40px;
}

.ui-icon-prepend,
.ui-icon-append {
	position: absolute;
	color: var(--color-text-secondary);
	display: flex;
	align-items: center;
	justify-content: center;
}

.ui-icon-prepend {
	left: 12px;
}

.ui-icon-append {
	right: 12px;
}

.ui-input:-webkit-autofill {
	box-shadow: 0 0 0 100px var(--color-bg) inset;
	-webkit-text-fill-color: var(--color-text-primary);
}

.ui-input:-webkit-autofill:focus {
	box-shadow: 0 0 0 100px var(--color-surface) inset, 0 0 0 2px var(--color-accent-shadow);
}
</style>
