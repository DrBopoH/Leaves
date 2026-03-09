// source/stores/user.ts
import { ref } from 'vue';
import { defineStore } from 'pinia';

export const useUserStore = defineStore('user', () => {
    const currentUser = ref<{ id: number; username: string } | null>(null);

    function setUser(user: { id: number; username: string } | null) {
        currentUser.value = user;
    }

    return { currentUser, setUser };
});