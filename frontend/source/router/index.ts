// source/router/index.ts
import { createRouter, createWebHistory } from 'vue-router'

import LandingPage from '../pages/LandingPage.vue'
import SignupPage from '../pages/SignupPage.vue'
import SigninPage from '../pages/SigninPage.vue'
import AppPage from '../pages/AppPage.vue'

const router = createRouter(
    {
        history: createWebHistory(import.meta.env.BASE_URL),
        routes: [
            { path: '/',       name: 'landing', component: LandingPage },
            { path: '/signup', name: 'signup', component: SignupPage },
            { path: '/signin', name: 'signin', component: SigninPage },
            { path: '/app',    name: 'app', component: AppPage },
        ],
    }
)

export default router