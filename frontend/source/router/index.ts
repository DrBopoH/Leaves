// source/router/index.ts
import { createRouter, createWebHashHistory } from 'vue-router'

import LandingPage from '../pages/LandingPage.vue'
import AuthPage from '../pages/AuthPage.vue'
import AppPage from '../pages/AppPage.vue'

const router = createRouter(
    {
        history: createWebHashHistory(import.meta.env.BASE_URL),
        routes: [
            { path: '/',       name: 'landing', component: LandingPage },
            { path: '/auth', name: 'auth', component: AuthPage },
            { path: '/app',    name: 'app', component: AppPage },
        ],
    }
)

export default router