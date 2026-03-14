// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/router/index.ts
import { createRouter, createWebHashHistory } from 'vue-router';
import AdminPage from '../pages/AdminPage.vue';
import HomePage from '../pages/HomePage.vue';
import AuthPage from '../pages/AuthPage.vue';
import AppPage from '../pages/AppPage.vue';

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{ path: '/', name: 'home', component: HomePage },
		{ path: '/auth', name: 'auth', component: AuthPage },
		{ path: '/app', name: 'app', component: AppPage },
		{ path: '/admin', name: 'admin', component: AdminPage },
	],
});

export default router;
