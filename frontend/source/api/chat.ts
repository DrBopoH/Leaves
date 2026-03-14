// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/api/chat.ts
const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

const defaultHeaders = {
	'Content-Type': 'application/json',
	'ngrok-skip-browser-warning': 'true'
};

export const fetchChatHistory = async () => {
	try {
		const response = await fetch(`${BASE_URL}/messages`, {
			method: 'GET',
			headers: defaultHeaders,
			credentials: 'include',
		});

		if (!response.ok) {
			if (response.status === 401) {
				throw new Error('Unauthorized. Please log in again to view chat history.');
			}
			throw new Error('Server failed to load chat messages.');
		}
		return await response.json();
	} catch (error: any) {
		if (error.name === 'TypeError') {
			throw new Error('Network error. Failed to load chat history.');
		}
		throw error;
	}
};

export const fetchUsers = async () => {
	try {
		const response = await fetch(`${BASE_URL}/users`, {
			method: 'GET',
			headers: defaultHeaders,
			credentials: 'include',
		});

		if (!response.ok) {
			throw new Error('Failed to retrieve active users list.');
		}
		return await response.json();
	} catch (error: any) {
		if (error.name === 'TypeError') {
			throw new Error('Network error while fetching users list.');
		}
		throw error;
	}
};
