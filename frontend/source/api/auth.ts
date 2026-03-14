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
	try {
		const response = await fetch(`${BASE_URL}/signup`, {
			method: 'POST',
			headers: defaultHeaders,
			body: JSON.stringify(data),
			credentials: 'include',
		});

		if (!response.ok) {
			const errorData = await response.json().catch(() => ({}));
			if (response.status === 409) {
				throw new Error('Email or username is already taken.');
			}
			if (response.status === 400) {
				throw new Error('Invalid data provided. Please check your inputs.');
			}
			throw new Error(errorData.message || 'An unexpected error occurred during signup.');
		}
		return await response.json();
	} catch (error: any) {
		if (error.name === 'TypeError') {
			throw new Error('Network error. Please check your internet connection.');
		}
		throw error;
	}
};

export const signinUser = async (data: AuthPayload) => {
	try {
		const response = await fetch(`${BASE_URL}/signin`, {
			method: 'POST',
			headers: defaultHeaders,
			body: JSON.stringify(data),
			credentials: 'include',
		});

		if (!response.ok) {
			if (response.status === 401) {
				throw new Error('Incorrect email or password.');
			}
			if (response.status === 429) {
				throw new Error('Too many login attempts. Please try again later.');
			}
			throw new Error('Server error occurred during login.');
		}
		return await response.json();
	} catch (error: any) {
		if (error.name === 'TypeError') {
			throw new Error('Network error. Please check your internet connection.');
		}
		throw error;
	}
};

export const fetchMe = async () => {
	try {
		const response = await fetch(`${BASE_URL}/me`, {
			method: 'GET',
			headers: defaultHeaders,
			credentials: 'include',
		});

		if (!response.ok) {
			return null;
		}

		const data = await response.json();
		return data.user;
	} catch (error) {
		console.warn('Failed to verify session. Assuming unauthenticated.');
		return null;
	}
};

