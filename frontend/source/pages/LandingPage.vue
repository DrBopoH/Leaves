<script setup lang="ts"> // source/pages/LandingPage.vue
import { RouterLink } from 'vue-router';
</script>

<template>
	<div class="home-container">
		<div class="glow-orb top-left"></div>
		<div class="glow-orb bottom-right"></div>

		<div class="hero">
			<h1 class="title">Your Personal <span class="highlight">Space</span></h1>
			<p class="subtitle">
				Leaves is a lightweight, fast, and secure messenger designed for those who value privacy and simplicity.
			</p>
			<RouterLink to="/signup" class="cta-button">Start Chatting</RouterLink>
		</div>

		<div class="features">
			<div class="feature-card">
				<div class="icon">
					<svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="#88ffb4" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2"></polygon></svg>
				</div>
				<h3>Lightning Fast</h3>
				<p>Zero lag and instant message delivery, powered by our custom backend.</p>
			</div>
			
			<div class="feature-card">
				<div class="icon">
					<svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="#88ffb4" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect><path d="M7 11V7a5 5 0 0 1 10 0v4"></path></svg>
				</div>
				<h3>Secure by Design</h3>
				<p>Your data belongs to you. Absolute privacy in every single message.</p>
			</div>
			
			<div class="feature-card">
				<div class="icon">
					<svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="#88ffb4" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 2v20M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"></path></svg>
				</div>
				<h3>Clean & Minimal</h3>
				<p>No ads, no unnecessary clutter. Just you and your conversations.</p>
			</div>
		</div>
	</div>
</template>

<style scoped>
/* Анимации */
@keyframes fadeInUp {
	from { opacity: 0; transform: translateY(30px); }
	to { opacity: 1; transform: translateY(0); }
}

@keyframes float {
	0% { transform: translateY(0px); }
	50% { transform: translateY(-10px); }
	100% { transform: translateY(0px); }
}

.home-container {
	position: relative;
	width: 100%;
	max-width: 1000px;
	margin: 0 auto;
	padding: 40px 20px;
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content: center;
	text-align: center;
}

/* Фоновое свечение */
.glow-orb {
	position: absolute;
	width: 300px;
	height: 300px;
	background: radial-gradient(circle, rgba(136, 255, 180, 0.08) 0%, transparent 70%);
	border-radius: 50%;
	z-index: -1;
	pointer-events: none;
}

.top-left { top: -100px; left: -100px; }
.bottom-right { bottom: -100px; right: -100px; }

/* Герой-секция */
.hero {
	max-width: 700px;
	margin-bottom: 80px;
	animation: fadeInUp 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}

.title {
	font-size: 3.5rem;
	font-weight: 800;
	color: #ffffff;
	margin-bottom: 20px;
	letter-spacing: -1px;
	line-height: 1.2;
}

.highlight {
	color: #88ffb4;
	text-shadow: 0 0 20px rgba(136, 255, 180, 0.3);
}

.subtitle {
	font-size: 1.25rem;
	color: #a0a0a0;
	margin-bottom: 40px;
	line-height: 1.6;
}

/* Кнопка */
.cta-button {
	display: inline-block;
	text-decoration: none;
	background-color: transparent;
	color: #88ffb4;
	padding: 16px 36px;
	border-radius: 50px; /* Сделали кнопку более круглой, современной */
	border: 2px solid #88ffb4;
	font-size: 1.1rem;
	font-weight: 600;
	transition: all 0.3s ease;
	box-shadow: 0 0 15px rgba(136, 255, 180, 0.1);
}

.cta-button:hover {
	background-color: #88ffb4;
	color: #121212;
	box-shadow: 0 0 30px rgba(136, 255, 180, 0.4);
	transform: translateY(-3px);
}

/* Секция преимуществ */
.features {
	display: grid;
	grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
	gap: 30px;
	width: 100%;
	animation: fadeInUp 0.8s cubic-bezier(0.16, 1, 0.3, 1) 0.2s forwards;
	opacity: 0; /* Для анимации с задержкой */
}

.feature-card {
	background: #1a1a1a;
	border: 1px solid rgba(136, 255, 180, 0.05);
	padding: 30px;
	border-radius: 16px;
	text-align: left;
	transition: all 0.3s ease;
}

.feature-card:hover {
	border-color: rgba(136, 255, 180, 0.3);
	transform: translateY(-5px);
	box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5);
}

.feature-card .icon {
	font-size: 2.5rem;
	margin-bottom: 20px;
	animation: float 4s ease-in-out infinite;
}

.feature-card h3 {
	color: #ffffff;
	font-size: 1.3rem;
	margin-bottom: 15px;
}

.feature-card p {
	color: #888888;
	line-height: 1.5;
	font-size: 0.95rem;
}

/* Адаптивность для телефонов */
@media (max-width: 768px) {
	.title { font-size: 2.5rem; }
	.subtitle { font-size: 1.1rem; }
	.hero { margin-bottom: 50px; }
}
</style>