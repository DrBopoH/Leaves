import { createRouter, createWebHistory } from 'vue-router'

import IndexView from '../views/IndexView.vue'
import SignupView from '../views/SignupView.vue'
import SigninView from '../views/SigninView.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        { path: '/', name: 'index', component: IndexView },
        { path: '/signup', name: 'signup', component: SignupView },
        { path: '/signin', name: 'signin', component: SigninView },
    ],
})

export default router