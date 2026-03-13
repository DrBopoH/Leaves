<script setup lang="ts">
import { ref, computed } from 'vue';

const props = defineProps<{
    isLogin: boolean;
    isLoading: boolean;
    serverMessage: string;
}>();

const emit = defineEmits<{
    (e: 'submitLogin', payload: any): void;
    (e: 'submitSignup', payload: any): void;
    (e: 'clearError'): void; // Очистка серверной ошибки при вводе
}>();

const username = ref('');
const email = ref('');
const password = ref('');
const rememberMe = ref(false);
const showPassword = ref(false);

const clientError = computed(() => {
    if (!email.value && !password.value && (props.isLogin || !username.value)) return '';
    if (!props.isLogin && username.value && (username.value.length < 3 || username.value.length > 30)) return 'Username must be between 3 and 30 characters';
    if (email.value && !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email.value)) return 'Invalid email format';
    if (password.value && password.value.length < 8) return 'Password must be at least 8 characters';
    return '';
});

const isFormValid = computed(() => {
    if (props.isLogin) {
        return email.value && password.value && !clientError.value;
    } else {
        return username.value && email.value && password.value && !clientError.value;
    }
});

const handleSubmit = () => {
    if (!isFormValid.value) return;
    
    if (props.isLogin) {
        emit('submitLogin', { email: email.value, password: password.value, rememberMe: rememberMe.value });
    } else {
        emit('submitSignup', { username: username.value, email: email.value, password: password.value });
    }
};

const handleInput = () => {
    emit('clearError');
};
</script>

<template>
    <form @submit.prevent="handleSubmit" class="form-content" @input="handleInput">
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
                        <svg width="20" height="20" @mousedown="showPassword = true" @mouseup="showPassword = false" @mouseleave="showPassword = false" @touchstart.prevent="showPassword = true" @touchend.prevent="showPassword = false" @touchcancel.prevent="showPassword = false" @contextmenu.prevent class="eye-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
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
</template>

<style scoped>
.form-content { display: flex; flex-direction: column; gap: 16px; margin-bottom: 24px; }
.inputs-container { display: flex; flex-direction: column; width: 100%; }
.input-box { position: relative; width: 100%; }
.input-box input {
    width: 100%; background-color: var(--color-surface-alt, #040605);
    border: 1px solid var(--color-border, #0f1714); border-radius: 16px;
    padding: 12px 16px; color: var(--color-text-primary, #c8c2b8);
    font-size: 14px; outline: none; transition: all 0.2s ease;
}
.input-box input::placeholder { color: var(--color-text-secondary, #64615c); }
.input-box input:focus {
    background-color: var(--color-surface, #050807); border-color: #5fca08;
    box-shadow: 0 0 0 2px rgba(95, 202, 8, 0.1);
}
.input-box input:-webkit-autofill, .input-box input:-webkit-autofill:hover, .input-box input:-webkit-autofill:focus, .input-box input:-webkit-autofill:active {
    -webkit-box-shadow: 0 0 0 1000px var(--color-surface-alt, #040605) inset !important;
    -webkit-text-fill-color: var(--color-text-primary, #c8c2b8) !important;
    transition: background-color 5000s ease-in-out 0s; border: 1px solid var(--color-border, #0f1714);
}
.input-box input:-webkit-autofill:focus {
    -webkit-box-shadow: 0 0 0 1000px var(--color-surface, #050807) inset, 0 0 0 2px rgba(95, 202, 8, 0.1) !important;
    border-color: #5fca08 !important;
}
.input-box.has-icon input { padding-right: 45px; }
.eye-icon {
    position: absolute; right: 16px; top: 50%; transform: translateY(-50%);
    width: 20px; height: 20px; color: var(--color-text-secondary, #64615c);
    cursor: pointer; transition: color 0.2s;
}
.eye-icon:hover { color: var(--color-text-primary, #8a867f); }
.options-row { display: flex; justify-content: space-between; align-items: center; font-size: 12px; padding: 0 10px; margin: 4px 0 6px 0; }
.remember-me { display: flex; align-items: center; gap: 8px; color: var(--color-text-secondary, #8a867f); cursor: pointer; transition: color 0.2s ease; }
.remember-me:hover { color: var(--color-text-primary, #c8c2b8); }
.remember-me input { display: none; }
.checkmark {
    width: 14px; height: 14px; background-color: var(--color-surface-alt, #040605);
    border: 1px solid var(--color-text-secondary, #64615c); border-radius: 3px;
    display: inline-block; position: relative; transition: all 0.2s ease;
}
.remember-me input:checked + .checkmark { background-color: #5fca08; border-color: #5fca08; }
.remember-me input:checked + .checkmark::after {
    content: ''; position: absolute; left: 4px; top: 1px; width: 4px; height: 8px;
    border: solid #050807; border-width: 0 2px 2px 0; transform: rotate(45deg);
}
.forgot-password { color: var(--color-text-secondary, #8a867f); text-decoration: none; transition: color 0.2s; }
.forgot-password:hover { color: var(--color-text-primary, #c8c2b8); }
.error-reservoir { display: grid; grid-template-rows: 0fr; transition: grid-template-rows 0.3s ease, margin 0.3s ease; }
.error-reservoir.has-error { grid-template-rows: 1fr; margin-top: 4px; margin-bottom: 8px; }
.error-content { overflow: hidden; display: flex; align-items: center; justify-content: center; }
.error-text { font-size: 13px; font-weight: 500; text-align: center; margin: 0; padding: 2px 0; }
.error-client { color: var(--color-text-secondary, #8a867f); }
.error-server { color: #ff5555; }
.submit-action {
    width: 100%; border: none; border-radius: 30px; padding: 14px; font-size: 15px;
    font-weight: 600; transition: all 0.3s ease; display: flex; justify-content: center; align-items: center;
}
.submit-action:disabled { background-color: var(--color-border, #0f1714); color: var(--color-text-secondary, #64615c); cursor: not-allowed; box-shadow: none; }
.submit-action:not(:disabled):not(.is-loading) { background-color: #5fca08; color: #050807; cursor: pointer; box-shadow: 0 4px 15px rgba(95, 202, 8, 0.15); }
.submit-action:not(:disabled):not(.is-loading):hover { background-color: #4da806; box-shadow: 0 6px 20px rgba(95, 202, 8, 0.25); transform: translateY(-1px); }
.submit-action.is-loading { background-color: var(--color-surface-alt, #040605); color: #5fca08; border: 1px solid #5fca08; cursor: wait; }
.fade-enter-active, .fade-leave-active { transition: opacity 0.2s ease, transform 0.2s ease; }
.fade-enter-from { opacity: 0; transform: translateY(-5px); }
.fade-leave-to { opacity: 0; transform: translateY(5px); }
</style>