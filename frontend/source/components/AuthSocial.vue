<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/components/AuthSocial.vue

import UiDivider from './ui-kit/UiDivider.vue';
import UiSocialLink from './ui-kit/UiSocialLink.vue';
</script>

<template>
	<UiDivider text="or continue with" class="social-divider" />

	<div class="social-circles">
		<UiSocialLink href="#" iconSrc="https://api.iconify.design/logos:google-icon.svg" iconAlt="Google" />

		<UiSocialLink href="#">
			<div class="social-icon github-icon" title="GitHub"></div>
		</UiSocialLink>

		<UiSocialLink href="#" steam />

		<UiSocialLink href="#" iconSrc="https://api.iconify.design/logos:discord-icon.svg" iconAlt="Discord" />
	</div>
</template>

<style scoped>
.social-divider {
	margin-bottom: 16px;
}

.social-circles {
	display: flex;
	justify-content: center;
	gap: 16px;
	margin-bottom: 24px;
}

.social-icon {
	width: 24px;
	height: 24px;
	object-fit: contain;
}

.github-icon {
	background-color: var(--color-text-primary);
	-webkit-mask: url('https://api.iconify.design/mdi:github.svg') no-repeat center / contain;
	mask: url('https://api.iconify.design/mdi:github.svg') no-repeat center / contain;
	transition: background-color 0.3s ease;
}
</style>
