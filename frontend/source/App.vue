<script setup lang="ts">
import { RouterLink, RouterView, useRoute } from 'vue-router';

const route = useRoute();
</script>

<template>
	<header class="navbar" v-if="!['/signin', '/signup', '/chats'].includes(route.path)">
		<div class="logo">
			<RouterLink to="/">
				<img src="/Logo.svg" alt="Leaves Logo" class="logo-icon" />
				Leaves
			</RouterLink>
		</div>
		<nav class="nav-links">
			<RouterLink to="/signin" class="nav-btn nav-btn-outline">Sign in</RouterLink>
		</nav>
	</header>

	<main class="content-wrapper">
		<RouterView v-slot="{ Component }">
			<transition name="discord-fade" mode="out-in">
      <component :is="Component" />
    </transition>
</RouterView>
	</main>
</template>

<style>
* {
	margin: 0;
	padding: 0;
	box-sizing: border-box;
	font-family: 'Inter', 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}
body {
	background-color: #0a0a0a;
	color: #ffffff;
	overflow-x: hidden;
}

/* --- ОБЩИЕ СТИЛИ ДЛЯ АВТОРИЗАЦИИ (Signup & Signin) --- */
.auth-layout {
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content: flex-start;
	width: 100%;
	max-width: 440px;
	margin: 0 auto;
	padding: 40px 20px;
}

.brand-header {
	text-align: center;
	margin-bottom: 30px;
}

.logo-wrapper {
	display: flex;
	align-items: center;
	justify-content: center;
	gap: 12px;
	margin-bottom: 12px;
}

.logo-wrapper img {
	width: 32px;
	height: 32px;
}

.logo-wrapper h1 {
	font-size: 28px;
	font-weight: 700;
	margin: 0;
	color: #fff;
	letter-spacing: -0.5px;
}

.brand-subtitle {
	font-size: 13px;
	color: #888;
	margin: 0;
	line-height: 1.5;
}

.auth-card {
	background-color: #1a1a1a;
	border: 1px solid #262626;
	border-radius: 20px;
	padding: 40px 32px;
	width: 100%;
	box-shadow: 0 20px 40px rgba(0, 0, 0, 0.5);
}

.card-headings {
	text-align: center;
	margin-bottom: 30px;
}

.card-headings h2 {
	font-size: 20px;
	font-weight: 600;
	margin: 0 0 8px 0;
	color: #fff;
}

.card-headings p {
	font-size: 13px;
	color: #888;
	line-height: 1.5;
	margin: 0;
}

.form-content {
	display: flex;
	flex-direction: column;
	gap: 16px;
	margin-bottom: 24px;
}

.input-box {
	position: relative;
	width: 100%;
}

.input-box input {
	width: 100%;
	background-color: #2a2a2a; 
	border: 1px solid #333;
	border-radius: 16px;
	padding: 14px 16px;
	color: #fff;
	font-size: 13px;
	outline: none;
	transition: all 0.2s ease;
}

.input-box input::placeholder { color: #555; }
.input-box input:focus {
	background-color: #222;
	border-color: #88ffb4;
	box-shadow: 0 0 0 2px rgba(136, 255, 180, 0.1);
}

.input-box.has-icon input { padding-right: 45px; }

.eye-icon {
	position: absolute;
	right: 16px;
	top: 50%;
	transform: translateY(-50%);
	width: 20px;
	height: 20px;
	color: #666;
	cursor: pointer;
	transition: color 0.2s;
	user-select: none;
}
.eye-icon:hover { color: #aaa; }

.error-reservoir {
	min-height: 20px; 
	display: flex;
	align-items: center;
	justify-content: center;
	text-align: center;
}
.error-text { font-size: 13px; font-weight: 500; }
.error-client { color: #888; }
.error-server { color: #ff5555; }

.submit-action {
	width: 100%;
	border: none;
	border-radius: 30px;
	padding: 14px;
	font-size: 15px;
	font-weight: 600;
	transition: all 0.3s ease;
	display: flex;
	justify-content: center;
	align-items: center;
}

.submit-action:disabled {
	background-color: #2a2a2a;
	color: #555;
	cursor: not-allowed;
	box-shadow: none;
}

.submit-action:not(:disabled):not(.is-loading) {
	background-color: #88ffb4;
	color: #000;
	cursor: pointer;
	box-shadow: 0 4px 15px rgba(136, 255, 180, 0.15);
}
.submit-action:not(:disabled):not(.is-loading):hover {
	background-color: #6ce09b;
	box-shadow: 0 6px 20px rgba(136, 255, 180, 0.25);
	transform: translateY(-1px);
}

.submit-action.is-loading {
	background-color: #000;
	color: #88ffb4;
	border: 1px solid #88ffb4;
	cursor: wait;
}

.social-divider {
	text-align: center;
	margin-bottom: 16px;
	position: relative;
}
.social-divider::before, .social-divider::after {
	content: '';
	position: absolute;
	top: 50%;
	width: 25%;
	height: 1px;
	background-color: #333;
}
.social-divider::before { left: 0; }
.social-divider::after { right: 0; }
.social-divider span { font-size: 12px; color: #666; padding: 0 10px; }

/* Выравниваем всю группу */
.social-circles {
    display: flex;
    justify-content: center;
    gap: 24px; /* Расстояние между иконками */
    margin-top: 16px;
    margin-bottom: 24px;
}

/* Убрали border и background, оставили только кликабельную зону */
.s-circle {
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    transition: transform 0.2s ease;
    text-decoration: none;
}

/* При наведении иконка будет чуть-чуть подпрыгивать */
.s-circle:hover {
    transform: translateY(-2px);
}

/* Настраиваем размер самих иконок, как на твоем оригинале */
.social-icon {
    width: 26px;
    height: 26px;
    object-fit: contain;
}

.google{background:white;border:1px solid #ddd;}
.github{background:#171515;}
.discord{background:#5865F2;}

.s-circle:hover{
transform:scale(1.1);
box-shadow:0 6px 15px rgba(0,0,0,0.2);
}1
.auth-footer {
	text-align: center;
	font-size: 13px;
	color: #888;
}
.auth-footer a { color: #fff; text-decoration: none; font-weight: 500; margin-left: 4px; }
.auth-footer a:hover { color: #88ffb4; text-decoration: underline; }
</style>

<style scoped>
/* Стили только для хедера */

.navbar {
	position: sticky;
	top: 0;
	z-index: 100;
	display: flex;
	justify-content: space-between;
	align-items: center;
	padding: 1rem 3rem;
	background: rgba(10, 10, 10, 0.7);
	backdrop-filter: blur(12px);
	border-bottom: 1px solid rgba(255, 255, 255, 0.05);
}
.logo a {
	display: flex;
	align-items: center;
	gap: 10px;
	font-size: 1.4rem;
	font-weight: 700;
	color: #fff;
	text-decoration: none;
	letter-spacing: -0.5px;
}
.logo-icon { width: 28px; height: 28px; }
.nav-links { display: flex; gap: 15px; }
.nav-btn {
	text-decoration: none;
	color: #000;
	background-color: #88ffb4;
	padding: 8px 20px;
	border-radius: 8px;
	font-weight: 600;
	font-size: 14px;
	transition: all 0.2s ease;
}
.nav-btn:hover { background-color: #6ce09b; }
.nav-btn-outline { background-color: transparent; border: 1px solid #333; color: #fff; }
.nav-btn-outline:hover { background-color: #1a1a1a; border-color: #555; }
.content-wrapper { display: flex; justify-content: center; align-items: center; padding: 20px; }
.discord-fade-enter-active,
.discord-fade-leave-active {
  transition: all 0.2s ease-out; /* Быстро, как в дискорде */
}

.discord-fade-enter-from {
  opacity: 0;
  transform: translateY(-10px);
}

.discord-fade-leave-to {
  opacity: 0;
  transform: translateY(10px);
}
.checkbox{
	color: #3f3f46;
	background-color: #3f3f46;
}
</style>