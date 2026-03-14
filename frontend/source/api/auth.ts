// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/api/auth.ts
const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

export interface AuthPayload {
    username?: string;
    email: string;
    password: string;
    rememberMe?: boolean;
}

const defaultHeaders = {
    'Content-Type': 'application/json',
    'ngrok-skip-browser-warning': 'true'
};

export const signupUser = async (data: AuthPayload) => {
    const response = await fetch(`${BASE_URL}/signup`, {
        method: 'POST',
        headers: defaultHeaders,
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
        headers: defaultHeaders,
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
        headers: defaultHeaders,
        credentials: 'include',
    });

    if (!response.ok) return null;

    const data = await response.json();
    return data.user;
};

