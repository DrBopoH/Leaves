// source/api/auth.ts
const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

export interface AuthPayload {
    username?: string; 
    email: string;
    password: string;
    rememberMe?: boolean;
}

export const signupUser = async (data: AuthPayload) => {
    const response = await fetch(`${BASE_URL}/signup`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data),
        credentials: 'include',
    });

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({}));
        throw new Error(errorData.message || 'Error during signup');
    }
    return response.json();
};

export const signinUser = async (data: AuthPayload) => {
    const response = await fetch(`${BASE_URL}/signin`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data),
        credentials: 'include',
    });

    if (!response.ok) {
        throw new Error('Invalid email or password'); 
    }
    return response.json(); 
};

export const fetchMe = async () => {
    const response = await fetch(`${BASE_URL}/me`, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
    });

    if (!response.ok) return null;
    
    const data = await response.json();
    return data.user;
};