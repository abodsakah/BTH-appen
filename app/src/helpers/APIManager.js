import axios from 'axios';
import { API_URL } from './Constants';
import * as SecureStore from 'expo-secure-store';

async function sendGETRequest(url, headers) {
	try {
		return await axios.get(url, { headers });
	} catch (error) {
		console.assert(error);
		return null;
	}
}

async function sendPOSTRequest(url, data, headers) {
	try {
		return await axios.post(url, data, headers);
	} catch (error) {
		console.assert(error);
		return null;
	}
}

async function sendPUTRequest(url, data, headers) {
	try {
		return await axios.put(url, data, headers);
	} catch (error) {
		console.assert(error);
		return null;
	}
}

async function sendDELETERequest(url, headers) {
	try {
		return await axios.delete(url, headers);
	} catch (error) {
		console.assert(error);
		return null;
	}
}

export async function loginUser(username, password) {
	const url = `${API_URL}/login`;
	const data = {
		username,
		password,
	};
	return await sendPOSTRequest(url, data);
}

export async function listAllExams() {
	const url = `${API_URL}/list-exams`;
	return await sendGETRequest(url, {
		headers: {
			jwt: `${await JSON.parse(SecureStore.getItemAsync('user'))?.jwt}`,
		},
	});
}

export async function listDueExams() {
	const url = `${API_URL}/list-due-exams`;
	return await sendGETRequest(url, {
		headers: {
			jwt: `${await JSON.parse(SecureStore.getItemAsync('user'))?.jwt}`,
		},
	});
}

export async function listExam(examId) {
	const url = `${API_URL}/list-exam/${examId}`;
	return await sendGETRequest(url, {
		headers: {
			jwt: `${await JSON.parse(SecureStore.getItemAsync('user'))?.jwt}`,
		},
	});
}

export async function registerExam(exam_id) {
	const url = `${API_URL}/register-exam`;
	return await sendPOSTRequest(
		url,
		{ exam_id },
		{
			headers: {
				jwt: `${await JSON.parse(SecureStore.getItemAsync('user'))?.jwt}`,
			},
		}
	);
}

export async function unregisterExam(exam_id) {
	const url = `${API_URL}/unregister-exam/`;
	return await sendPOSTRequest(
		url,
		{ exam_id },
		{
			headers: {
				jwt: `${await JSON.parse(SecureStore.getItemAsync('user'))?.jwt}`,
			},
		}
	);
}

export async function deleteExams(exam_id) {
	const url = `${API_URL}/delete-exam/`;
	return await sendPOSTRequest(
		url,
		{ exam_id },
		{
			headers: {
				jwt: `${await JSON.parse(SecureStore.getItemAsync('user'))?.jwt}`,
			},
		}
	);
}

export async function addExpoPushToken(pushToken) {
	const url = `${API_URL}/add-user-expo-push-token`;
	console.log(
		await sendPOSTRequest(
			url,
			{ expo_push_token: pushToken },
			{
				headers: {
					jwt: `${await JSON.parse(SecureStore.getItemAsync('user'))?.jwt}`,
				},
			}
		)
	);
}
