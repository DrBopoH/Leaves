// Copyright (C) 2026 MorangTong Creative Studio
// SPDX-License-Identifier: AGPL-3.0-or-later

// source/api/chat.ts
const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

const defaultHeaders = {
	'Content-Type': 'application/json',
	'ngrok-skip-browser-warning': 'true'
};

export const fetchChatHistory = async () => {
	const response = await fetch(`${BASE_URL}/messages`, {
		method: 'GET',
		headers: defaultHeaders,
		credentials: 'include',
	});

	if (!response.ok) throw new Error('Failed to fetch history');
	return response.json();
};

export const fetchUsers = async () => {
	const response = await fetch(`${BASE_URL}/users`, {
		method: 'GET',
		headers: defaultHeaders,
		credentials: 'include',
	});

	if (!response.ok) throw new Error('Failed to fetch users');
	return response.json();
};
