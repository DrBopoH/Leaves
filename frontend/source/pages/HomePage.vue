<script setup lang="ts">
// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/pages/HomePage.vue
import { useRouter, RouterLink } from 'vue-router';
import { useTheme } from '../composables/useTheme';
import { useUserStore } from '../stores/user';
import UiLogo from '../components/ui-kit/UiLogo.vue';
import UiButton from '../components/ui-kit/UiButton.vue';
import UiUserItem from '../components/ui-kit/UiUserItem.vue';
import UiIconButton from '../components/ui-kit/UiIconButton.vue';
import UiCard from '../components/ui-kit/UiCard.vue';
import UiGlowBackground from '../components/ui-kit/UiGlowBackground.vue';

const router = useRouter();
const userStore = useUserStore();
const { isDark, toggleTheme } = useTheme();
</script>

<template>
	<div class="landing-page">
		<UiGlowBackground />

		<!-- Header specifically for landing -->
		<header class="landing-navbar">
			<RouterLink to="/" class="logo-link">
				<UiLogo size="medium" text="Leaves" />
			</RouterLink>

			<nav class="nav-links">
				<UiIconButton
					@click="toggleTheme"
					:title="isDark ? 'Switch to Light Theme' : 'Switch to Dark Theme'"
					class="theme-btn"
				>
					<svg v-if="isDark" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
						<path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path>
					</svg>
					<svg v-else viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
						<circle cx="12" cy="12" r="5"></circle>
						<line x1="12" y1="1" x2="12" y2="3"></line>
						<line x1="12" y1="21" x2="12" y2="23"></line>
						<line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line>
						<line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line>
						<line x1="1" y1="12" x2="3" y2="12"></line>
						<line x1="21" y1="12" x2="23" y2="12"></line>
						<line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line>
						<line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line>
					</svg>
				</UiIconButton>

				<template v-if="!userStore.currentUser">
					<UiButton variant="outline" @click="router.push('/auth')">
						Sign in
					</UiButton>
				</template>

				<template v-else>
					<UiUserItem
						:username="userStore.currentUser.username"
						:color="userStore.getUserColor(userStore.currentUser.username)"
						interactive
						@click="router.push('/app')"
					/>
				</template>
			</nav>
		</header>

		<div class="home-container">
			<div class="hero">
				<h1 class="title">Your Personal <span class="highlight">Space</span></h1>
				<p class="subtitle">
					Leaves is a lightweight, fast, and secure messenger designed for those who value privacy and simplicity.
				</p>
				<UiButton variant="primary" size="large" @click="router.push('/auth')">
					Start Chatting
				</UiButton>
			</div>

			<div class="features">
				<UiCard
					title="Lightning Fast"
					description="Zero lag and instant message delivery, powered by our custom backend."
					interactive
				>
					<template #icon>
						<svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
							<polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2"></polygon>
						</svg>
					</template>
				</UiCard>

				<UiCard
					title="Secure by Design"
					description="Your data belongs to you. Absolute privacy in every single message."
					interactive
				>
					<template #icon>
						<svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
							<rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
							<path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
						</svg>
					</template>
				</UiCard>

				<UiCard
					title="Clean & Minimal"
					description="No ads, no unnecessary clutter. Just you and your conversations."
					interactive
				>
					<template #icon>
						<svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
							<path d="M12 2v20M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"></path>
						</svg>
					</template>
				</UiCard>
			</div>
		</div>
	</div>
</template>

<style scoped>
@keyframes fadeInUp {
	from { opacity: 0; transform: translateY(30px); }
	to { opacity: 1; transform: translateY(0); }
}

@keyframes float {
	0% { transform: translateY(0px); }
	50% { transform: translateY(-10px); }
	100% { transform: translateY(0px); }
}

.landing-page {
	position: relative;
	width: 100%;
	display: flex;
	flex-direction: column;
	min-height: 100vh;
}

/* Landing Navbar */
.landing-navbar {
	position: fixed;
	top: 0;
	left: 0;
	right: 0;
	width: 100%;
	z-index: 100;
	display: flex;
	justify-content: space-between;
	align-items: center;
	padding: 1rem 3rem;
	background: var(--color-overlay);
	backdrop-filter: blur(12px);
	border-bottom: 1px solid var(--color-border);
	transition: background-color 0.3s ease, border-color 0.3s ease;
}

.logo-link {
	text-decoration: none;
	display: inline-flex;
}

.nav-links {
	display: flex;
	gap: 15px;
	align-items: center;
}

.theme-btn {
	margin-right: 10px;
}

/* Home Container */
.home-container {
	position: relative;
	width: 100%;
	max-width: 1000px;
	margin: 0 auto;
	padding: 100px 20px 40px 20px;
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content: center;
	text-align: center;
	flex: 1;
}

/* Геройский блок */
.hero {
	position: relative;
	z-index: 1;
	max-width: 700px;
	margin-bottom: 80px;
	animation: fadeInUp 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}

.title {
	font-size: 3.5rem;
	font-weight: 800;
	color: var(--color-text-primary);
	margin-bottom: 20px;
	letter-spacing: -1px;
	line-height: 1.2;
	transition: color 0.3s ease;
}

.highlight {
	color: var(--color-accent);
	text-shadow: 0 0 20px var(--color-accent-shadow);
	transition: color 0.3s ease, text-shadow 0.3s ease;
}

.subtitle {
	font-size: 1.25rem;
	color: var(--color-text-secondary);
	margin-bottom: 40px;
	line-height: 1.6;
	transition: color 0.3s ease;
}

/* Блок преимуществ */
.features {
	position: relative;
	z-index: 1;
	display: grid;
	grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
	gap: 30px;
	width: 100%;
	animation: fadeInUp 0.8s cubic-bezier(0.16, 1, 0.3, 1) 0.2s forwards;
	opacity: 0;
}

/* Медиа-запрос */
@media (max-width: 768px) {
	.landing-navbar {
		padding: 1rem 1.5rem;
	}
	.title { font-size: 2.5rem; }
	.subtitle { font-size: 1.1rem; }
	.hero { margin-bottom: 50px; }
}
</style>
