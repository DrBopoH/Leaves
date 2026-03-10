<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { signinUser, signupUser, fetchMe } from '../api/auth';
import { useUserStore } from '../stores/user';

const router = useRouter();
const userStore = useUserStore();

const isLogin = ref(true);

const username = ref('');
const email = ref('');
const password = ref('');
const rememberMe = ref(false);
const serverMessage = ref('');
const isLoading = ref(false);
const showPassword = ref(false);

const clientError = computed(() => {
    if (!email.value && !password.value && (isLogin.value || !username.value)) return '';    
    if (email.value && !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email.value)) return 'Invalid email format';
    if (password.value && password.value.length < 6) return 'Password is too short';
    return '';
});

const isFormValid = computed(() => {
    if (isLogin.value) {
        return email.value && password.value && !clientError.value;
    } else {
        return username.value && email.value && password.value && !clientError.value;
    }
});

const toggleMode = () => {
    isLogin.value = !isLogin.value;
    serverMessage.value = '';
};

const handleSubmit = async () => {
    if (!isFormValid.value) return;

    isLoading.value = true;
    serverMessage.value = '';

    try {
        if (isLogin.value) {
            const result = await signinUser({ email: email.value, password: password.value, rememberMe: rememberMe.value });
            serverMessage.value = result.message; 
            
            const user = await fetchMe();
            if (user) userStore.setUser(user);
            router.push('/app');
        
        } else {
            const result = await signupUser({ username: username.value, email: email.value, password: password.value });
            serverMessage.value = result.message; 

            setTimeout(() => { isLogin.value = true; }, 1500);
        }
    } catch (error: any) {
        serverMessage.value = error.message || 'Internal server error';
    } finally {
        isLoading.value = false;
    }
};
</script>

<template>
    <div class="auth-layout">
        <div class="glow-orb top-left"></div>
        <div class="glow-orb bottom-right"></div>
        
        <div class="brand-header">
            <div class="logo-wrapper">
                <img src="/Logo.svg" alt="Leaves" />
                <h1>Leaves</h1>
            </div>
            <p class="brand-subtitle">Create your account and start building your conversations today.</p>
        </div>

        <div class="auth-card">
            <transition name="fade" mode="out-in">
                <div class="card-headings" :key="isLogin ? 'login-head' : 'reg-head'">
                    <h2>{{ isLogin ? 'Welcome back!' : 'Get started' }}</h2>
                    <p>{{ isLogin ? 'Sign in to continue your conversations and stay connected.' : 'Sign up to connect with your communities and start messaging instantly.' }}</p>
                </div>
            </transition>

            <form @submit.prevent="handleSubmit" class="form-content">
                <div style="position: relative;">
                    <transition name="fade" mode="out-in">
                        <div :key="isLogin ? 'login-inputs' : 'reg-inputs'" class="inputs-container">
                            
                            <div v-if="!isLogin" class="input-box" style="margin-bottom: 28px;">
                                <input type="text" v-model="username" placeholder="username" required />
                            </div>

                            <div class="input-box" style="margin-bottom: 16px;">
                                <input type="email" v-model="email" placeholder="your@email.com" required />
                            </div>

                            <div class="input-box has-icon">
                                <input :type="showPassword ? 'text' : 'password'" v-model="password" placeholder="password" required />
                                <svg 
                                    @mousedown="showPassword = true" 
                                    @mouseup="showPassword = false" 
                                    @mouseleave="showPassword = false"
                                    @touchstart.prevent="showPassword = true"
                                    @touchend.prevent="showPassword = false"
                                    @touchcancel.prevent="showPassword = false"
                                    @contextmenu.prevent
                                    class="eye-icon" 
                                    viewBox="0 0 24 24" 
                                    fill="none" 
                                    stroke="currentColor" 
                                    stroke-width="2" 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round"
                                >
                                    <path v-if="!showPassword" d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24M1 1l22 22"/>
                                    <path v-else d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path>
                                    <circle v-if="showPassword" cx="12" cy="12" r="3"></circle>
                                </svg>
                            </div>
                            
                            <div v-if="isLogin" class="options-row">
                                <label class="remember-me">
                                    <input type="checkbox" v-model="rememberMe" />
                                    <span class="checkmark"></span>
                                    <span>remember me</span>
                                </label>
                                <a href="#" class="forgot-password">forgot password?</a>
                            </div>
                        </div>
                    </transition>
                </div>  

                <div class="error-reservoir" :class="{ 'has-error': clientError || serverMessage }">
                    <div class="error-content">
                        <span v-if="clientError" class="error-text error-client">{{ clientError }}</span>
                        <span v-else-if="serverMessage" class="error-text error-server">{{ serverMessage }}</span>
                    </div>
                </div>

                <button type="submit" class="submit-action" :class="{ 'is-loading': isLoading }" :disabled="!isFormValid || isLoading">
                    <span v-if="!isLoading">{{ isLogin ? 'Sign in' : 'Sign up' }}</span>
                    <svg v-else width="24" height="24" viewBox="0 0 24 24" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
                        <circle cx="4" cy="12" r="3"><animate id="a1" begin="0;a3.end-0.25s" attributeName="r" dur="0.75s" values="3;.2;3"/></circle>
                        <circle cx="12" cy="12" r="3"><animate begin="a1.begin+0.15s" attributeName="r" dur="0.75s" values="3;.2;3"/></circle>
                        <circle cx="20" cy="12" r="3"><animate id="a3" begin="a1.begin+0.3s" attributeName="r" dur="0.75s" values="3;.2;3"/></circle>
                    </svg>
                </button>
            </form>

            <div class="social-divider">
                <span>or continue with</span>
            </div>

            <div class="social-circles">
                <a href="#" class="s-circle">
                    <img src="https://api.iconify.design/logos:google-icon.svg" class="social-icon" alt="Google" />
                </a>
                <a href="#" class="s-circle">
                    <img src="https://api.iconify.design/skill-icons:github-dark.svg" class="social-icon" alt="GitHub" />
                </a>
                <a href="#" class="s-circle">
                    <img src="https://api.iconify.design/mdi:steam.svg?color=%2300adee" class="social-icon" alt="Steam" />
                </a>
                <a href="#" class="s-circle">
                    <img src="https://api.iconify.design/logos:discord-icon.svg" class="social-icon" alt="Discord" />
                </a>
            </div>

            <transition name="fade" mode="out-in">
                <div class="auth-footer" :key="isLogin ? 'login-foot' : 'reg-foot'">
                    <span v-if="isLogin">
                        Don't have an account? <a href="#" class="leaf-link" @click.prevent="toggleMode">Sign up</a>
                    </span>
                    <span v-else>
                        Already have an account? <a href="#" class="leaf-link" @click.prevent="toggleMode">Sign in</a>
                    </span>
                </div>
            </transition>
        </div>
    </div>
</template>

<style scoped>
.auth-layout {
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
    width: 100%;
    max-width: 440px;
    margin: 0 auto;
    padding: 40px 20px;
}

.glow-orb {
    position: absolute;
    width: 400px;
    height: 400px;
    background: radial-gradient(circle, #0f1714 0%, transparent 70%);
    border-radius: 50%;
    z-index: -1;
    pointer-events: none;
}
.top-left { top: -100px; left: -150px; }
.bottom-right { bottom: -100px; right: -150px; }

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
    color: #c8c2b8;
    letter-spacing: -0.5px;
}

.brand-subtitle {
    font-size: 13px;
    color: #8a867f;
    margin: 0;
    line-height: 1.5;
}

.auth-card {
    background-color: #080b0a;
    border: 1px solid #0f1714;
    border-radius: 20px;
    padding: 40px 32px;
    width: 100%;
    box-shadow: 0 20px 40px rgba(0, 0, 0, 0.5);
    transition: all 0.3s ease;
}

.card-headings {
    text-align: center;
    margin-bottom: 30px;
}

.card-headings h2 {
    font-size: 20px;
    font-weight: 600;
    margin: 0 0 8px 0;
    color: #c8c2b8;
}

.card-headings p {
    font-size: 13px;
    color: #8a867f;
    line-height: 1.5;
    margin: 0;
}

.form-content {
    display: flex;
    flex-direction: column;
    gap: 16px;
    margin-bottom: 24px;
}

.inputs-container {
    display: flex;
    flex-direction: column;
    width: 100%;
}

.input-box {
    position: relative;
    width: 100%;
}

.input-box input {
    width: 100%;
    background-color: #040605; 
    border: 1px solid #0f1714;
    border-radius: 16px;
    padding: 12px 16px;
    color: #c8c2b8;
    font-size: 14px;
    outline: none;
    transition: all 0.2s ease;
}

.input-box input::placeholder {
    color: #64615c;
}

.input-box input:focus {
    background-color: #050807;
    border-color: #5fca08;
    box-shadow: 0 0 0 2px rgba(95, 202, 8, 0.1);
}

.input-box.has-icon input {
    padding-right: 45px;
}

.eye-icon {
    position: absolute;
    right: 16px;
    top: 50%;
    transform: translateY(-50%);
    width: 20px;
    height: 20px;
    color: #64615c;
    cursor: pointer;
    transition: color 0.2s;
    user-select: none;
    -webkit-user-select: none;
    -webkit-touch-callout: none;
    -webkit-tap-highlight-color: transparent;
    touch-action: none;
}

.eye-icon:hover {
    color: #8a867f;
}

.options-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-size: 12px;
    padding: 0 10px;
    margin: 4px 0 6px 0;
}

.remember-me {
    display: flex;
    align-items: center;
    gap: 8px;
    color: #8a867f;
    cursor: pointer;
    transition: color 0.2s ease;
}

.remember-me:hover {
    color: #c8c2b8;
}

.remember-me input {
    display: none;
}

.checkmark {
    width: 14px;
    height: 14px;
    background-color: #040605;
    border: 1px solid #64615c;
    border-radius: 3px;
    display: inline-block;
    position: relative;
    transition: all 0.2s ease;
}

.remember-me input:checked + .checkmark {
    background-color: #5fca08;
    border-color: #5fca08;
}

.remember-me input:checked + .checkmark::after {
    content: '';
    position: absolute;
    left: 4px;
    top: 1px;
    width: 4px;
    height: 8px;
    border: solid #050807;
    border-width: 0 2px 2px 0;
    transform: rotate(45deg);
}

.remember-me:hover input:not(:checked) + .checkmark {
    border-color: #8a867f;
}

.forgot-password {
    color: #8a867f;
    text-decoration: none;
    transition: color 0.2s;
}

.forgot-password:hover {
    color: #c8c2b8;
}

.error-reservoir {
    display: grid;
    grid-template-rows: 0fr;
    transition: grid-template-rows 0.3s ease, margin 0.3s ease;
}

.error-reservoir.has-error {
    grid-template-rows: 1fr;
    margin-top: 4px;
    margin-bottom: 8px;
}

.error-content {
    overflow: hidden;
    display: flex;
    align-items: center;
    justify-content: center;
}

.error-text {
    font-size: 13px;
    font-weight: 500;
    text-align: center;
    margin: 0;
    padding: 2px 0;
}

.error-client {
    color: #8a867f;
}

.error-server {
    color: #ff5555;
}

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
    background-color: #0f1714;
    color: #64615c;
    cursor: not-allowed;
    box-shadow: none;
}

.submit-action:not(:disabled):not(.is-loading) {
    background-color: #5fca08;
    color: #050807;
    cursor: pointer;
    box-shadow: 0 4px 15px rgba(95, 202, 8, 0.15);
}

.submit-action:not(:disabled):not(.is-loading):hover {
    background-color: #4da806;
    box-shadow: 0 6px 20px rgba(95, 202, 8, 0.25);
    transform: translateY(-1px);
}

.submit-action.is-loading {
    background-color: #040605;
    color: #5fca08;
    border: 1px solid #5fca08;
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
    background-color: #0f1714;
}

.social-divider::before { left: 0; }
.social-divider::after { right: 0; }

.social-divider span {
    font-size: 12px;
    color: #64615c;
    padding: 0 10px;
    background-color: #080b0a;
}

.social-circles {
    display: flex;
    justify-content: center;
    gap: 16px;
    margin-bottom: 24px;
}

.s-circle {
    width: 44px;
    height: 44px;
    background-color: #040605;
    border: 1px solid #0f1714;
    border-radius: 32px;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    transition: all 0.2s;
    text-decoration: none;
}

.s-circle:hover {
    background-color: #0f1714;
    border-color: #64615c;
}

.social-icon {
    width: 24px;
    height: 24px;
    object-fit: contain;
}

.auth-footer {
    text-align: center;
    font-size: 13px;
    color: #8a867f;
}

.auth-footer .leaf-link {
    color: #c8c2b8;
    text-decoration: none;
    font-weight: 500;
    margin-left: 4px;
    transition: color 0.2s;
}

.auth-footer .leaf-link:hover {
    color: #5fca08;
    text-decoration: underline;
}

.fade-enter-active, .fade-leave-active { transition: opacity 0.2s ease, transform 0.2s ease; }
.fade-enter-from { opacity: 0; transform: translateY(-5px); }
.fade-leave-to { opacity: 0; transform: translateY(5px); }

.fade-form-enter-active, .fade-form-leave-active { transition: opacity 0.2s ease, transform 0.2s ease; }
.fade-form-leave-active { position: absolute; top: 0; left: 0; width: 100%; }
.fade-form-enter-from { opacity: 0; transform: translateY(-10px); }
.fade-form-leave-to { opacity: 0; transform: translateY(10px); }
</style>