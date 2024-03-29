'use server';
import axios from 'axios';
import { unstable_noStore as noStore } from "next/cache";

const apiClient = axios.create({
    baseURL: 'http://localhost:8080',
    headers: {
        'Content-Type': "application/json",
    },
});

export const sendMessageStreamChat = async (chatId: string, message: string) => {
    noStore();
    try {
        const response = await apiClient.post('/chat/stream', { chat_id: +chatId, message: message });
        return response.data;
    } catch (error) {
        console.error('API request Error:', error);
        throw error;
    }
};

export const sendMessageJSONChat = async (chatId: string, message: string) => {
    noStore();
    console.log("sendMessageJSONChat called!!!")
    try {
        const response = await apiClient.post('/chat/json', { chat_id: +chatId, message: message });
        return response.data;
    } catch (error) {
        console.error('API request Error:', error);
        throw error;
    }
};

export const fetchStreamChatHistory = async (chatId: string) => {
    noStore();
    try {
        const response = await apiClient.get(`/chat/stream/history?id=${chatId}`);
        return response.data;
    } catch (error) {
        console.error('API request Error:', error);
        throw error;
    }
};

export const fetchJSONChatHistory = async (chatId: string) => {
    noStore();
    console.log("fetchJSONChatHistory called!!!");
    try {
        const response = await apiClient.get(`/chat/json/history?id=${chatId}`);
        return response.data;
    } catch (error) {
        console.error('API request Error:', error);
        throw error;
    }
};
