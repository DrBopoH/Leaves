// source/composables/useTheme.ts
import { ref, watch, computed } from 'vue';

type Theme = 'black' | 'light';

const currentTheme = ref<Theme>((localStorage.getItem('theme') as Theme) || 'black');

export function useTheme() {
    // Вычисляемое свойство: является ли текущая тема черной
    const isDark = computed(() => currentTheme.value === 'black');

    // Функция переключения темы
    const toggleTheme = () => {
        currentTheme.value = isDark.value ? 'light' : 'black';
    };

    watch(currentTheme, (newTheme) => {
        if (newTheme === 'black') {
            document.documentElement.removeAttribute('data-theme');
        } else {
            document.documentElement.setAttribute('data-theme', newTheme);
        }
        localStorage.setItem('theme', newTheme);
    }, { immediate: true }); 

    return {
        isDark,
        toggleTheme
    };
}