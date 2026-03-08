// frontend/src/api/auth.ts

const BASE_URL = import.meta.env.VITE_API_URL;

export interface AuthPayload {
    email: string;
    password: string;
}

export const signupUser = async (data: AuthPayload) => {
    const response = await fetch(`${BASE_URL}/signup`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data),
    });

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({}));
        throw new Error(errorData.message || 'Ошибка при регистрации');
    }
    
    return response.json();
};

export const signinUser = async (data: AuthPayload) => {
    const response = await fetch(`${BASE_URL}/signin`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data),
    });

    if (!response.ok) {
        throw new Error('Неверный email или пароль'); 
    }
    
    return response.json(); 
};