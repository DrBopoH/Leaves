import { createRouter, createWebHistory } from 'vue-router'

import IndexView from '../pages/LandingPage.vue'
import SignupView from '../pages/SignupPage.vue'
import SigninView from '../pages/SigninPage.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        { path: '/', name: 'index', component: IndexView },
        { path: '/signup', name: 'signup', component: SignupView },
        { path: '/signin', name: 'signin', component: SigninView },
    ],
})

export default router